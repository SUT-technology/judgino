package htmltmp

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"html/template"

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

	return handler
}

type ProfileData struct {
	Username string
	Phone string
	Email string
	Role string
	Attempted int64
	Accepted int64
	total int64
	err error
}

func (h *ProfileHndlr) HandleDynamicProfile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	userIDstr , _ := strconv.ParseUint(strings.TrimPrefix(path, "/profile/"),10,64)
	userId := uint(userIDstr)
	profileDto := dto.ProfileDTO{
		UserId: userId,
	}

	user,err:=h.Services.PrflSrvc.GetProfileById(context.Background(),profileDto)

	data := ProfileData {
		Username: user.Username,
		Phone: user.Phone,
		Email: user.Email,
		Role: user.Role,
		Accepted: user
		err: err,
	}


	tmpl := template.Must(template.ParseFiles("D:/GOprojects/practice/judgino/templates/profile.html"))
	tmpl.Execute(w,data)
}
