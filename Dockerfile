FROM golang:1.20-alpine AS buildbase

WORKDIR /go/src/github.com/RofaBR/link-shortener-svc

RUN apk add git build-base

COPY . .

RUN GOOS=linux go build -o /usr/local/bin/link-shortener-svc

FROM alpine:3.9
COPY --from=buildbase /usr/local/bin/link-shortener-svc /usr/local/bin/link-shortener-svc
RUN apk add --no-cache ca-certificates

CMD ["/usr/local/bin/link-shortener-svc"]