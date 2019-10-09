package transport

import(
	"errors"
	
	"github.com/streadway/amqp"
)

// RabbitMQTransport is a transport via RabbitMQ
type RabbitMQTransport struct {
	url string // AQMP URL
	connection *amqp.Connection // AMQP connection
	channel *amqp.Channel // AMQP channel
	exchange string // AMQP exchange name
	routingKey string // AMQP routing key
	queue amqp.Queue // AMQP queue
	options RabbitMQTransportOptions // Special AMQP options
}

// RabbitMQTransportOptions are options for RabbitMQ transport
type RabbitMQTransportOptions struct {
	Temporary bool
}

// NewRabbitMQTransport is a constructor for RabbitMQ transport
func NewRabbitMQTransport(url string, exchange string, routingKey string, options RabbitMQTransportOptions) (*RabbitMQTransport, error) {
	// Connect to broker
	conn, err := amqp.Dial(url)
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
	// Create queue
	queue, err := ch.QueueDeclare(
		routingKey, // name
		durable, // durable
		autoDeleted, // delete when unused
		false, // exclusive
		false, // no-wait
		nil, // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}
	// Bind queue
	err = ch.QueueBind(
		queue.Name, // queue name
		routingKey, // routing key
		exchange, // exchange
		false, // no-wait
		nil, // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}
	// Define QoS
	err = ch.Qos(
		1, // prefetch count
		0, // prefetch size
		false, // global
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}
	return &RabbitMQTransport{
		url: url,
		connection: conn,
		channel: ch,
		exchange: exchange,
		routingKey: routingKey,
		queue: queue,
		options: options,
	}, nil
}

// Send sends a message via RabbitMQ
func (t *RabbitMQTransport) Send(m Message) (bool, error) {
	deliveryMode := amqp.Persistent
	if t.options.Temporary {
		deliveryMode = amqp.Transient
	}
	err := t.channel.Publish(
		t.exchange, // exchange
		t.routingKey, // routing key
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
	err := t.channel.Close()
	if err != nil {
		return err
	}
	return t.connection.Close()
}
