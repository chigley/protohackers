package smoke

import (
	"fmt"
	"io"
	"log"
	"net"
)

func Listen(address string) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("smoke: listen error: %w", err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			return fmt.Errorf("smoke: accept error: %w", err)
		}

		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	if _, err := io.Copy(c, c); err != nil {
		log.Printf("smoke: copy error: %s\n", err)
	}
}
