package htmltmp

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/service"
)

type ProfileHndlr struct {
	Services service.Service
}

func NewProfileHndlr(g *Group, srvc service.Service) ProfileHndlr {
	handler := ProfileHndlr{
		Services: srvc,
	}

	g.Handle("GET", "/", handler.HandleDynamicProfile)
	g.Handle("POST", "/change-role", handler.HandleChangeRole)



	return handler
}

type ProfileData struct {
	UserId    uint
	CurrentUserId uint
	Username string
	Phone string
	Email string
	Role string
	NotAccepted int64
	Accepted int64
	Total int64
	SolvedPercentage int
	IsCurrentUserAdmin bool
	err error
}

func (h *ProfileHndlr) HandleChangeRole(w http.ResponseWriter, r *http.Request) {

	var data dto.UpdateUserDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	fmt.Println(data)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	h.Services.PrflSrvc.ChangeRole(context.Background(),data)
}


func (h *ProfileHndlr) HandleDynamicProfile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	currentUserProfileDto := dto.ProfileDTO{
		UserId: 1,
	}
	currentUser,_:=h.Services.PrflSrvc.GetProfileById(context.Background(),currentUserProfileDto)

	userID64 , _ := strconv.ParseUint(strings.TrimPrefix(path, "/profile/"),10,64)
	userId := uint(userID64)
	profileDto := dto.ProfileDTO{
		UserId: userId,
	}
	user,err:=h.Services.PrflSrvc.GetProfileById(context.Background(),profileDto)

	var p int

	if user.SubmissionsCount == 0 {
		p=0
	} else {
		p=100*int(user.SolvedQuestionsCount/user.SubmissionsCount)
	}

	data := ProfileData {
		UserId: userId,    
		CurrentUserId: currentUser.ID,         
		Username: user.Username,
		Phone: user.Phone,
		Email: user.Email,
		Role: user.Role,
		Accepted: user.SolvedQuestionsCount,
		NotAccepted: user.SubmissionsCount-user.SolvedQuestionsCount,
		Total: user.SubmissionsCount,
		SolvedPercentage:  p,
		IsCurrentUserAdmin: currentUser.Role=="admin",
		err: err,
	}


	tmpl := template.Must(template.New("profile.html").Funcs(template.FuncMap{
		"eqs": func(a, b string) bool {
			return a == b
		},
		"equi": func(a, b uint) bool {
			return a == b
		},
	}).ParseFiles("D:/GOprojects/practice/judgino/templates/profile.html"))
	tmpl.Execute(w,data)
}
