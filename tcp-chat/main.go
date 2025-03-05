package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("error accepting connection : %s", err.Error())
			return
		}

		c := s.newClient(conn)

		go c.readInput()
	}
}
