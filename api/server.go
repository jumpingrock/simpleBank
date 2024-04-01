package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/simpleBank/db/sqlc"
)

// Server this server will serve all HTTP request for our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	//to suppress debug mode for gin
	//gin.SetMode(gin.ReleaseMode)
	server := &Server{store: store}
	router := gin.Default()

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := v.RegisterValidation("currency", validCurrency)
		if err != nil {
			fmt.Println("Error registering currency validation:")
		}
	}

	router.POST("/users", server.createUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
