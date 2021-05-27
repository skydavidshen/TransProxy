FROM golang:alpine
ENV GO111MODULE=on
# 国内用国内的镜像，线上需要改成国外的资源URL
ENV GOPROXY=https://goproxy.cn
ENV DIR=/go/src/TransProxy
ARG JOB=daemon-trans

WORKDIR ${DIR}
COPY . .
RUN mkdir bin \
    && go env \
    && chmod +x ./run_daemon_by_one.sh && ./run_daemon_by_one.sh ${JOB} ${DIR}

# alpine:latest: A minimal Docker image based on Alpine Linux, 因为编译后的可执行二进制文件，可以不需要golang
FROM alpine:latest
WORKDIR /go/src/TransProxy
COPY --from=0 /go/src/TransProxy/bin/job ./
COPY --from=0 /go/src/TransProxy/config.yaml ./

ENTRYPOINT ./job


# 打包镜像
# docker build -t 844827150/trans-proxy:{版本号} -f daemon.Dockerfile . --build-arg JOB=call
# docker build -t 844827150/trans-proxy-daemon-trans:1.0.1 --build-arg JOB=trans -f daemon_by_one.Dockerfile .
# docker build -t 844827150/trans-proxy-daemon-call:1.0.1 --build-arg JOB=call -f daemon_by_one.Dockerfile .

# 启动容器
# docker run --name trans-proxy-deamon-trans -d 844827150/trans-proxy-daemon-trans:{版本号}
# docker run --name trans-proxy-deamon-call -d 844827150/trans-proxy-daemon-call:{版本号}
