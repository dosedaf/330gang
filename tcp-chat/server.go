package main

import (
	"fmt"
	"net"
	"strings"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_MSG:
			fmt.Println("here")
			s.msg(cmd.client, cmd.args)
		}
	}
}

func (s *server) newClient(c net.Conn) *client {
	return &client{
		conn:     c,
		name:     "anon",
		room:     nil,
		commands: s.commands,
	}
}

func (s *server) join(c *client, args []string) {
	roomName := args[1]

	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}

		s.rooms[roomName] = r
	}

	r.members[c.conn.RemoteAddr()] = c
	c.room = r
}

func (s *server) msg(c *client, args []string) {
	msg := strings.Join(args[1:], " ")
	c.room.broadcast(c, msg)
}
