package routes

import (
	"github.com/coleYab/erestourant/internal/handler"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	e *gin.Engine
}

func NewUserRoute(e *gin.Engine) *UserRoutes {
	return &UserRoutes{e: e}
}

func (r *UserRoutes) RegisterRoutes(userHandler *handler.UserHandler) {
	r.e.GET("/user/", userHandler.GetAll)
	r.e.GET("/user/:id", userHandler.GetDetail)
	r.e.DELETE("/user/:id", userHandler.DeleteUser)
}
