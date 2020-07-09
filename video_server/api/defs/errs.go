package defs

// Err 错误码
type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

// ErroResponse 响应
type ErroResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErroResponse{
		HttpSC: 400,
		Error: Err{
			Error:     "Request is not body",
			ErrorCode: "001",
		},
	}
	ErrorNotAuthUser = ErroResponse{
		HttpSC: 401,
		Error: Err{
			Error:     "user authenticat",
			ErrorCode: "002",
		},
	}
)
