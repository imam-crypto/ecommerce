package helpers

type DefaultResponse struct {
	Status     int         `json:"status"`
	Is_Success bool        `json:"is_success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
type ResponsePaginate struct {
	Status     int         `json:"status"`
	Is_Success bool        `json:"is_success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
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
func ConvResponsePaginate(Status int, Is_success bool, Message string, data interface{}, pagination interface{}) ResponsePaginate {

	Result := ResponsePaginate{
		Status:     Status,
		Is_Success: Is_success,
		Message:    Message,
		Data:       data,
		Pagination: pagination,
	}

	return Result
}
