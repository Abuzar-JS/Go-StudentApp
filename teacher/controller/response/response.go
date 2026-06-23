package response

type TeacherResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Subject  string `json:"subject"`
	SchoolID int    `json:"school_id"`
}
