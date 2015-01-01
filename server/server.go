package server

import (
	"io"
	"log"
	"net"
)

func Start() {
	l, err := net.Listen("tcp", ":6900")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(conn)
	}
}
