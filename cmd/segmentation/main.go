package main

import (
	"github.com/capkeik/backend-trainee-assignment-2023/internal/config"
	"log"
)

func main() {
	log.Println("Starting Segmentation Service")
	cfg := config.Get()
	_ = cfg
}
