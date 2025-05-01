package server

import (
	"net/http"

	"github.com/way-platform/rfms-go/cmd/rfms/v4/internal/dashboard/ui"
)

func (s *Server) registerRoutes(mux *http.ServeMux) {
	mux.Handle(s.staticRoute())
	mux.Handle(s.indexRoute())
}

func (s *Server) staticRoute() (string, http.Handler) {
	return "/static/", http.StripPrefix("/static/", http.FileServer(http.FS(ui.StaticFS)))
}

func (s *Server) indexRoute() (string, http.Handler) {
	return "/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.renderTemplate(w, r, "index.page.tmpl", nil)
	})
}
