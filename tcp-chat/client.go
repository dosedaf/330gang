package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	name     string
	room     *room
	commands chan<- command
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

		cmd := strings.TrimSuffix(args[0], "\n")

		// pass command to chan
		switch cmd {
		case "/join":
			c.commands <- command{
				id:     CMD_JOIN,
				client: c,
				args:   args,
			}
		case "/msg":
			fmt.Println("here?")
			c.commands <- command{
				id:     CMD_MSG,
				client: c,
				args:   args,
			}
		}

		// c.conn.Write([]byte("Message received.\n"))
	}
}

func (c *client) err(err error) {
	c.conn.Write([]byte("err: " + err.Error() + "\n"))
}

func (c *client) msg(m string) {
	c.conn.Write([]byte(">" + m + "\n"))
}
