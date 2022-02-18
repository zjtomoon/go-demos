FROM golang:1.15 as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

# 设置工作区
WORKDIR /go/cache

#复制文件下载依赖
ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/build

ADD . .

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o app main.go

FROM scratch as runner

# exe

# views

# static

# conf
EXPOSE 3000

#启动服务
CMD ["/app"]

