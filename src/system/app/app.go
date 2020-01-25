package app

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gorilla/handlers"
	"github.com/sanket143/lisa/src/system/router"
)

// Server Model
type Server struct {
	port string
	Db   *dynamodb.DynamoDB
}

// NewServer Function
func NewServer() Server {
	return Server{}
}

// Init all values
func (s *Server) Init(port string, db *dynamodb.DynamoDB) {
	log.Println("Initializing Server...")
	s.port = ":" + port
	s.Db = db
}

// Start the
func (s *Server) Start() {
	log.Println("Starting server on port" + s.port)

	r := router.NewRouter()
	r.Init(s.Db)

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{""}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(r.Router))

	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	newServer := http.Server{
		Handler:      handler,
		Addr:         "0.0.0.0" + s.port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(newServer.ListenAndServe())
}
