# for docker-compose
FROM golang:1.19.2-alpine3.16 as build-phase
WORKDIR /var/src
COPY ./go-echo-sample .
RUN ls -lR
RUN pwd
RUN  go mod tidy && \
     CGO_ENABLED=0 \
     GOOS=linux \
     GOARCH=amd64 \
     go build -a -o /var/src/go-echo-sample main.go

FROM scratch as run-phase
COPY --from=build-phase /var/src/go-echo-sample /var/app/go-echo-sample
WORKDIR /var/app
EXPOSE 8080
CMD ["./go-echo-sample"]
