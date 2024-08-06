package request

type CreateStudentRequest struct {
	Name     string `validate:"required,min=1,max=200" json:"name"`
	Class    string `validate:"required,max=200,min=1" json:"class"`
	SchoolID int    `validate:"required,max=200,min=1" json:"school_id"`
}

type UpdateStudentRequest struct {
	Name     *string `json:"name"`
	Class    *string `json:"class"`
	SchoolID *int    `json:"school_id"`
}
