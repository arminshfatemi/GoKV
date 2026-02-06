package main

import (
	"GoKV/internal/config"
	"GoKV/internal/server/tcp"
	"flag"
	"log"
)

func main() {
	configPath := flag.String("config", "configs/gokv.json", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	srv := tcp.NewServer(cfg)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
