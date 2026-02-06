package tcp

import (
	"GoKV/internal/auth"
	"log"
	"net"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	authStore, err := auth.NewStore()
	if err != nil {
		return err
	}

	log.Printf("TCP server listening on %s\n", s.addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}

		go handleConnection(conn, authStore)
	}
}
