# MYJWT Middleware for Gin Framework

本项目基于gin-jwt框架开发,扩展了以下两个功能:

1. 支持从POST请求的表单数据中获取并解析token;

2. 支持同一条路由同时支持未登录和登录用户访问.

这些功能按照原框架的代码风格开发, 使用时只需在authMiddleware中间件中增加相应配置即可.

## usage

### 扩展功能1

设置TokenLookup, 增加从postform的"token"字段中解析token:

```Go
TokenLookup:   "header: Authorization, query: token, cookie: jwt, postform: token", // 最后一个是我们新增的解析字段
```

即可支持从POST请求的表单数据中获取并解析token.

### 扩展功能2

设置FilteredURL, 填入想要同时支持未登录和登录用户访问的URL. 如果有多个路由, 用逗号隔开:

```Go
FilteredURL:   "/douyin/feed, /douyin/publish/list, /douyin/favorite/list/, /douyin/comment/list/, /douyin/relation/follow/list/, /douyin/relation/follower/list/", // 设置你需要跳过认证的url
```

设置后, 这些路由只需和需要认证的路由一样, 使用authMiddleware中间件即可.