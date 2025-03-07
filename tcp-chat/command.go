package main

type commandID int

// increment ggwp
const (
	CMD_USERNAME commandID = iota
	CMD_JOIN
	CMD_MSG
	CMD_ROOMS
	CMD_QUIT
)

type command struct {
	id     commandID
	client *client
	args   []string
}
