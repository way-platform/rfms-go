package server

import (
	"html/template"
	"log/slog"
	"net/http"

	"github.com/way-platform/rfms-go/cmd/rfms/v4/internal/dashboard/ui"
)

// Server is an rFMS dashboard server.
type Server struct {
	templates *template.Template
	mux       *http.ServeMux
}

// NewServer creates a new rFMS dashboard server.
func NewServer() (*Server, error) {
	templates, err := template.ParseFS(ui.TemplateFS, "templates/**/*.tmpl")
	if err != nil {
		return nil, err
	}
	server := &Server{
		templates: templates,
		mux:       http.NewServeMux(),
	}
	server.registerRoutes(server.mux)
	return server, nil
}

// ServeHTTP implements [http.Handler].
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) renderTemplate(w http.ResponseWriter, r *http.Request, name string, data any) {
	if err := s.templates.ExecuteTemplate(w, name, data); err != nil {
		slog.ErrorContext(r.Context(), "failed to render template", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
