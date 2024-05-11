package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/andrei1998Front/go_course/homework_9/telnet-server/config"
)

func handleConnection(ctx context.Context, conn net.Conn) {
	defer conn.Close()

	msg := fmt.Sprintf("Welcome to %s, friend from %s\n", conn.LocalAddr(), conn.RemoteAddr())
	conn.Write([]byte(msg))

	scanner := bufio.NewScanner(conn)
Loop:
	for {
		select {
		case <-ctx.Done():
			log.Print("TIMEOUT")
			conn.Write([]byte("TIMEOUT"))
			break Loop
		default:
			if !scanner.Scan() {
				log.Print("CAN NOT SCAN")
				break Loop
			}
			text := scanner.Text()
			if text == "quit" || text == "exit" {
				break Loop
			}

			log.Printf("RECEIVED: %s", text)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error happend on connection with %s: %v", conn.RemoteAddr(), err)
	}

	log.Printf("Closing connection with %s", conn.RemoteAddr())
}

func main() {
	cfg := config.New()
	if err := cfg.Init(); err != nil {
		log.Fatal(err)
	}

	port := "127.0.0.1:" + strconv.Itoa(cfg.Port)

	ctx, _ := context.WithTimeout(context.Background(), cfg.Timeout)

	l, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Cannot accept: %v", err)
		}
		go handleConnection(ctx, conn)
	}
}
