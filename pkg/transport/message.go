package transport

//Message is the transported message
type Message struct {
	Body []byte //Message body
}

//String creates string representation of a message
func (message *Message) String() string {
	return string(message.Body)
}
