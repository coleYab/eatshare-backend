package handler

import (
	"fmt"
	"net/http"

	"github.com/coleYab/erestourant/internal/dto"
	"github.com/coleYab/erestourant/internal/store"
	"github.com/coleYab/erestourant/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	us *store.UserStore
}

func NewAuthHandler(us *store.UserStore) *AuthHandler {
	return &AuthHandler{us: us}
}

func (a *AuthHandler) Login(ctx *gin.Context) {
	payload := dto.LoginDto{}
	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := a.us.GetUserByEmail(payload.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid credentials")
		return
	}

	// TODO: fix it later
	// if err := utils.ComparePassword(payload.Password, user.Password); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, "invalid credentials "+err.Error())
	// 	return
	// }
	if user.Password != payload.Password {
		ctx.JSON(http.StatusBadRequest, "invalid credentials")
		return
	}

	token, err := utils.CreateJwtToken(user.Email, user.ID.String())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "unable to create jwt token")
		return
	}

	// TODO: add jwt token
	ctx.JSON(200, gin.H{"message": "login successfull", "token": token})
}

func (a *AuthHandler) Register(ctx *gin.Context) {
	payload := dto.RegisterDto{}
	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := a.us.CreateUser(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("unable to create user %v", err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"message": "user created successfully", "user": user})
}
