package main

import (
	"context"
	"log"
	"net/http"
)

//go:generate go run ./../../generator

func main() {
	ctx := context.Background()
	httpHandler := generated.CreateHandler(ctx)
	log.Fatalln(http.ListenAndServe(":3000", httpHandler))
}
