FROM golang:alpine
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR /go/src/TransProxy
COPY . .
RUN mkdir bin \
    && go env \
    && go build -o bin/translate daemon/translate/main.go \
    && go build -o bin/call-insert-trans daemon/call-insert-trans/main.go

# alpine:latest: A minimal Docker image based on Alpine Linux, 因为编译后的可执行二进制文件，可以不需要golang
FROM alpine:latest
WORKDIR /go/src/TransProxy
COPY --from=0 /go/src/TransProxy/bin/translate ./
COPY --from=0 /go/src/TransProxy/bin/call-insert-trans ./
COPY --from=0 /go/src/TransProxy/config.yaml ./
COPY --from=0 /go/src/TransProxy/run_daemon.sh ./

RUN chmod +x ./run_daemon.sh

ENTRYPOINT ./run_daemon.sh


# 打包镜像
# docker build -t 844827150/trans-proxy:{版本号} -f daemon.Dockerfile .
# 启动容器
# docker run --name trans-proxy -d -p 8888:8888 844827150/trans-proxy:{版本号}
