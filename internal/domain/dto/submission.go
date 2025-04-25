package dto

type SubmissionRequest struct {
	SubmissionValue string `json:"submissionFilter" query:"submissionFilter"`
	FinalValue      string `json:"finalFilter" query:"finalFilter"`
	PageParam       uint   `json:"pageInfo" query:"pageInfo"`
	PageAction      string `json:"page" query:"page"`
}

type Submission struct {
	QuestionTitle string `json:"QuestionTitle"`
	UserName      string `json:"UserName"`
	Status        int64  `json:"Status"`
	Date          string `json:"Date"`
	Type          string `json:"Type"`
}
type SubmissionsResponse struct {
	Submissions      []Submission `json:"Submissions"`
	TotalPages       int          `json:"TotalPages"`
	QuestionId       int          `json:"QuestionId"`
	CurrentPage      int          `json:"CurrentPage"`
	SubmissionFilter string       `json:"SubmissionFilter"`
	FinalFilter      string       `json:"FinalFilter"`
	Error            error        `json:Error`
}

type SubmissionRun struct {
	ID             uint   `json:"id"`
	Code           string `json:"code"`
	Input          string `json:"input"`
	ExpectedOutput string `json:"expectedOutput"`
	TimeLimit      int    `json:"timeLimit"`
	MemoryLimit    int    `json:"memoryLimit"`
}

type SubmissionRunResp struct {
	Submissions []SubmissionRun `json:"submissions"`
}
type SubmissionResult struct {
	ID     uint `json:"id"`
	Status int  `json:"status"`
}
