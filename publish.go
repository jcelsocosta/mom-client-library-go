package mmiddleware

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	PublishTypeDefault = 1
)

func (client *Client) Publish(messageId string, topic string, message string) error {
	buf := new(bytes.Buffer)

	// 1. clientType (1 byte)
	buf.WriteByte(PublishTypeDefault)

	// 2. messageId (length + content)
	binary.Write(buf, binary.BigEndian, uint16(len(messageId)))
	buf.WriteString(messageId)

	// 3. topic (length + content)
	binary.Write(buf, binary.BigEndian, uint16(len(topic)))
	buf.WriteString(topic)

	// 4. message (length + content)
	binary.Write(buf, binary.BigEndian, uint32(len(message)))
	buf.WriteString(message)

	// Envia tudo
	_, err := client.conn.Write(buf.Bytes())
	fmt.Println(buf.Bytes())
	return err
}
