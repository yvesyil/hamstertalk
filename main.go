package main

import (
	"log"
	"net"
)

const (
	protocol = "tcp"
	port     = ":8888"
)

func main() {
	s := newServer()
	go s.run()
	listener, err := net.Listen(protocol, port)
	if err != nil {
		log.Fatalf("Unable to start the server, %v", err.Error())
	}
	defer listener.Close()
	log.Printf("Started the %s server at port %s...", protocol, port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection, %v", err.Error())
			continue
		}
		go s.newHamster(conn)
	}
}
