package mom

import (
	"net"
)

const (
	PublishTypeDefault = 1
) // Value for clientType

func Publish(conn net.Conn, message string) error {
	packet := make([]byte, 1+len(message))

	packet[0] = PublishTypeDefault

	copy(packet[1:], message)

	_, err := conn.Write(packet)
	if err != nil {
		return err
	}
	return nil
}
