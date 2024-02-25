package response

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOk    = "200"
	StatusError = "500"
)

func Ok() Response {
	return Response{
		Status: StatusOk,
	}
}

func Error(errorMessage string) Response {
	return Response{
		Status: StatusError,
		Error:  errorMessage,
	}
}
