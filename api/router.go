package api

import "github.com/gin-gonic/gin"

func (server *Server) setupRouter() {
	router := gin.Default()

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts/create", server.createAccount)
	authRoutes.GET("/accounts/get/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)
	authRoutes.PUT("/accounts/update", server.updateAccount)
	authRoutes.POST("/transfer", server.createTransfer)
	router.POST("/users/create", server.createUser)
	authRoutes.GET("/users/get/:username", server.getUser)
	router.POST("/users/login", server.loginUser)
	server.router = router

}
