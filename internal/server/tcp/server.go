package tcp

import (
	"GoKV/internal/auth"
	"GoKV/internal/config"
	"log"
	"net"
)

type Server struct {
	cfg *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{cfg: config}
}

func (s *Server) Start() error {
	adr := s.cfg.Server.Address

	ln, err := net.Listen("tcp", adr)
	if err != nil {
		return err
	}
	defer ln.Close()

	authStore, err := auth.NewStore(s.cfg.Auth.SaltFile)
	if err != nil {
		return err
	}

	log.Printf("TCP server listening on %s\n", adr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}

		go handleConnection(conn, authStore)
	}
}
