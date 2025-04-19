package dto

type QuestionRequest struct {
	UserId 	uint 
	SearchFilter string `form:"SearchFilter"`
	QuestionValue string `form:"questionFilter"`
	SortValue string `form:"sortFilter"`
	PageParam int `form:"pageInfo"`
	PageAction string `form:"page"`
}
type Question struct {
	Title string `json:"title"`
	PublishDate string `json:"publish_date"`
	Deadline string `json:"deadline"`
}
type QuestionsResponse struct {
	Questions []Question `json:"Questions"`
	TotalPages  int          `json:"TotalPages"`
	CurrentPage int `json:"CurrentPage"`
	SearchFilter string `json:"SearchFilter"`
	QuestionFilter string `json:"questionFilter"`
	SortFilter string `json:"sortFilter"`
}