package mmiddleware

import (
	"net"
)

type Client struct {
	conn net.Conn
}

func (client *Client) NewClient(url string) (*Client, error) {
	conn, err := net.Dial("tcp", url)

	if err != nil {
		return nil, err
	}

	return &Client{conn: conn}, nil
}
