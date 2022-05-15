package jsonapi

//

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (resp *ErrorResponse) SetContent(code int, mess string) {
	resp.Code = code
	resp.Message = mess
}

func NewErrorResponse() Response {
	return new(ErrorResponse)
}

//

type ErrorResponseShort struct {
	Code int `json:"code"`
}

func (resp *ErrorResponseShort) SetContent(code int, mess string) {
	resp.Code = code
}

func NewErrorResponseShort() Response {
	return new(ErrorResponseShort)
}
