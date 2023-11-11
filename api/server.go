package api

import (
	db "github.com/LKarrie/learn-go-project/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server servers HTTP requests for our banking service.
type Server struct {
	store db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v,ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency",validCurrency)
	}

	router.POST("/accounts",server.createAccount)
	router.GET("/accounts/:id",server.getAccount)
	router.GET("/accounts",server.listAccount)

	router.POST("/transfers",server.createTransfer)

	// add routers to router
	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error":err.Error()}
}