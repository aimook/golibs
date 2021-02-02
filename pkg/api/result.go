package api

const (
	SUCCESS = 10000
)

type Result struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

func ResultCreate() *Result {
	r := defaultSuccessResult()
	return r
}

func defaultSuccessResult() *Result {
	r := &Result{
		Success: true,
		Message: "",
		Code:    SUCCESS,
	}
	return r
}

type Option func(r *Result)

func (r *Result) WithCode(code int) *Result {
	r.Code = code
	return r
}

func (r *Result) WithMessage(message string) *Result {
	r.Message = message
	return r
}

func (r *Result) WithData(data interface{}) *Result {
	r.Data = data
	return r
}

func (r *Result) WithSuccess() *Result {
	r.Code = SUCCESS
	r.Success = true
	return r
}

func (r *Result) WithFail() *Result {
	r.Success = false
	return r
}
