package htmltmp

import (
	"net/http"
	"strings"

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

func (h *ProfileHndlr) HandleDynamicProfile(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from the URL
	path := r.URL.Path // e.g., /profile/soroush
	userID := strings.TrimPrefix(path, "/profile/")
	if userID == "" {
		http.NotFound(w, r)
		return
	}

	// Example: fetch user profile using userID (just return it for now)
	w.Write([]byte("Showing profile for user: " + userID))
}
