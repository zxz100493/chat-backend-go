FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/app

ADD go.mod .
ADD go.sum .
RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod download

RUN export GO111MODULE=on && \
    export GOPROXY=https://goproxy.cn && \
    go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /app/apps ./cmd
ADD build/config.yaml /app/
ADD build/service.sh /app/service.sh

FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata curl bash tar yarn
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/apps /app/apps
COPY --from=builder /app/config.yaml /app/config.yaml

EXPOSE 8080     

EXPOSE 8888     

CMD ["/app/service.sh"]