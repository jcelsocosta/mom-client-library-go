package main

import (
	"fmt"
	"log"
	"net"

	mmessage "github.com/jcelsocosta/mom-client-library-go"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("Erro ao conectar ao servidor:", err)
		return
	}
	defer conn.Close()

	err = mmessage.Subscriber(conn, "topic.message")
	fmt.Println("error", err)
}
