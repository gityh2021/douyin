# 配置数据库读写分离

## 说明

主库负责增删改，从库负责查询。

## 执行`docker-compose.yaml`

在当前目录下执行`sudo docker-compose up -d`构建镜像，构建完成后,

执行`sudo docker-compose ps -a` 查看容器状态。

确认两个节点已经启动。

## 确认节点的网络状态

执行`sudo docker network ls`查看两个节点的网络状态

![](https://fastly.jsdelivr.net/gh/Draculabo/gallery/img/202206051531542.png)

再执行`docker inspect 70db4c6c062d`查看详细信息

![](https://fastly.jsdelivr.net/gh/Draculabo/gallery/img/202206051532492.png)

获取到两个节点的`IP`地址，因为Docker网络的配置，两个连接到同一network的容器会直接相互连通。

## 开始进行主从配置

分别进入`master`主节点和`slave`从节点

例如：`docker exec -it 7453daedfa7d /bin/bash`

### `master`节点

1. 执行`mysql -u root -p douyin`进入数据库

2. 执行`show master status`确认**日志文件名称** 和 **数据同步起始位置**

![](https://fastly.jsdelivr.net/gh/Draculabo/gallery/img/202206051533385.png)

### slave节点

1. 执行`mysql -u root -p douyin`进入数据库

2. 执行`CHANGE MASTER TO MASTER_HOST='172.19.0.2', MASTER_USER='root', MASTER_PASSWORD='root', MASTER_LOG_FILE='replicas-mysql-bin.000004',`配置主节点信息 。

   > 说明：`MASTER_HOST`为之前查看的`master`节点的`IP`地址，`MASTER_LOG_FILE`为`master`节点的日志文件名称

3. 执行` show slave status\G;`查看`salve`节点状态，保证标红处都为`Yes`

![](https://fastly.jsdelivr.net/gh/Draculabo/gallery/img/202206051533629.png)

## 测试

可以使用`navicat`连接数据库进行测试

1. 在`linux`系统中执行`ifconfig`查看`IP`地址

2. 执行`docker ps -a`查看容器状态，获取端口号。

   > `docker-compose.yaml`中默认配置`master`为33065，`slave`为33066

3. 连接数据库，根用户名为`root`

![](https://fastly.jsdelivr.net/gh/Draculabo/gallery/img/202206051533036.png)

4. 在`master`主库中的`douyin`数据库创建表

```
SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for course
-- ----------------------------
DROP TABLE IF EXISTS `course`;
CREATE TABLE `course` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `lesson_period` double(5,0) DEFAULT NULL,
  `score` double(10,0) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

5. 如果看到`slave`从库中也出现了对应的表，则主从成功配置