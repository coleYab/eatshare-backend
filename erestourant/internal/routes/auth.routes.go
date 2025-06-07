package routes

import (
	"github.com/coleYab/erestourant/internal/handler"
	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	e *gin.Engine
}

func NewAuthRoute(e *gin.Engine) *AuthRoutes {
	return &AuthRoutes{e: e}
}

func (r *AuthRoutes) RegisterRoutes(authHandler *handler.AuthHandler) {
	r.e.POST("/auth/login", authHandler.Login)
	r.e.POST("/auth/register", authHandler.Register)
}
