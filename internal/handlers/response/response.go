package response

type Message struct {
	Success bool          `json:"success"`
	Error   *ErrorMessage `json:"error,omitempty"`
	Payload interface{}   `json:"payload,omitempty"`
}

type ErrorMessage struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

var MessageOK = Message{
	Success: true,
}

func CreateErrorResponse(code int, err error) Message {
	message := Message{
		Success: false,
		Error: &ErrorMessage{
			Code:    code,
			Message: err.Error(),
		},
	}

	return message
}
