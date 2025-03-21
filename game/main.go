package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Running on port %d\n", ln.Addr().(*net.TCPAddr).Port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("New Connection From %s\n", conn.RemoteAddr())
	fmt.Fprintln(conn, "Hello client")

	buffer := make([]byte, 1024)

	x := make(chan int)
	y := make(chan int)

	go game(conn, x, y)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Fprintln(conn, err.Error())
			break
		}
		fmt.Printf("Received from %s: %s", conn.RemoteAddr(), buffer[:n])
		handleMessage(conn, string(buffer[:n]), x, y)
	}
}
func handleMessage(conn net.Conn, message string, x, y chan int) {
	switch message {
	case "w":
		cy := <-y
		cy++
		y <- cy
	case "clear\n":
		fmt.Fprint(conn, "\033[2J")
	}
}
func game(conn net.Conn, x, y chan int) {
	x <- int(90/2) - 5
	y <- int(15/2) - 1
	for {
		fmt.Fprint(conn, "\033[2J", grid(90, 15, x, y), "\n")
		<-time.After(time.Millisecond * 500)
	}
}
func grid(w, h int, x, y chan int) (result string) {
	// x := int(math.Floor(rand.Float64() * float64(w)))
	// y := int(math.Floor(rand.Float64() * float64(h)))
	// x := int(w/2) - 5
	// y := int(h/2) - 1
	fmt.Println(<-x, <-y)
	for i := range h {
		for j := range w {
			cx := <-x
			cy := <-y
			if cx <= j && cx+5 > j && cy <= i && cy+3 > i {
				result += "$"
			} else {
				result += "."
			}
			if j == w-1 {
				result += "\n"
			}
		}
	}
	return
}
