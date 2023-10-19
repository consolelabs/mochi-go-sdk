package model

type ErrorMessage struct {
	Err        string `json:"err"`
	Msg        string `json:"msg"`
	StatusCode int    `json:"status_code"`
	ErrorCode  string `json:"error_code_code"`
}
