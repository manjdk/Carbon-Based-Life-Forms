package error

type Message struct {
	Error string `json:"error"`
}

func NewErrorMessage(err error) *Message {
	return &Message{
		Error: err.Error(),
	}
}
