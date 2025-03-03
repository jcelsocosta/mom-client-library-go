package mom

import (
	"fmt"
	"net"
)

func Publish(conn net.Conn, message string) {
	messageBytes := []byte(message)

	_, err := conn.Write(messageBytes)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	fmt.Println("Message sent:", message)
}
