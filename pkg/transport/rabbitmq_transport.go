package transport

import(
	"os"
	"errors"
	"github.com/streadway/amqp"
)

// RabbitMQTransport is a transport via RabbitMQ
type RabbitMQTransport struct {
	Connection *amqp.Connection
	Channel *amqp.Channel
	Exchange string
	RoutingKey string
	Options RabbitMQTransportOptions
}

// RabbitMQTransportOptions are options for RabbitMQ transport
type RabbitMQTransportOptions struct {
	Temporary bool
}

// NewRabbitMQTransport is a constructor for RabbitMQ transport
func NewRabbitMQTransport(exchange string, routingKey string, options RabbitMQTransportOptions) (*RabbitMQTransport, error) {
	// Connect to broker
	conn, err := amqp.Dial(os.Getenv("BROKER_URL"))
	if err != nil {
		return nil, err
	}
	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}
	// Create an exchange
	durable := true
	autoDeleted := false
	if options.Temporary {
		durable = false
		autoDeleted = true
	}
	err = ch.ExchangeDeclare(
		exchange, // name
		"direct", // type
		durable, // durable
		autoDeleted, // auto-deleted
		false, // internal
		false, // no-wait
		nil, // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}
	return &RabbitMQTransport{
		Connection: conn,
		Channel: ch,
		Exchange: exchange,
		RoutingKey: routingKey,
		Options: options,
	}, nil
}

// Send sends a message via RabbitMQ
func (t *RabbitMQTransport) Send(m Message) (bool, error) {
	deliveryMode := amqp.Persistent
	if t.Options.Temporary {
		deliveryMode = amqp.Transient
	}
	err := t.Channel.Publish(
		t.Exchange, // exchange
		t.RoutingKey, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			DeliveryMode: deliveryMode,
			ContentType: "application/protobuf",
			Body: MessageMarshal(m),
		},
	)
	if err != nil {
		return false, err
	}	
	return true, nil
}

// Receive a message via RabbitMQ
func (t *RabbitMQTransport) Receive(rp RequestParams) (chan Message, chan error) {
	ch := make(chan error, 1)
	defer close(ch)
	ch <- errors.New("rabbitmqtransport: Receive not implemented")
	return nil, ch
}

// Close closes transport
func (t *RabbitMQTransport) Close() error {
	err := t.Channel.Close()
	if err != nil {
		return err
	}
	return t.Connection.Close()
}
