package dto

type SubmissionRequest struct {
	UserId 	uint 
	IsAdmin bool
	QuestionId int
	SubmissionValue string `form:"submissionFilter"`
	FinalValue string `form:"finalFilter"`
	PageParam uint `form:"pageNumber"`
	PageAction string `form:"page"`
}

type Submission struct {
	QuestionTitle string `json:"QuestionTitle"`
	UserName string `json:"UserName"`
	Status int64 `json:"Status"`
	Date string `json:"Date"`
	Type string `json:"Type"`
}
type SubmissionsResponse struct {
	Submissions []Submission `json:"Submissions"`
	TotalPages  int          `json:"TotalPages"`
	QuestionId int `json:"QuestionId"`
	CurrentPage int `json:"CurrentPage"`
}