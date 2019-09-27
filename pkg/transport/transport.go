package transport

// RequestParams are request parameters for Receive
type RequestParams interface {
	Hash() string
	Map() map[string]string
}

// EmptyRequestParams are empty request parameters
type EmptyRequestParams struct {}

// Hash returns hash representation of EmptyRequestParams
func (p EmptyRequestParams) Hash() string {
	return "";
}

// Map returns map representation of EmptyRequestParams
func (p EmptyRequestParams) Map() map[string]string {
	return map[string]string{}
}

//Transport is a message transport interface
type Transport interface {
	Send(Message) (bool, error)
	Receive(RequestParams) chan Message
}
