# for docker-compose
FROM golang:1.19.2-alpine3.16 as binary-build
WORKDIR /var/src
COPY ./go-echo-sample .
RUN ls -lR
RUN pwd
RUN  go mod tidy && \
     CGO_ENABLED=0 \
     GOOS=linux \
     GOARCH=amd64 \
     go build -a -o /var/src/go-echo-sample.bin main.go # 明示的にコンパイル済みバイナリファイル名をつける(.bin)

# 完全にバイナリだけを動かす極小コンテナの作成
FROM scratch as run-phase
# buildしたコンテナからファイルをコピーする
COPY --from=binary-build /var/src/go-echo-sample.bin /var/app/go-echo-sample.bin
COPY --from=binary-build /var/src/sample-if /var/app/sample-if
WORKDIR /var/app
EXPOSE 8080
CMD ["./go-echo-sample.bin"]
