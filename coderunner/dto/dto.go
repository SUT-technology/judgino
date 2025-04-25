package dto

type SubmissionRun struct {
	ID             int    `json:"id"`
	Code           string `json:"code"`
	Input          string `json:"input"`
	ExpectedOutput string `json:"expectedOutput"`
	TimeLimit      int    `json:"timeLimit"`
	MemoryLimit    int    `json:"memoryLimit"`
}

type SubmissionRunResp struct {
	Submissions []SubmissionRun `json:"submissions"`
}
