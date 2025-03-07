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
		case CMD_USERNAME:
			s.username(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.getRooms(cmd.client)
		case CMD_QUIT:
			s.quit(cmd.client)
		}
	}
}

func (s *server) newClient(c net.Conn) *client {
	return &client{
		conn:     c,
		username: "anon",
		room:     nil,
		commands: s.commands,
	}
}

func (s *server) username(c *client, args []string) {
	if len(args) != 2 {
		c.msg("username can only contains 1 word")
		return
	}

	c.username = args[1]
	c.msg("username set! welcome " + c.username + " !")
}

func (s *server) join(c *client, args []string) {
	fmt.Println("join")
	roomName := args[1]

	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}

		s.rooms[roomName] = r
		// c.msg("room " + roomName + " created!")
	}

	s.quitCurrentRoom(c)

	r.members[c.conn.RemoteAddr()] = c
	c.room = r
	c.msg("joined the room!")
}

func (s *server) msg(c *client, args []string) {
	if c.room == nil {
		c.msg("you havent joined any room yet")
		return
	}
	msg := strings.Join(args[1:], " ")
	c.room.broadcast(c, msg)
}

func (s *server) getRooms(c *client) {
	c.msg("rooms list")
	for key, _ := range s.rooms {
		c.msg("-> " + key)
	}
}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		oldRoom := c.room
		delete(s.rooms[c.room.name].members, c.conn.RemoteAddr()) // hapus addr A dari map members
		oldRoom.broadcast(c, fmt.Sprintf("user %s has left the room!", c.username))
		c.room = nil
		return
	}
}

func (s *server) quit(c *client) {
	s.quitCurrentRoom(c)

	c.msg("gbye!")

	c.conn.Close()
}
