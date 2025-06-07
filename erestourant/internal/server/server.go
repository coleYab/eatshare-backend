package server

import (
	"github.com/coleYab/erestourant/internal/db/repository"
	"github.com/coleYab/erestourant/internal/handler"
	"github.com/coleYab/erestourant/internal/routes"
	"github.com/coleYab/erestourant/internal/store"
	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	r    *gin.Engine
	addr string
	qry  *repository.Queries
}

func New(addr string, qry *repository.Queries) *ApiServer {
	return &ApiServer{
		r:    gin.Default(),
		addr: addr,
		qry:  qry,
	}
}

func (s *ApiServer) RegisterRoutes() {
	userStore := store.NewUserStore(s.qry)

	// /auth - > subroute
	authRoutes := routes.NewAuthRoute(s.r)
	authHandler := handler.NewAuthHandler(userStore)
	authRoutes.RegisterRoutes(authHandler)

	// /user - > subroute
	userRoutes := routes.NewUserRoute(s.r)
	UserHandler := handler.NewUserHandler(userStore)
	userRoutes.RegisterRoutes(UserHandler)
}

func (s *ApiServer) Run() {
	s.r.Run(s.addr)
}
