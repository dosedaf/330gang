package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn net.Conn
	name string
	room *room
}

func newClient(c net.Conn) *client {
	return &client{
		conn: c,
		name: "anon",
		room: nil,
	}
}

func (c *client) readInput() {
	reader := bufio.NewReader(c.conn)
	for {
		// e.g. /join mantap
		msg, err := reader.ReadString('\n')
		if err != nil {
			c.conn.Close()
			return
		}

		args := strings.Split(msg, " ")

		cmd := args[0]
		fmt.Println(cmd)

		switch cmd {
		case "/join":
			// c.room = room?
		case "/msg":
			// msg will now send to room?
		}

		// c.conn.Write([]byte("Message received.\n"))
	}
}
