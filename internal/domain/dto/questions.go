package dto

import (
	"github.com/SUT-technology/judgino/internal/domain/entity"
	"github.com/SUT-technology/judgino/internal/domain/model"
)

type QuestionSummeryRequest struct{
	SearchFilter  string `json:"SearchFilter" query:"SearchFilter"`
	QuestionValue string `json:"questionFilter" query:"questionFilter"`
	SortValue     string `json:"sortFilter" query:"sortFilter"`
	PageParam     int    `json:"pageInfo" query:"pageInfo"`
	PageAction    string `json:"page" query:"page"`
}
type QuestionSummery struct {
	Title string 	   `json:"title"`
	PublishDate string `json:"publish_date"`
	Deadline    string `json:"deadline"`
	QuestionId int64   `json:"questionId"`
	Status string	   `json:"published"`
	Publisher string   `json:"publisher"`
	PublisherId int64  `json:"publisherID"`
	IsCurrentUserAdmin bool `json:"IsCurrentUserAdmin"`
}
type QuestionsSummeryResponse struct {
	Questions      []QuestionSummery `json:"Questions"`
	CurrentUserId int64
	TotalPages     int        		 `json:"TotalPages"`
	CurrentPage    int        		 `json:"CurrentPage"`
	SearchFilter   string     		 `json:"SearchFilter"`
	QuestionFilter string     		 `json:"questionFilter"`
	SortFilter     string     		 `json:"sortFilter"`
	IsCurrentUserAdmin bool	  		 `json:"isCurrentUserAdmin"`
	Error          error      		 `json:Error`
}
type CreateQuestionRequest struct {
	UserID  int64 	  `json:"UserID"`
	Status   string   `json:"Status"`
	Title  string     `json:"Title"`
	Body    string    `json:"Body"`
	TimeLimit int64   `json:"TimeLimit"`
	MemoryLimit int64 `json:"MemoryLimit"`
	InputURL  string  `json:"InputURL"`
	Deadline  string  `json:"Deadline"`
	OutputURL string  `json:"OutputURL"`
}

type CreateQuestionResponse struct {
	Error        bool
	Status       model.UserMessage
	QuestionID   int64
	UserID       int64
	Title        bool
	Body         bool
	TimeLimit    bool
	MemoryLimit  bool
	InputURL     bool
	OutputURL    bool
}

type GetQuestionResponse struct {
	Question *entity.Question
	CurrentUserId int64
}


type PublishResponse struct {
	Msg string
}

