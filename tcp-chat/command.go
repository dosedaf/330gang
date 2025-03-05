package main

type commandID int

const (
	CMD_JOIN commandID = iota
	CMD_MSG
)

type command struct {
	id     commandID
	client *client
	args   []string
}
