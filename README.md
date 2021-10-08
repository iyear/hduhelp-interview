## 简介

这个项目是面试杭电助手的任务，也是一个自己总结出来的 `Gin Web` 项目模板。

面试任务在根目录下的 `task.md`

以前一直没有用 `Go` 正经写过 `Web` 项目，趁这个机会把开发小型 `Web` 项目的整个流程学习一下，并总结自己的一套项目结构。

## 文档
### API
```
// <host>/doc/index.html
swag init
```
### 前端获取AccessToken
在登录回调成功后，后端会发起一个 `<host>/#/token=<AccessToken>` 的重定向
   
前端可在此阶段获取 `token` 并存于 `Cookie` 或 `LocalStorage`，并在每次请求时将 `token` 置于鉴权头中
## 配置文件

根目录下新建 `config.yml`

第一次启动后需要在 `photos` 与 `departs` 表中手动添加 `id` 为 `-1` 的默认照片和默认部门

```yaml
# 监听端口
port: 8080 
# DEBUG Mode
# 开启后将启用 Gin Logger,Gorm Logger,Swagger Doc
debug: true
# 前端主页地址，末尾带/
url: https://baidu.com/
# 对单客户端的每秒请求次数限制，根据接口、ip拦截，示例为限制两秒一次
limit: 0.5
photo:
  # 图片上传大小范围，单位:KB
  min: 50
  max: 5120
  # 图片MIME限制
  mime:
    - image/jpeg
    - image/png
path:
  # 以 / 结尾
  # 图片保存位置
  photo: ./data/upload/
  # 临时目录(暂时无用)
  temp: ./data/tmp/
mysql:
  # mysql信息
  host: 127.0.0.1
  port: 3306
  user: root
  password: root
  database: hduhelp
auth:
  # 杭电助手OAuth信息
  client_id: xxxxxxxxx
  client_secret: xxxxxxxxxxxxxxx
  # 后端登录Callback地址
  redirect_url: http://localhost:8080/auth
```

## 结构
<details>
<summary>项目结构</summary>

- project
    - api (controller) 版本分离
        - v1
            - stu.go
            - ...
        - api.go  base resp,读参数和响应的处理函数
        - auth.go 登录callback
    - conf 配置及常量级的包
        - e 错误码
            - code.go
            - msg.go
        - app.go 主要配置
        - conf.go 入口
    - db 数据库
        - db.go 入口
        - mysql.go 初始化
        - bolt.go 初始化
        - ...
    - middleware 中间件
        - auth.go 鉴权
        - ...
    - model 模型和请求、响应结构体定义
        - stu.go 学生
        - photo.go 照片
        - ...
    - pkg 其他包
        - oss
        - ...
    - router 路由
        - router.go 路由配置
    - server 服务器
        - server.go 项目入口，初始化和启动服务
    - service 服务
        - srv_stu 学生
            - stu.go 增删改查和与 controller 对接的函数
        - srv_photo 照片
            - photo.go
        - ...
    - util 工具类
        - base.go 基础工具
        - ...
    - main.go 总入口
</details>

-------

补充：这个结构对于该项目粒度还是过细

小型项目可以在 `service` 直接操作数据库

大型项目可以增加 `protocol` 层用于协议转发，同时分离出 `db` 层

## 技术点

开发的过程中解决的一些问题和一些经验

### 上传文件类型校验

根据 https://github.com/gin-gonic/gin#single-file ，`FileName` 不可信，所以判断 `ContentType` 须通过文件头特征。

### 中间件编写

以前一直不敢写中间件，感觉高大上，这次看了官方的 `gin-contrib` 后发现简单的中间件还是很好写的，就是一个流程控制

### Swagger文档

因为这是一个正经的项目，终于试了试 `swagger-go` ，用着真舒服，直接可以在线调试，结构清晰

### 不要用布尔型

用 `int` 的 `1`和 `2` 代替布尔值，布尔值在参数校验和输出要特意关注忽略和默认值导致的误认为未传参，非常麻烦，血泪教训。

### 不要用任何默认值代表实际意义

> 前段时间遇到过这种问题，某个参数是必填的，但是可选的值包括此类型值的零值，最后把这个类型的值设置为指针类型的方式解决的

### .....

## License

MIT