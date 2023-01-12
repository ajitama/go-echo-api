# for docker-compose-simple.yaml
FROM golang:1.19.2-alpine3.16 as binary-build
WORKDIR /var/src
COPY ./go-echo-sample .

EXPOSE 8080
CMD ["go","run","main.go"]
