package dto

type QuestionRequest struct {
	UserId 	uint 
	SearchValue string
	QuestionValue string
	SortValue string
	PageParam uint
}