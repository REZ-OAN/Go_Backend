package api

import "github.com/gin-gonic/gin"

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/accounts/create", server.createAccount)
	router.GET("/accounts/get/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.PUT("/accounts/update/:id", server.updateAccount)
	router.POST("/transfer", server.createTransfer)
	router.POST("/users/create", server.createUser)
	router.GET("/users/get/:username", server.getUser)
	router.POST("/users/login", server.loginUser)
	server.router = router

}
