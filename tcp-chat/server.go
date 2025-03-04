package main

type server struct {
	rooms map[string]*room
}

func newServer() *server {
	return &server{
		rooms: make(map[string]*room),
	}
}
