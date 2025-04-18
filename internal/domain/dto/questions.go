package dto

type QuestionsDto struct {
	UserId 	uint 
	SearchValue string
	QuestionValue string
	SortValue string
	PageParam uint
}