package main

import (
	"context"
	"fmt"
	"grpc_service_logger/internal/config"
	"log"
	"time"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_ = ctx
}
