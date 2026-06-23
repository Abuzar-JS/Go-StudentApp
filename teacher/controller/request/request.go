package request

type CreateTeacherRequest struct {
	Name     string `validate:"required,min=1,max=200" json:"name"`
	Subject  string `validate:"required,min=1,max=200" json:"subject"`
	SchoolID int    `json:"school_id"`
}

type UpdateTeacherRequest struct {
	Name     *string `json:"name"`
	Subject  *string `json:"subject"`
	SchoolID *int    `json:"school_id"`
}
