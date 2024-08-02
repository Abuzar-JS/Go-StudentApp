package request

type CreateSchoolRequest struct {
	Name string `validate:"required,min=1,max=200" json:"name"`
}

type UpdateSchoolRequest struct {
	Id   int    `validate:"required" json:"id"`
	Name string `validate:"required,max = 200, min =1" json:"name"`
}
