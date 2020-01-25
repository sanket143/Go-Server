package app

import (
	"log"
	"net/http"

	"github.com/sanket143/lisa/src/system/router"
)

// Server Model
type Server struct {
	port string
}

// NewServer Function
func NewServer() Server {
	return Server{}
}

// Init all values
func (s *Server) Init() {
	log.Println("Initializing Server...")
	s.port = ":8000"
}

// Start the
func (s *Server) Start() {
	log.Println("Starting server!")

	r := router.NewRouter()
	r.Init()

	http.ListenAndServe(s.port, r.Router)
}
