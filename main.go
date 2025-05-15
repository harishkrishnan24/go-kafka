package main

import (
	"fmt"
	"io"
	"log/slog"
	"net"
)

type Message struct {
	data []byte
}

type Server struct {
	coffsets map[string]int
	buffer   []Message
	ln       net.Listener
}

func (s *Server) Start() error {
	return nil
}

func NewServer() *Server {
	return &Server{
		coffsets: make(map[string]int),
		buffer:   make([]Message, 0),
	}
}

func (s *Server) Listen() error {
	ln, err := net.Listen("tcp", ":9092")
	if err != nil {
		return err
	}
	s.ln = ln
	for {
		conn, err := ln.Accept()
		if err != nil {
			if err == io.EOF {
				return err
			}
			slog.Error("server accept error", "err", err)
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	fmt.Println("new connection", conn.RemoteAddr())

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			slog.Error("connection read error", "err", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

func main() {
	server := NewServer()
	server.Listen()
}
