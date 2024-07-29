FROM golang:1.20-alpine AS buildbase
WORKDIR /go/src/github.com/RofaBR/link-shortener-svc
COPY go.mod go.sum ./
RUN apk add git build-base
RUN go mod download
COPY . .
RUN GOOS=linux go build -o /usr/local/bin/link-shortener-svc

FROM alpine:3.9
COPY --from=buildbase /usr/local/bin/link-shortener-svc /usr/local/bin/link-shortener-svc
RUN apk add --no-cache ca-certificates

# Створення папки для логів
RUN mkdir -p /var/log/link-shortener

# Вказання шляху для логів
ENV LOG_PATH=/var/log/link-shortener/log.txt

CMD ["/usr/local/bin/link-shortener-svc"]
