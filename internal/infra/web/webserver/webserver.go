package webserver

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type typeOfRequest struct {
	Method, Path string
}

type WebServer struct {
	Router        chi.Router
	Handlers      map[typeOfRequest]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[typeOfRequest]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	s.Handlers[typeOfRequest{
		Method: method,
		Path:   path,
	}] = handler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for request, handler := range s.Handlers {

		switch request.Method {
		case http.MethodGet:
			s.Router.Get(request.Path, handler)
		case http.MethodPost:
			s.Router.Post(request.Path, handler)
		case http.MethodPut:
			s.Router.Put(request.Path, handler)
		case http.MethodDelete:
			s.Router.Delete(request.Path, handler)
		// Adicione outros métodos aqui, se necessário
		default:
			panic("Método HTTP não suportado: " + request.Method)
		}
	}

	// Serve static files
	staticDir := "./static"
	FileServer(s.Router, "/static", http.Dir(staticDir))

	// Default route
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, staticDir+"/index.html")
	})

	err := http.ListenAndServe(s.WebServerPort, s.Router)
	if err != nil {
		panic(err)
	}
}

// FileServer is a utility for serving static files
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}
	fs := http.StripPrefix(path, http.FileServer(root))
	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}
