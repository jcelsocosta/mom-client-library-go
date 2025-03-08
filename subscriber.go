package mom

import (
	"fmt"
	"io"
	"net"
)

const (
	SubscriberTypeDefault = 2
) // Value for clientType

func Subscriber(conn net.Conn) {
	packet := make([]byte, 1)
	packet[0] = SubscriberTypeDefault

	_, err := conn.Write(packet)

	if err != nil {
		fmt.Println("Error conn consumer")
		return
	}

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error receiving message:", err)
				return
			}
			return
		}

		message := string(buffer[:n])

		fmt.Println("Message received:", message)
	}
}
