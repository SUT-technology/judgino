package htmltmp

import (
	// "encoding/json"
	"html/template"
	"net/http"
	"strconv"

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

	g.Handle("GET", "/", handler.ShowData)

	return handler
}

type Question struct {
	Title string `json:"title"`
	PublishDate string `json:"publish_date"`
	Deadline string `json:"deadline"`
}

func (q *QuestionsHndlr) ShowData(w http.ResponseWriter, r *http.Request) {
	// Retrieve query parameters from the URL
	var userId int
	if r.Header.Get("Authorization") == "" {userId = 0} else {userId, _ = strconv.Atoi(r.Context().Value("userId").(string))}

	searchValue := r.URL.Query().Get("searchValue")
	questionValue := r.URL.Query().Get("questionValue")
	sortValue := r.URL.Query().Get("sortValue")
	pageParam, _ := strconv.Atoi(r.URL.Query().Get("page"))
	// fmt.Println(r.Context().Value("page"))


	// searchValue := ""
	// questionValue := "all"
	// sortValue := "deadline"
	// pageParam := 1
	
	questionsDto := dto.QuestionsDto{
		UserId:       uint(userId),
		SearchValue:  searchValue,
		QuestionValue: questionValue,
		SortValue:    sortValue,
		PageParam: uint(pageParam),
	}

	questions, err := q.Services.QuestionsSrvc.GetDate(r.Context(), questionsDto)
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


	data := map[string]interface{}{
		"questions":  questions,
		"totalPages": 10,
	}
	// jsonData, err := json.Marshal(data)
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	
}