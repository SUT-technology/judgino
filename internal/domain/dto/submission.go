package dto

type SubmissionRequest struct {
	UserId 	uint 
	QuestionId uint
	SubmissonValue string
	FinalValue string
	PageParam uint
}