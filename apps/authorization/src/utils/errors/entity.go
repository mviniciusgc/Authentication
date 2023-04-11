package errors

type ErrorBody struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Op      string      `json:"op,omitempty"`
	Err     *ErrorValue `json:"err,omitempty"`
	Detail  string      `json:"detail,omitempty"`
}

type ErrorValue struct {
	Code    *int   `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
