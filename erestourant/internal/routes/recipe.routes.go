package routes

import (
	"github.com/coleYab/erestourant/internal/handler"
	"github.com/gin-gonic/gin"
)

type RecipeRoutes struct {
	e *gin.Engine
}

func NewRecipeRoute(e *gin.Engine) *RecipeRoutes {
	return &RecipeRoutes{e: e}
}

func (r *RecipeRoutes) RegisterRoutes(recipeHandler *handler.RecipeHandler) {
	r.e.POST("recipe", recipeHandler.CreateRecipe)
	r.e.GET("recipe", recipeHandler.GetAll)
	r.e.GET("recipe/:id", recipeHandler.GetRecipeById)
	r.e.DELETE("recipe/:id", recipeHandler.DeleteRecipe)
}
