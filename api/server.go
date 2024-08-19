package api

import (
	db "github.com/REZ-OAN/simplebank/database/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts/create", server.createAccount)
	router.GET("/accounts/get/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.PUT("/accounts/update/:id", server.updateAccount)
	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
