package api

import (
	"github.com/gin-gonic/gin"
	"github.com/senn404/bookmark-managent/internal/handler"
	"github.com/senn404/bookmark-managent/internal/service"
)

type Engine interface {
	Start() error
}

type api struct {
	app *gin.Engine
}

func New() Engine {
	a := &api{
		app: gin.New(),
	}
	a.registerEP()
	return a
}

func (a *api) Start() error {
	return a.app.Run(":8080")
}

func (a *api) registerEP() {
	passSvc := service.NewPassword()
	passHandler := handler.NewPassword(passSvc)
	a.app.GET("/gen-pass", passHandler.GenPass)
}
