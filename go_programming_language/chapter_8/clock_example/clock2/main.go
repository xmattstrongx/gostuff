package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

// go build -o clock2
// ./clock2 &
// nc localhost 8000
// 20:50:18
// 20:50:19
// 20:50:20
// 20:50:21

// killall clock1

// note this does not stop the app on the port. Have to kill the PID
