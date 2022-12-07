package main

import (
	"fmt"

	"github.com/perfectogo/template-service-with-kafka/config"
)

func main() {
	cfg := config.Load()
	fmt.Println(cfg)
}
