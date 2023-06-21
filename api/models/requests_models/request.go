package requests_models

// Response is sent as a response with information about the success of the request
type Response struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

const (
	StandardAnswerOnError = "Произошла ошибка при выполнении операции, попробуйте позже"
)