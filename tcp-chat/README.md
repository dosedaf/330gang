# Go TCP Chat Server

A simple TCP chat server written in Go. This server allows multiple clients to connect, join chat rooms, and send messages to each other in real-time.

## Features

- Set a username with `/username <name>`.
- Join or create a room with `/join <room_name>`.
- Send messages to the current room with `/msg <message>`.
- List all available rooms with `/rooms`.
- Quit the chat with `/quit`.

## How It Works

The server manages multiple chat rooms and clients. When a client connects, they can set a username, join a room, and start chatting. Messages sent in a room are broadcast to all members of that room.

## Running the Server
```bash
go run .
```

## Running the Client
```bash
go run client/hit.go
```
or navigate to the `/client` directory and run:
```bash
go run hit.go
```