package dto

import (
	"time"

	"github.com/SUT-technology/judgino/internal/domain/model"
)

type QuestionSummeryRequest struct {
	SearchFilter  string `json:"SearchFilter" query:"SearchFilter"`
	QuestionValue string `json:"questionFilter" query:"questionFilter"`
	SortValue     string `json:"sortFilter" query:"sortFilter"`
	PageParam     int    `json:"pageInfo" query:"pageInfo"`
	PageAction    string `json:"page" query:"page"`
}
type QuestionSummery struct {
	Title       string `json:"title"`
	PublishDate string `json:"publish_date"`
	Deadline    string `json:"deadline"`
}
type QuestionsSummeryResponse struct {
	Questions      []QuestionSummery `json:"Questions"`
	TotalPages     int        `json:"TotalPages"`
	CurrentPage    int        `json:"CurrentPage"`
	SearchFilter   string     `json:"SearchFilter"`
	QuestionFilter string     `json:"questionFilter"`
	SortFilter     string     `json:"sortFilter"`
	Error          error      `json:Error`
}
type CreateQuestionRequest struct {
	UserID  int64 `json:"UserID"`
	Status   string`json:"Status"`
	Title  string `json:"Title"`
	Body    string `json:"Body"`
	TimeLimit int64 `json:"TimeLimit"`
	MemoryLimit int64 `json:"MemoryLimit"`
	InputURL  string `json:"InputURL"`
	Deadline  time.Time `json:"Deadline"`
	OutputURL string `json:"OutputURL"`
}

type CreateQuestionResponse struct {
	Error bool
	Status model.UserMessage
	questionId int64
	Title   bool
	Body bool
	TimeLimit bool
	MemoryLimit bool
	InputURL bool
	OutputURL bool
}