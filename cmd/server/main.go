package main

import (
	"GoKV/internal/server/tcp"
	"log"
)

func main() {
	srv := tcp.NewServer(":6379")
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
