package helpers

type DefaultResponse struct {
	Status     int         `json:"status"`
	Is_Success bool        `json:"is_success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func ConvDefaultResponse(Status int, Is_success bool, Message string, data interface{}) DefaultResponse {

	Result := DefaultResponse{
		Status:     Status,
		Is_Success: Is_success,
		Message:    Message,
		Data:       data,
	}

	return Result
}
