Todo

Obter o endereço IP da máquina em questão
construir uma lib completa
unmarshall e marshall
reconectar ao servidor


//sugestão chat gpt
type Client struct {
    conn net.Conn
}

func NewClient(addr string) (*Client, error) {
    conn, err := net.Dial("tcp", addr)
    if err != nil {
        return nil, err
    }
    return &Client{conn: conn}, nil
}

func (c *Client) Publish(id, topic, message string) (ResponseType, error) {
    // lógica
}