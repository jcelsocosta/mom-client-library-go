package middleware

import (
	"fmt"
	"net"
)

type IMMiddlewareIntegration interface {
	Bind(url string)
}

type MiddlewareIntegration struct{}

func (middlewareIntegration *MiddlewareIntegration) Bind(url string) {
	client, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Error ao se conectar")
		return
	}
	fmt.Println("Conectado")
	defer client.Close()

}
