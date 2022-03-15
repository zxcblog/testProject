# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.17-alpine3.15 as builder

# 更新下载软件
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update \
    && apk add --no-cache git ca-certificates make bash yarn nodejs

# 设置go环境变量
RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app

RUN git clone https://gitee.com/zxcblog/test-project.git \
    && cd test-project\
    && go install github.com/swaggo/swag/cmd/swag@latest \
    && swag init \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

FROM alpine:3.15

RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -S app \
    && adduser -S -g app app

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

WORKDIR /app

# 将上一个阶段test-project文件夹下的所有文件复制进来
COPY --from=builder /app/test-project .

RUN chown -R app:app ./

EXPOSE 19610

USER app