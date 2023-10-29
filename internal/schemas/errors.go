package schemas

type ErrorResponse struct {
	/* Схема ответа с ошибкой */
	error string
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		error: err.Error(),
	}
}
