package transport

// Message is the transported message
type Message struct {
	Body []byte //Message body
}

// String creates string representation of a message
func (m *Message) String() string {
	return string(m.Body)
}

// MessageMarshal encodes message a binary
func MessageMarshal(m Message) []byte {
	return m.Body
}

// MessageUnmarshal decodes message from binary
func MessageUnmarshal(data []byte) Message {
	return Message{
		Body: data,
	}
}
