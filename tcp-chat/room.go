package main

import (
	"fmt"
	"net"
)

type room struct {
	name    string
	members map[net.Addr]*client
}

func (r *room) broadcast(sender *client, msg string) {
	fmt.Println("tes")
	for addr, m := range r.members {
		if sender.conn.RemoteAddr() != addr {
			m.msg(sender.username + "> " + msg)
		}
	}
}
