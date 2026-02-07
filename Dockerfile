# 使用官方Golang镜像作为构建环境
FROM golang:1.24.7-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件（如果存在）
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 使用轻量级alpine镜像作为运行环境
FROM alpine:latest

# 安装必要的包
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从builder阶段复制编译好的二进制文件
COPY --from=builder /app/main .

# 暴露端口（根据main.go中看到的应用监听端口）
EXPOSE 8991 8992

# 运行应用
CMD ["./main"]