package main

import (
	"go-echo-api/routes"
	"log"
)

func main() {
	server, err := routes.SetUpEchoServer("./sample-if/go-echo-api.yaml")
	if err != nil {
		log.Fatal(err)
	}
	serverError := server.Start("0.0.0.0:8080")

	// サーバー起動に失敗した場合
	server.Logger.Fatal(serverError)
}
