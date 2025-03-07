package main

import (
	"bufio"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	username string
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

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		// pass command to chan
		switch cmd {
		case "/username":
			c.commands <- command{
				id:     CMD_USERNAME,
				client: c,
				args:   args,
			}
		case "/join":
			c.commands <- command{
				id:     CMD_JOIN,
				client: c,
				args:   args,
			}
		case "/msg":
			c.commands <- command{
				id:     CMD_MSG,
				client: c,
				args:   args,
			}
		case "/rooms":
			c.commands <- command{
				id:     CMD_ROOMS,
				client: c,
			}
		case "/quit":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
			}
		default:
			c.msg("invalid input! please use the available commands!")
		}
	}
}

func (c *client) err(err error) {
	c.conn.Write([]byte("err: " + err.Error() + "\n"))
}

func (c *client) msg(m string) {
	c.conn.Write([]byte(m + "\n"))
}
