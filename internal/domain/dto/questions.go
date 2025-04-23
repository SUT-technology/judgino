package dto

type QuestionRequest struct {
	SearchFilter  string `json:"SearchFilter" query:"SearchFilter"`
	QuestionValue string `json:"questionFilter" query:"questionFilter"`
	SortValue     string `json:"sortFilter" query:"sortFilter"`
	PageParam     int    `json:"pageInfo" query:"pageInfo"`
	PageAction    string `json:"page" query:"page"`
}
type Question struct {
	Title       string `json:"title"`
	PublishDate string `json:"publish_date"`
	Deadline    string `json:"deadline"`
}
type QuestionsResponse struct {
	Questions      []Question `json:"Questions"`
	TotalPages     int        `json:"TotalPages"`
	CurrentPage    int        `json:"CurrentPage"`
	SearchFilter   string     `json:"SearchFilter"`
	QuestionFilter string     `json:"questionFilter"`
	SortFilter     string     `json:"sortFilter"`
	Error          error      `json:Error`
}

type PublishResponse struct {
	Msg string
}
