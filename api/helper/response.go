package helper

type Response struct {
	Status  int        `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type EmptyObj struct{}

func BuildResponse(status int, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Data:    data,
		Message: message,
	}
	return res
}

func BuildErrorResponse(status int,message string, data interface{}) Response {
	//splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  status,
		Data:    data,
		Message: message,
	}
	return res
}
