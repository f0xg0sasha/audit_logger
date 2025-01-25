package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/f0xg0sasha/audit_logger/internal/config"
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
