package testing

import (
	"fmt"
	"time"
	
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/transport"
)

// TransportMock is a mock for transport layer
type TransportMock struct {
	sentMessages       []transport.Message
	receivableMessages map[string][]transport.Message
}

// NewTransportMock is a constructor for TransportMock
func NewTransportMock() *TransportMock {
	return &TransportMock{}
}

// Send sends a message to a fake receiver
func (t *TransportMock) Send(message transport.Message) (bool, error) {
	t.sentMessages = append(t.sentMessages, message)
	return true, nil
}

// Receive receives fake message
func (t *TransportMock) Receive(params transport.RequestParams) (chan transport.Message, chan error) {
	messageChannel := make(chan transport.Message, 1)
	errorChannel := make(chan error, 1)
	hash := params.Hash()
	messages := t.receivableMessages[hash]
	len := len(messages)
	if len == 0 {
		errorChannel <- fmt.Errorf("no messages to receive for %v", hash)
		close(messageChannel)
		close(errorChannel)
		return nil, errorChannel
	}
	message, x := messages[len-1], messages[:len-1]
	t.receivableMessages[hash] = x
	go func(){
		time.Sleep(50 * time.Millisecond)
		messageChannel <- message
		close(messageChannel)
		close(errorChannel)
	}()
	return messageChannel, nil
}

// Close closes transport
func (t *TransportMock) Close() error {
	return nil
}

// AddReceivableMessage adds a message which can be received with
// Receive for specified request parameters
func (t *TransportMock) AddReceivableMessage(message transport.Message, params transport.RequestParams) {
	hash := params.Hash()
	if t.receivableMessages == nil {
		t.receivableMessages = make(map[string][]transport.Message)
	}
	t.receivableMessages[hash] = append(t.receivableMessages[hash], message)
}

// GetLastSentMessageAsString returns last sent message as string
func (t *TransportMock) GetLastSentMessageAsString() (string, bool) {
	if len(t.sentMessages) == 0 {
		return "", false
	}
	return t.sentMessages[len(t.sentMessages)-1].String(), true
}

// GetLastSentMessage returns last sent message
func (t *TransportMock) GetLastSentMessage() (*transport.Message, bool) {
	if len(t.sentMessages) == 0 {
		return nil, false
	}
	return &t.sentMessages[len(t.sentMessages)-1], true
}
