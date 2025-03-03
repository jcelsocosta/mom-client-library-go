package mom

import (
	"fmt"
	"io"
	"net"
)

func Subscriber(conn net.Conn) {
	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		if err != io.EOF {
			fmt.Println("Error receiving message:", err)
		}
		return
	}

	message := string(buffer[:n])

	fmt.Println("Message received:", message)
}
