package mom

import (
	"net"
)

func Bind(url string) (net.Conn, error) {
	client, err := net.Dial("tcp", url)

	if err != nil {
		return nil, err
	}

	return client, nil
}
