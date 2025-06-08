package handler

import (
	"net/http"

	"github.com/coleYab/erestourant/internal/dto"
	"github.com/coleYab/erestourant/internal/store"
	"github.com/coleYab/erestourant/internal/utils"
	"github.com/gin-gonic/gin"
)

type RecipeHandler struct {
	us *store.UserStore
}

func NewRecipeHandler(us *store.UserStore) *RecipeHandler {
	return &RecipeHandler{us: us}
}

func (a *RecipeHandler) Login(ctx *gin.Context) {
	var payload dto.LoginDto
	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid_request",
			"message": "Invalid request payload.",
		})
		return
	}

	user, err := a.us.GetUserByEmail(payload.Email)
	if err != nil || user.Password != payload.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "invalid_credentials",
			"message": "Invalid email or password.",
		})
		return
	}

	token, err := utils.CreateBasicToken(user.Email, user.ID.String())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "token_generation_failed",
			"message": "Failed to create authentication token.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id":       user.ID,
			"username": "username", // You might want to fix this to actually use user.Username
			"email":    user.Email,
			"name":     user.Name,
		},
		"token": token,
	})
}

func (a *RecipeHandler) Register(ctx *gin.Context) {
	var payload dto.RegisterDto
	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid_request",
			"message": "Invalid registration data.",
		})
		return
	}

	if _, err := a.us.GetUserByEmail(payload.Email); err == nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"error":   "user_exists",
			"message": "User already exists with this email.",
		})
		return
	}

	user, err := a.us.CreateUser(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "user_creation_failed",
			"message": "Failed to create user.",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}
