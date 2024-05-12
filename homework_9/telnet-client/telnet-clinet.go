package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	"github.com/andrei1998Front/go_course/homework_9/telnet-client/config"
)

func readRoutine(conn net.Conn, sigCh chan<- int) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()
		log.Printf("From server: %s", text)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error happend on connection with %s: %v", conn.RemoteAddr(), err)
	}

	conn.Close()
	sigCh <- 1
	log.Printf("Finished readRoutine")
}

func writeRoutine(conn net.Conn, sigCh <-chan int) {
	scanner := bufio.NewScanner(os.Stdin)
LOOPWRITE:
	for {
		select {
		case <-sigCh:
			break LOOPWRITE
		default:
			if !scanner.Scan() {
				break LOOPWRITE
			}
			str := scanner.Text()
			log.Printf("To server %v\n", str)

			conn.Write([]byte(fmt.Sprintf("%s\n", str)))
		}
	}
	log.Printf("Finished writeRoutine")
}

func main() {
	cfg := config.New()

	if err := cfg.Init(); err != nil {
		log.Fatal(err)
	}

	url := cfg.Host + ":" + strconv.Itoa(cfg.Port)

	log.Print("Start tcp-client")
	log.Print("Connecting to " + url + "...")
	conn, err := net.DialTimeout("tcp", url, cfg.Timeout)

	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}
	log.Print("Connection to the " + conn.RemoteAddr().String() + " was successful")

	defer func() {
		conn.Close()
		log.Print("Closing connection to the " + conn.RemoteAddr().String() + " was successful")
	}()

	sigCh := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		readRoutine(conn, sigCh)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		writeRoutine(conn, sigCh)
		wg.Done()
	}()
	wg.Wait()
}
