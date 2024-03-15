package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go holdConn(conn)
	}

}

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

// broadcaster 广播：
func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}

	}
}

func holdConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := struct {
		name string
		ch   chan string
	}{
		name: conn.RemoteAddr().String(),
		ch:   ch,
	}
	ch <- "input your name,and press enter to continue..."
	// who := conn.RemoteAddr().String()
	ch <- "You are " + who.name
	name, _, _ := bufio.NewReader(conn).ReadLine()
	who.name = string(name)
	messages <- who.name + " has arrived"
	entering <- ch

	tick := time.NewTimer(10 * time.Second)
	go func() {
		<-tick.C
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who.name + ": " + input.Text()
		tick.Reset(10 * time.Second)
	}

	leaving <- ch
	messages <- who.name + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
