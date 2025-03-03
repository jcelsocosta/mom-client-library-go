package middleware

import (
	"fmt"
	"net"
)

func Bind(url string) (net.Conn, error) {
	client, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Error ao se conectar")
		return nil, err
	}

	return client, nil
}
