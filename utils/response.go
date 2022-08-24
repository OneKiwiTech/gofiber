package utils

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(message string, data interface{}) Response {
	res := Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, data interface{}) Response {
	res := Response{
		Status:  false,
		Message: message,
		Data:    data,
	}
	return res
}
