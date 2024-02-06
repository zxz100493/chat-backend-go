FROM golang:alpine3.19 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0 \
    GOOS linux \
    GO111MODULE on \
    GOPROXY https://goproxy.cn,direct

WORKDIR /build/app

COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod download

RUN export GO111MODULE=on && \
    export GOPROXY=https://goproxy.cn && \
    go mod tidy && \
    go mod download

COPY . .

RUN go mod tidy && go build -ldflags="-s -w" -o /app/apps ./cmd
ADD build/config.yaml /app/

FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata curl bash tar yarn
ENV TZ=Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/apps /app/apps
COPY --from=builder /app/config.yaml /app/config.yaml

EXPOSE 8080     

EXPOSE 8888     

CMD ["/app/apps"]