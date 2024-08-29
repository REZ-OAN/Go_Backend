package api

import (
	"fmt"

	db "github.com/REZ-OAN/simplebank/database/sqlc"
	"github.com/REZ-OAN/simplebank/utils"
	"github.com/go-playground/validator/v10"

	token "github.com/REZ-OAN/simplebank/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Server struct {
	store      db.Store
	tokenMaker token.Maker
	config     utils.Config
	router     *gin.Engine
}

func NewServer(store db.Store, config utils.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TOKEN_SYM_KEY)
	if err != nil {
		return nil, fmt.Errorf("cannot make token : %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
