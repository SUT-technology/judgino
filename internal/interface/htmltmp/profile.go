package htmltmp

import (
	"context"
	"encoding/json"
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


func (h *ProfileHndlr) HandleChangeRole(w http.ResponseWriter, r *http.Request) {

	var data dto.ChangeRoleDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	h.Services.PrflSrvc.ChangeRole(context.Background(),data)
}


func (h *ProfileHndlr) HandleDynamicProfile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	userID64 , _ := strconv.ParseUint(strings.TrimPrefix(path, "/profile/"),10,64)
	userId := uint(userID64)
	profileDto := dto.ProfileDTO{
		UserId: userId,
	}
	profile,err:=h.Services.PrflSrvc.GetProfileById(r.Context(),profileDto)
	if err != nil {
		http.Error(w, "profile not found", http.StatusNotFound)
		return
	}


	tmpl := template.Must(template.New("profile.html").Funcs(template.FuncMap{
		"eqs": func(a, b string) bool {
			return a == b
		},
		"equi": func(a, b uint) bool {
			return a == b
		},
	}).ParseFiles("D:/GOprojects/practice/judgino/templates/profile.html"))
	tmpl.Execute(w,profile)
}
