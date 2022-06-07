**本项目是2022字节跳动后端青训营项目，具体的说明文档地址为：https://www.yuque.com/yanghong-kpkbj/gg6r7n/wiexpw**

## 1、需求规格说明
- 项目地址：[https://github.com/gityh2021/douyin](https://github.com/gityh2021/douyin)
- 线上地址：[http://139.196.152.44:31986/](http://139.196.152.44:31986/)
- 项目介绍：

实现简单版本的抖音，通过对项目的不断迭代开发，完成对课程各个知识点的实践。

- 项目功能说明：
| **功能项** | **说明** |
| --- | --- |
| 视频 Feed 流、视频投稿、个人信息 | 支持所有用户刷抖音，按投稿时间倒序推出，登录用户可以自己拍视频投稿，查看自己的基本信息和投稿列表，注册用户流程简化。 |
| 点赞列表、用户评论 | 登录用户可以对视频点赞，并在视频下进行评论，在个人主页能够查看点赞视频列表。 |
| 关注列表、粉丝列表 | 登录用户可以关注其他用户，能够在个人信息页查看本人的关注数和粉丝数，点击打开关注列表和粉丝列表。 |

## 2、软件概要设计说明

- **总体框架：**

项目采用微服务架构，总共分成三个服务：api、video以及user。其中video负责视频相关数据的读写，包括视频信息，视频的评论以及视频的点赞。user模块则负责用户相关行为，包括用户的登录，注册，关注等行为。video和user都会分别注册成RPC服务，供api服务调用。Api服务则是对外提供Http服务的service，用来接收客户端的请求，在API服务内内对video和user的数据进行整合将数据返回给客户端。
框架调用图：
![image.png](https://cdn.nlark.com/yuque/0/2022/png/13007828/1654568375426-cdf731ba-7d6b-421a-93e5-ac57682c09a5.png#clientId=ua6c16e40-bb9b-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=409&id=u617682a7&margin=%5Bobject%20Object%5D&name=image.png&originHeight=608&originWidth=860&originalType=binary&ratio=1&rotation=0&showTitle=false&size=44445&status=done&style=none&taskId=u7a2797ff-c02e-4ce4-b0b4-51e3a888d0a&title=&width=578)

- **api服务**

api服务负责对外提供HTTP服务，接受客户端的请求并解析域名并调用相应的RPC服务，将数据返回给客户端。实现的接口如下：

|  | **解析域名** | **RPC调用函数** |
| --- | --- | --- |
| 基础接口 | /douyin/feed - 视频流接口 | rpc.GetVideosFeed()
rpc.GetUsersByIds() |
|  | /douyin/user/register/ - 用户注册接口 | rpc.CreateUser()
rpc.CheckUser() |
|  | /douyin/user/login/ -用户登录接口 | rpc.CheckUser() |
|  | /douyin/user/ - 用户信息 | rpc.InfoGetUser() |
|  | /douyin/publish/action/ - 视频投稿 | rpc.CreateVideo() |
|  | /douyin/publish/list - 发布列表 | rpc.GetPublishVideoList()
rpc.GetUsersByIds() |
| 扩展接口-1 | /douyin/favourite/action/ - 赞操作 | rpc.FavoriteByUser() |
|  | /douyin/favourite/list/ - 点赞列表 | rpc.GetFavoriteList() |
|  | /douyin/comment/action/ - 评论操作 | rpc.PostComment() |
|  | /douyin/comment/list/ - 视频评论列表 | rpc.GetCommentsByVideoId()
rpc.GetUsersByIds() |
| 扩展接口-2 | /douyin/relation/action/ - 关系操作 | rpc.UpdateUser() |
|  | /douyin/relation/follow/list/ - 用户关注列表 | rpc.MGetUser() |
|  | /douyin/relation/follower/list/ - 用户粉丝列表 | rpc.MGetUser() |

- **user服务**

user服务负责用户相关行为，如注册登录与社交功能。提供的服务如下：

| **服务名称** | **服务功能** | **说明** |
| --- | --- | --- |
| CreateUser | 用户注册 | 检查用户名与密码并创建一个新用户。 |
| CheckUser | 用户登录 | 检查用户名与密码并登录。 |
| InfoGetUser | 获取用户信息 | 根据用户ID查询登录用户相关信息 |
| MGetUser  | 获取关注/粉丝列表  | 根据传入user_id查询其关注/粉丝列表，同时查询登录用户是否已关注这些用户（未登录默认未关注）。 |
| UpdateUser | 关注/取关用户 | 根据传入user_id与token执行关注/取关操作 |
| GetUserInfoList | 批量获取用户信息 | 根据用户ID列表批量获取用户信息，同时查询登录用户是否已关注这些用户。用于获取视频作者列表、获取用户关注、粉丝列表。 |

- **video服务**

video负责视频相关数据的读写，包括视频信息，视频的评论以及视频的点赞。提供的服务如下：

| **服务名称** | **服务功能** | **说明** |
| --- | --- | --- |
| GetPublishListByUser | 获取作品列表 | 根据用户ID查询其所有发布的视频。 |
| GetVideosByLastTime | 获取视频流 | 获取首页视频流， 以时间倒序排列。 |
| PublishVideo | 发布视频 | 保存用户发布的视频。 |
| FavoriteByUser | 赞操作 | 执行点赞、取消点赞的操作。 |
| GetFavoriteListBYUser | 点赞列表 | 根据用户ID获取该用户的点赞过的视频列表。 |
| GetCommentListByVideo | 评论列表 | 根据视频ID获取该视频的评论列表，包括评论时间、评论内容以及评论用户。 |
| PostComment | 发布评论 | 登录用户发布评论。 |

## 3、数据库设计说明

- **数据库表单**
| **表单名** | **表单字段** | **说明** |
| --- | --- | --- |
| user | user_name, password, follow_count, follower_count | 存储所有注册用户的信息，密码存储为md5。 |
| videos | author_id, play_url, cover_url, favourite_count, comment_count, is_favourite, title | 存储所有上传视频的信息。 |
| favourites | user_id, video_id | 存储所有用户点赞视频的记录。 |
| comments | user_id, video_id, content | 存储所有视频的评论及评论者。 |
| follower | user_id, follower_id | 存储用户之间关注关系。 |

## 4、项目非功能性和高可用保证。

- **拓展gin-jwt保证网关功能。**

本项目在使用 github.com/appleboy/gin-jwt/v2 框架实现JWT认证时，遇到了几个该框架不支持的问题：

   1. 发布视频的请求将token放在POST请求的表单数据中，而该gin-jwt框架不支持从这里获取token的内容；
   1. 抖声APP的许多请求同时支持登录状态和未登录状态(比如说，登录用户和未登录用户都可以获取视频流、获取其他用户信息等），而且登录状态和未登录状态的这些请求路由是一样的。而该gin-jwt框架对于同一条路由只能开启认证或者不开启认证。具体来说，如果获取视频流的路由不开启认证，那么登录用户访问时，我们就无法从token字段中获取当前登录用户的id；如果获取视频流的路由开启认证，那么未登录用户访问时就会因为认证不通过（token为空）而直接返回错误。

为了解决上述两个问题，我们在该框架的基础上进行了拓展，实现了：

   1. 增加从POST请求的表单数据中获取和解析token内容的功能；
   1. 增加“过滤URL”功能，对于同一条路由支持处理登录用户和未登录用户的请求。

具体实现方法可以看我的笔记：[扩展gin-jwt | 青训营笔记](https://juejin.cn/post/7105422755863986212)，也可以看子文档：

- **使用K8S部署。**

本项目使用K8S进行部署，实现每个服务的多Pod实例，即使某一个副本挂了也可以有其它副本可用，并且K8S还会自动重启副本，之后还可以有动态扩缩容的机制。通过这些手段来增强服务的高可用。并且在部署过程中实现端口与服务分离，可以灵活部署服务。具体见6:部署上线。

- **数据库读写分离，一主一从。**

        见[数据库主从同步配置文档](https://www.yuque.com/yanghong-kpkbj/gg6r7n/cwdr95)

- **数据库事务。**

 数据库的事务(Transaction)是一种机制、一个操作序列，包含了一组数据库操作命令。事务把所有的命令作为一个整体一起向系统提交或撤销操作请求，即这一组数据库命令要么都执行，要么都不执行，因此事务是一个不可分割的工作逻辑单元。  
在我们的一些数据库函数中，涉及多个增删补改的操作的，我们都增加了事务机制。
例如在./cmd/user/dal/db/user.go中的UpdateUser函数操作，当一个用户想要关注另一个用户，需要涉及User表以及Follower的操作。User表中follow_count与follower_count需要修改，并且Follower表中需要增删字段。这样就需要我们添加事务，以保证原子性、一致性、隔离性和持久性。
详细的代码讲解在子文档：[事务操作文档](https://www.yuque.com/yanghong-kpkbj/gg6r7n/vr680u)
## 5、软件测试
基于Postman进行测试：建立测试组
![image.png](https://cdn.nlark.com/yuque/0/2022/png/13007828/1654582988535-af4e7859-a425-47b0-a93a-5a72d6dc3cee.png#clientId=u0854dacc-c24b-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=799&id=ub244faff&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1598&originWidth=2840&originalType=binary&ratio=1&rotation=0&showTitle=false&size=460154&status=done&style=none&taskId=u6b3c8e32-7d17-4176-911d-54bc9e87430&title=&width=1420)
## 6、部署上线

- **服务高可用：**

创建三个K8S service实例，并为每个service配置2-3个副本，保证服务的高可用性。
```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: api-service
  labels:
    app: api-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: api-service
  template:
    metadata:
      labels:
        app: api-service
    spec:
      containers:
        - name: video-service
          image: fduyh2021/api:v2
          imagePullPolicy: Always
          env:
            - name: NETWORK_IP
              value: "172.19.109.141"
            - name: USER_PORT
              value: "30111"
            - name: VIDEO_PORT
              value: "30112"
          ports:
            - containerPort: 8082
          resources:
            requests:
              cpu: 100m
              memory: 500Mi
```

- **灵活部署：**

所有的服务端口都以环境变量的形式传入文件中进行部署，保证端口可以灵活配置。
```dockerfile
FROM golang:1.17.2-alpine
ENV USER_PORT=8087 \
    VIDEO_PORT=8088 \
    API_PORT=8082
RUN go env -w GO111MODULE="on"
RUN go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR $GOPATH/douyin
COPY . $GOPATH/douyin
RUN go build ./cmd/api/
EXPOSE 8082
ENTRYPOINT ["./api"]
```

- **部署上线：**[http://139.196.152.44:31986/](http://139.196.152.44:31986/)

![image.png](https://cdn.nlark.com/yuque/0/2022/png/13007828/1654569079495-e28d0f50-da45-4514-9869-88d0ca9377ba.png#clientId=ua6c16e40-bb9b-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=452&id=u08deff44&margin=%5Bobject%20Object%5D&name=image.png&originHeight=904&originWidth=2248&originalType=binary&ratio=1&rotation=0&showTitle=false&size=391964&status=done&style=none&taskId=ufeb5caf9-c7f0-4a74-b9b4-b46b5c9bfff&title=&width=1124)

