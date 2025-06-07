package handler

import (
	"net/http"

	"github.com/coleYab/erestourant/internal/db/repository"
	"github.com/coleYab/erestourant/internal/store"
	"github.com/coleYab/erestourant/internal/utils"
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
		utils.RespondError(ctx, http.StatusNotFound, "user_not_found", "User not found", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User found",
		"user":    user,
	})
}

func (a *UserHandler) GetAll(ctx *gin.Context) {
	users, err := a.us.GetAllUsers()
	if err != nil {
		utils.RespondError(ctx, http.StatusInternalServerError, "user_fetch_failed", "Unable to retrieve users", err.Error())
		return
	}

	if users == nil {
		users = []repository.User{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"users":   users,
	})
}

func (a *UserHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := a.us.DeleteUser(id); err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, "user_deletion_failed", "Unable to delete user", err.Error())
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
