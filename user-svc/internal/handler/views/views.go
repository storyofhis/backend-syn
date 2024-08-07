package views

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

const (
	M_BAD_REQUEST           = "BAD_REQUEST"
	M_INVALID_CREDENTIALS   = "M_INVALID_CREDENTIALS"
	M_CREATED               = "CREATED"
	M_OK                    = "OK"
	M_INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	M_EMAIL_ALREADY_USED    = "EMAIL_ALREADY_USED"
)

func SuccessResponse(status int, message string, payload interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Payload: payload,
	}
}

func ErrorResponse(status int, message string, err error) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Error:   err.Error(),
	}
}
