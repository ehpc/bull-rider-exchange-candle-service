package transport

// Transport is a message transport interface
type Transport interface {
	Send(Message) (bool, error)
	Receive(RequestParams) (chan Message, chan error)
	Close() error
}
