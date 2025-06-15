package mmiddleware

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

const (
	SubscriberTypeDefault = 2
)

type Message struct {
	MessageId string
	Message   string
}

func Subscriber(conn net.Conn, topic string) error {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Println("Error server listenner")
	}
	portStr := strings.Split(listener.Addr().String(), ":")[1]

	portInt, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("erro ao converter porta: %w", err)
	}

	buf := new(bytes.Buffer)
	port := uint16(portInt)

	// 1. clientType (1 byte)
	buf.WriteByte(SubscriberTypeDefault)

	// 2. port (2 bytes - uint16)
	binary.Write(buf, binary.BigEndian, uint16(port))

	// 3. topic length (2 bytes - uint16) + topic content (len(topic) bytes)
	binary.Write(buf, binary.BigEndian, uint16(len(topic)))
	buf.WriteString(topic)

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		return err
	}

	listenner(listener)

	return nil
}

func unmarshal(reader *bufio.Reader) (Message, error) {
	var msg Message

	var idLen uint16
	err := binary.Read(reader, binary.BigEndian, &idLen)
	if err != nil {
		fmt.Println("Erro ao ler tamanho do messageId:", err)
		return msg, err
	}

	idBytes := make([]byte, idLen)
	_, err = io.ReadFull(reader, idBytes)
	if err != nil {
		fmt.Println("Erro ao ler messageId:", err)
		return msg, err
	}
	msg.MessageId = string(idBytes)

	var msgLen uint32
	err = binary.Read(reader, binary.BigEndian, &msgLen)
	if err != nil {
		fmt.Println("Erro ao ler tamanho da mensagem:", err)
		return msg, err
	}

	msgBytes := make([]byte, msgLen)
	_, err = io.ReadFull(reader, msgBytes)
	if err != nil {
		fmt.Println("Erro ao ler mensagem:", err)
		return msg, err
	}
	msg.Message = string(msgBytes)
	fmt.Println("msg", msg.Message, msg.MessageId)
	return msg, nil
}

func listenner(listener net.Listener) {
	for {
		subscriberAccept, err := listener.Accept()
		if err != nil {
			fmt.Println("Error ao aceitar conexão", err)
		}

		reader := bufio.NewReader(subscriberAccept)

		unmarshal(reader)
	}
}
