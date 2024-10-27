FROM golang:1.19-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 1
# http://172.16.20.30:32888 为自建的go第三方库代理
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build

ADD go.mod .
ADD go.sum .
# ADD pkg/scanner/ksubdomain/go.mod ./pkg/scanner/ksubdomain/go.mod
COPY . .
COPY static/template /app/static/template
RUN go mod download
COPY cmd/etc /app/etc

RUN go build -ldflags="-s -w" -o /app/cert cmd/main.go

FROM alpine:3.17

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENV TZ Asia/Shanghai

RUN echo -e https://mirrors.aliyun.com/alpine/v3.6/main > /etc/apk/repositories
RUN apk --no-cache update && \
apk --no-cache add tzdata

WORKDIR /app
COPY --from=builder /app/cert /app/cert
COPY --from=builder /app/etc /app/etc
COPY --from=builder /app/static/template /app/static/template

# EXPOSE 8088
ENTRYPOINT ["./cert", "-f"]
CMD ["etc/config.yaml", "-m", "prod"]

