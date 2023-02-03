FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 在含go环境的镜像中将代码编译成二进制可执行文件 app
RUN go build -o wobbs_app .

###################
# 接下来创建一个小镜像 因为此时我们已经得到可执行文件了 不需要有go环境了
###################
FROM debian:stretch-slim

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY ./wait-for.sh /
COPY ./template template/
COPY ./static static/
COPY ./config config/
COPY --from=builder /build/wobbs_app /

# set命令用法 https://www.ruanyifeng.com/blog/2017/11/bash-set.html
RUN set -eux; \
    cp /etc/apt/sources.list /etc/apt/sources.list.bak; \
    sed -i 's/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list; \
    sed -i 's/security.debian.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list; \
	apt-get update; \
	apt-get install -y \
		--no-install-recommends \
		netcat; \
        chmod 755 /wait-for.sh

EXPOSE 8088

# 需要运行的命令
ENTRYPOINT ["/wobbs_app", "-f", "config/config.yaml"]