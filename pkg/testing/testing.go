package testing

import (
	"ehpc.io/bull-rider/exchange-candle-service/pkg/transport"
)

//TransportMock is a mock for transport layer
type TransportMock struct {
	sentMessages       []transport.Message
	receivableMessages []transport.Message
}

//Send sends a message to a fake receiver
func (t *TransportMock) Send(message transport.Message) bool {
	t.sentMessages = append(t.sentMessages, message)
	return true
}

//Receive receives fake message
func (t *TransportMock) Receive() transport.Message {
	message, x := t.receivableMessages[len(t.receivableMessages)-1], t.receivableMessages[:len(t.receivableMessages)-1]
	t.receivableMessages = x
	return message
}

//AddReceivableMessage add message which can be received with Receive
func (t *TransportMock) AddReceivableMessage(message transport.Message) {
	t.receivableMessages = append(t.receivableMessages, message)
}

//GetLastSentMessageAsString returns last sent message as string
func (t *TransportMock) GetLastSentMessageAsString() string {
	return t.sentMessages[len(t.sentMessages)-1].String()
}
