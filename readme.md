# user service

## version 1 
- register and auto login
- login

### start
#### 1. Setup Basic Dependence
```shell
sudo docker compose up
```

#### 2.Run User RPC Server
```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

#### 3.Run API Server
```shell
cd cmd/api
chmod +x run.sh
sh run.sh
```

### API test
postman desktop
#### 1.register

<img src="images/shot_register.jpg" width="2850"  alt=""/>

#### 2.login
<img src="images/shot_login.jpg" width="2850"  alt=""/>

## 20220525-现在的问题是:

- 1.关注/取关的时候,相应用户的粉丝数量不会动

- 2.关注列表,微服务的server显示成功response,但数据在送往前端的路上报了内存地址错误,怀疑指针数组传来传去写错了

## 20220526

- 基本实现所有API,postman上测试过

- 问题1: 关注/取关仍然可以重复操作

不知道前端能不能重复点关注和取关按钮,不过就算前端有保护,后端也应该拒绝重复操作...待改进

- 问题2:app还没测试通过

目前可能是返回格式还有问题,所以app没有正确显示.但后端能正常响应并返回数据的...待改进

## 20220526: release v0.0.1

### 1. 配置

模仿easy_note,基于docker部署.

数据库和etcd的配置和easy_note一样.

### 2. 功能

- 删除了链路追踪功能

基本保留了easy_note中的设置,比如限制CPU使用等;

- 用户注册

基于jwt实现token分发;

注意: app限制了密码不小于6个字符

- 用户登录

基于jwt实现token校验;

- 用户信息查询

目前只能查自己账号的信息,我看那个app没有搜索用户功能,不知道能不能从视频点进对方个人主页,如果可以的话还要改;

- 关注列表/粉丝列表查询

配合app上的关注/取关按钮还能实现关注/取关操作,可以反复点;

- 关注/取关操作

目前实现逻辑上还比较简单,处理正常逻辑没问题,但没有做成数据库事务(这样误操作还能回滚);

### 3. start
#### 1. Setup Basic Dependence
```shell
sudo docker compose up
```

#### 2. Run User RPC Server
```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

#### 3. Run API Server
```shell
cd cmd/api
chmod +x run.sh
sh run.sh
```

### 4. Postman测试

#### 用户注册

请求:

```shell
POST
127.0.0.1:8080/douyin/user/register/?username=yang&password=123456
```

响应:

```shell
{
    "status_code": 0,
    "status_msg": "Success",
    "user_id": 7,
    "name": "user",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTM1NjgyNjMsImlkIjo3LCJvcmlnX2lhdCI6MTY1MzU2NDY2M30.NR6ZOiANDBNijCHm4o9vIIYOWCoPjBNiM-M8TFsgZD4",
    "expire": "2022-05-26T20:31:03+08:00"
}
```

#### 用户登录

请求:

```shell
POST
127.0.0.1:8080/douyin/user/login/?username=bjz&password=123456
```

响应:

```shell
{
    "status_code": 0,
    "status_msg": "Success",
    "user_id": 1,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTM1NjYwNjcsImlkIjoxLCJvcmlnX2lhdCI6MTY1MzU2MjQ2N30.tbqK6lKg9oCTxkYXnNjK8pu2FxArSbHD7Kfx1nerOyc",
    "expire": "2022-05-26T19:54:27+08:00"
}
```


#### 用户信息查询

请求:

```shell
GET
127.0.0.1:8080/douyin/user/          # user_id可加可不加,因为我的id从token中拿
header 'Authorization: Bearer $token'  # 在postman中设置header字段
```

响应:

```shell
{
    "status_code": 0,
    "status_msg": "Success",
    "id": 1,
    "name": "bjz",
    "follow_count": 4,
    "follower_count": 2,
    "is_follow": false
}
```

#### 关注

请求:

```shell
POST
127.0.0.1:8080/douyin/relation/action/?user_id=1&to_user_id=4&action_type=1
header 'Authorization: Bearer $token'  # 在postman中设置header字段
```

响应:

```shell
{
    "status_code": 0,
    "status_msg": "Success",
    "data": NULL
}
```

#### 取关

请求:

```shell
POST
127.0.0.1:8080/douyin/relation/action/?user_id=1&to_user_id=4&action_type=2
header 'Authorization: Bearer $token'  # 在postman中设置header字段
```

响应:

```shell
{
    "status_code": 0,
    "status_msg": "Success",
    "data": NULL
}
```

#### 关注列表

请求:

```shell
GET
127.0.0.1:8080/douyin/relation/follow/list/?user_id=1
header 'Authorization: Bearer $token'  # 在postman中设置header字段
```

响应:

```shell
{
    "status_code": 0,
    "status_msg": "Success",
    "user_list": [
        {
            "id": 2,
            "name": "leafccc",
            "follow_count": 1,
            "follower_count": 1,
            "is_follow": true
        },
        {
            "id": 3,
            "name": "Bob",
            "follow_count": 0,
            "follower_count": 1,
            "is_follow": true
        },
        {
            "id": 5,
            "name": "Bob1",
            "follow_count": 0,
            "follower_count": 1,
            "is_follow": true
        },
        {
            "id": 6,
            "name": "yang",
            "follow_count": 0,
            "follower_count": 1,
            "is_follow": true
        }
    ]
}
```

#### 粉丝列表

请求:

```shell
GET
127.0.0.1:8080/douyin/relation/follower/list/?user_id=1
header 'Authorization: Bearer $token'  # 在postman中设置header字段
```

响应:

```shell
{
    "status_code": 0,
    "status_msg": "Success",
    "user_list": [
        {
            "id": 2,
            "name": "leafccc",
            "follow_count": 1,
            "follower_count": 1,
            "is_follow": true
        },
        {
            "id": 4,
            "name": "little Bob",
            "follow_count": 1,
            "follower_count": 0,
            "is_follow": false
        }
    ]
}
```

### 5. APP测试

我是在ubuntu双系统下跑的,电脑连宿舍WiFi,然后查询电脑IP(比如是192.168.x.x).

手机上安装最新版app(注:字节那边经常会更新),双击右下角的"我",将url设置为: http://192.168.x.x:8080/ , 保存并重启.

单击"我",可以登录数据库中已有的账户; 如果没有账号,在登录页面点击"注册账号",注册完成后会自动登录.

登录后app会直接跳转到个人主页.由于目前的api还不全,此时的主页是假主页(字节说了,app会判断个人主页的接口是否全部实现,如果不是就返回假主页),但是可以点击"关注"/"粉丝",跳转后的关注列表和粉丝列表就是真的了(需要通过Postman预先插入一些关注行为,因为目前还不支持在app中搜索用户).

在"关注"和"粉丝"中可以看到数量和名字,点击对应用户右侧的"关注"和"取消关注"即可以实现关注和取关操作.但是页面不会实时刷新,要刷新可以退出关注/粉丝列表然后重新进入."关注"/"取消关注"按钮点击后会自动toggle,可以反复点,这样就反复关注/取关.
