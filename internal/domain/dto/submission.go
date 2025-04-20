package dto

type SubmissionRequest struct {
	SubmissionValue string `json:"submissionFilter" query:"submissionFilter"`
	FinalValue string `json:"finalFilter" query:"finalFilter"`
	PageParam uint `json:"pageInfo" query:"pageInfo"`
	PageAction string `json:"page" query:"page"`
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
	SubmissionFilter string `json:"SubmissionFilter"`
	FinalFilter string `json:"FinalFilter"`
	Error error `json:Error`
}