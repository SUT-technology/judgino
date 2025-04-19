package htmltmp

import (
	// "encoding/json"
	"fmt"
	// "fmt"
	"html/template"
	"net/http"
	"strconv"

	// "github.com/goccy/go-json"
	"github.com/gorilla/mux"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/service"
)

type QuestionsHndlr struct {
	Services service.Service
}

func NewQuestionsHndlr(g *Group, srvc service.Service) QuestionsHndlr {
	handler := QuestionsHndlr{
		Services: srvc,
	}

	g.Handle("GET", "/", handler.ShowQuestions)

	//TODO fix
	g.Handle("GET", "/{question_id}/submissions", handler.ShowSubmissions)

	return handler
}

type Question struct {
	Title string `json:"title"`
	PublishDate string `json:"publish_date"`
	Deadline string `json:"deadline"`
}
type QuestionsResponse struct {
	Questions []Question `json:"questions"`
	TotalPages  int          `json:"totalPages"`
}

func (q *QuestionsHndlr) ShowQuestions(w http.ResponseWriter, r *http.Request) {
	// Retrieve query parameters from the URL
	var userId int
	if r.Header.Get("Authorization") == "" {userId = 0} else {userId, _ = strconv.Atoi(r.Context().Value("userId").(string))}
	

	searchValue := r.URL.Query().Get("searchValue")
	questionValue := r.URL.Query().Get("questionValue")
	sortValue := r.URL.Query().Get("sortValue")
	pageParam, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if pageParam == 0 {
		pageParam = 1
	}
	// fmt.Println(r.Context().Value("page"))


	// searchValue := ""
	// questionValue := "all"
	// sortValue := "deadline"
	// pageParam := 1
	
	questionsDto := dto.QuestionRequest{
		UserId:       uint(userId),
		SearchValue:  searchValue,
		QuestionValue: questionValue,
		SortValue:    sortValue,
		PageParam: uint(pageParam),
	}

	questions, err := q.Services.QuestionsSrvc.GetQuestions(r.Context(), questionsDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	

	// Create the data to pass to the template
	questionsData := make([]Question, len(questions))
	for i, question := range questions {
		questionsData[i] = Question{
			Title:         question.Title,
			PublishDate: question.PublishDate.Format("2006-01-02 15:04:05"),
			Deadline: 	question.Deadline.Format("2006-01-02 15:04:05"),
		}
	}

	// Render the template with the data
	// tmpl := template.Must(template.New("questions").ParseFiles("templates/questions.html"))

	
	tmpl, err := template.ParseFiles("templates/questions.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	data := QuestionsResponse{
		Questions:  questionsData,
		TotalPages: 10,
	}
	if searchValue != "" || questionValue != "" || sortValue != "" || pageParam != 1 {
		fmt.Println("hello")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	
	
	
}

type Submission struct {
	QuestionTitle string `json:"question_title"`
	UserName string `json:"user_name"`
	Status int64 `json:"status"`
	Date string `json:"date"`
	Type string `json:"type"`
}
type SubmissionsResponse struct {
	Submissions []Submission `json:"submissions"`
	TotalPages  int          `json:"totalPages"`
}

func (q *QuestionsHndlr) ShowSubmissions(w http.ResponseWriter, r *http.Request) {
	var userId int
	if r.Header.Get("Authorization") == "" {userId = 0} else {userId, _ = strconv.Atoi(r.Context().Value("userId").(string))}

	vars := mux.Vars(r)
	questionID, _ := strconv.Atoi(vars["question_id"])
	submissonValue := r.URL.Query().Get("submissionValue")
	finalValue := r.URL.Query().Get("finalValue")
	pageParam, _ := strconv.Atoi(r.URL.Query().Get("page"))

	submissionsDto := dto.SubmissionRequest{
		UserId:         uint(userId),
		QuestionId:     uint(questionID),
		SubmissonValue: submissonValue,
		FinalValue:     finalValue,
		PageParam:      uint(pageParam),
	}

	submissions, err := q.Services.QuestionsSrvc.GetSubmissions(r.Context(), submissionsDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	

	// Create the data to pass to the template
	submissionsData := make([]Submission, len(submissions))
	for i, submission := range submissions {
		qt, err := q.Services.QuestionsSrvc.GetQuestion(r.Context(), submission.QuestionID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ut, err := q.Services.UserSrvc.GetUser(r.Context(), submission.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var typ string
		if submission.IsFinal {
			typ = "Final"
		} else {
			typ = "Normal"
		}
		submissionsData[i] = Submission{
			QuestionTitle: qt.Title,
			UserName: ut.FirstName,
			Status: submission.Status,
			Date: submission.SubmitTime.Format("2006-01-02 15:04:05"),
			Type: typ,
		}
	}
	

	tmpl, err := template.ParseFiles("templates/submissions.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	data := SubmissionsResponse{
		Submissions:  submissionsData,
		TotalPages: 10,
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}