# for docker-compose
FROM golang:1.19.2-alpine3.16 as binary-build
WORKDIR /var/src
COPY ./go-echo-sample .
RUN apk update && \
    apk add --no-cache git && \
    go get github.com/pilu/fresh # from kurone

EXPOSE 8080
CMD ["fresh"]
CMD ["go","run","main.go"]
