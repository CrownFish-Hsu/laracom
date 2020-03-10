FROM alpine:latest

# 更新安装源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update && apk add --no-cache bash supervisor
RUN mkdir /app
WORKDIR /app
#将Makefile生成的 demo-service服务二进制文件 copy到容器中 可以直接运行
ADD demo-service /app/demo-service
ADD supervisord.conf /etc/supervisord.conf

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]