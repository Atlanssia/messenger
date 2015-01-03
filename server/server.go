package server

import (
	"fmt"
	// "io"
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

		go connHandler(conn, l)
	}

}

func connHandler(c net.Conn, l net.Listener) {
	var content string
	var buf []byte
	for {
		n, err := c.Read(buf)
		if err != nil {
			fmt.Println(err)
		}

		if n <= 0 {
			break
		}
		content += string(buf[:n])
	}

	fmt.Println(content)

	// io.Copy(c, c)

	c.Write([]byte("abc"))

	c.Close()
	// l.Close()
}
