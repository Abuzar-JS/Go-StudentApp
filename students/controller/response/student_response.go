package response

type StudentResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
