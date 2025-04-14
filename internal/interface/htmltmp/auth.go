package htmltmp

import (
	"context"
	"html/template"
	"net/http"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/service"
)

type AuthHndlr struct {
	Services service.Service
}

func NewAuthHndlr(g *Group, srvc service.Service) AuthHndlr {
	handler := AuthHndlr{
		Services: srvc,
	}

	g.Handle("GET", "/login", handler.Login)

	return handler
}

type LoginData struct {
	FirstName string
	err       error
	UserId    int
}

func (h *AuthHndlr) Login(w http.ResponseWriter, r *http.Request) {

	loginDto := dto.LoginDTO{
		Username: "test",
		Password: "sa",
	}
	userId := r.Context().Value("userId").(int)
	firstname, err := h.Services.AuthSrvc.Login(context.Background(), loginDto)
	tmpl := template.Must(template.New("step1").Parse(`
        <html>
        <body>
            <h1>first name: {{.FirstName}}</h1>
            <h1>userId: {{.UserId}}</h1>
            <h1>error {{.err}}</h1>
        </body>
        </html>`))

	// Create the data to pass to the template
	data := LoginData{
		FirstName: firstname, // Replace "John" with dynamic data as needed
		err:       err,
		UserId:    userId,
	}

	// Execute the template with the data
	tmpl.Execute(w, data)
}
