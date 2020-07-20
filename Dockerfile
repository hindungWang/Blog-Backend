FROM golang:1.14.6-alpine as build

ENV GOPROXY https://goproxy.cn/
ENV GO111MODULE on

RUN mkdir -p /usr/local/app
WORKDIR /usr/local/app

COPY . .

RUN go build .

FROM alpine as prod

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

COPY --from=build /usr/local/app/Blog-Backend /Blog-Backend

EXPOSE 8080

CMD ["/Blog-Backend"]