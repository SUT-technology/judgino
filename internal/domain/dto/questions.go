package dto

type QuestionRequest struct {
	UserId 	uint 
	SearchValue string `form:"searchFilter"`
	QuestionValue string `form:"questionFilter"`
	SortValue string `form:"sortFilter"`
	PageParam uint `form:"pageNumber"`
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
}