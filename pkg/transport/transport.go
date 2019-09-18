package transport

//Transport is a message transport interface
type Transport interface {
	Send(Message) bool
	Receive() Message
}
