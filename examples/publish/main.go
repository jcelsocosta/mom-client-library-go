package main

import (
	"fmt"
	"log"

	mmiddleware "github.com/jcelsocosta/mom-client-library-go"
)

func main() {
	client, err := mmiddleware.NewClient()
	if err != nil {
		log.Println("Erro ao conectar ao servidor:", err)
		return
	}
	defer conn.Close()

	message := mmessage.Publish(conn, "123", "topic.message", "dorgival dantas")
	fmt.Println("Mensagem recebida pelo servidor:", message)

}
