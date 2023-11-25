## 启动

1. 克隆项目：`https://github.com/novohit/wobbs.git`。
2. 修改配置文件 `wobbs/wobbs-server/config/config.yaml`，主要修改数据库、redis相关配置。
3. 创建数据库 `wobbs`。
4. 启动后端服务，运行 `wobbs/wobbs-server/main.go`，会自动创建表。
5. 启动前端服务，进入 `wobbs-frontend` 目录，执行 `npm run dev`。
6. 浏览器访问 `http://localhost:8080`，可以看到如下页面：

![image-20230201151611673](https://zwx-images-1305338888.cos.ap-guangzhou.myqcloud.com/img/2023/02/01/image-20230201151611673.png)

表示启动成功。

## 功能介绍

go练手demo

## 技术栈

- Gin
- Gorm
- MySQL
- Redis
- Vue
- 鉴权：jwt
- 配置文件管理：viper
- 分布式ID：snowflake
- 参数校验：validator
- 日志库：zap
- 热重启：air
