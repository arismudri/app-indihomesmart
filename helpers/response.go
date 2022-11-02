package helpers

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r Response) Success() Response {
	r.Code = 200
	r.Message = "Success"

	return r
}

func (r Response) Fail() Response {
	if r.Code == 0 {
		r.Code = 400
	}

	return r
}
