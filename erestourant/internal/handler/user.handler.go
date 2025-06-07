package handler

import (
	"net/http"

	"github.com/coleYab/erestourant/internal/db/repository"
	"github.com/coleYab/erestourant/internal/store"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	us *store.UserStore
}

func NewUserHandler(us *store.UserStore) *UserHandler {
	return &UserHandler{us: us}
}

func (a *UserHandler) GetDetail(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := a.us.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user found",
		"user":    user,
	})
}

func (a *UserHandler) GetAll(ctx *gin.Context) {
	users, err := a.us.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to get all users",
			"error":   err.Error(),
		})
		return
	}

	if users == nil {
		users = []repository.User{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"users":   users,
	})
}

func (a *UserHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := a.us.DeleteUser(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to delete user",
			"error":   err.Error(),
		})
	}

	ctx.JSON(http.StatusNoContent, nil)
}
