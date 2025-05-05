package server

import (
	"html/template"
	"io/fs"
	"log/slog"
	"mime"
	"net/http"

	"github.com/way-platform/rfms-go/cmd/rfms/v4/internal/dashboard/ui"
	"github.com/way-platform/rfms-go/v4"
)

// Server is an rFMS dashboard server.
type Server struct {
	client        *rfms.Client
	templates     *template.Template
	staticHandler http.Handler
	mux           *http.ServeMux
}

// NewServer creates a new rFMS dashboard server.
func NewServer(client *rfms.Client) (*Server, error) {
	if err := mime.AddExtensionType(".js", "text/javascript"); err != nil {
		return nil, err
	}
	templates, err := template.ParseFS(ui.TemplateFS, "templates/*/*.html")
	if err != nil {
		return nil, err
	}
	rootedStaticFS, err := fs.Sub(ui.StaticFS, "static")
	if err != nil {
		return nil, err
	}
	server := &Server{
		client:        client,
		templates:     templates,
		staticHandler: http.StripPrefix("/static/", http.FileServer(http.FS(rootedStaticFS))),
		mux:           http.NewServeMux(),
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
		s.handleError(w, r, "failed to render template", err)
	}
}

func (s *Server) handleError(w http.ResponseWriter, r *http.Request, message string, err error) {
	slog.ErrorContext(r.Context(), message, "error", err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
