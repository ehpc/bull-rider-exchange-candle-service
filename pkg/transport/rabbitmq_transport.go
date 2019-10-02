package transport

// RabbitMQTransport is a transport via RabbitMQ
type RabbitMQTransport struct{

}

// NewRabbitMQTransport is a constructor for RabbitMQ transport
func NewRabbitMQTransport() RabbitMQTransport {
	return RabbitMQTransport{}
}


// Send sends a message via RabbitMQ
func (t *RabbitMQTransport) Send(m Message) (bool, error) {
	return false, nil
}

// Receive a message via RabbitMQ
func (t *RabbitMQTransport) Receive(rp RequestParams) (chan Message, chan error) {
	return nil, nil
}
