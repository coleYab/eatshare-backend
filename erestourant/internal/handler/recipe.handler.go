package handler

import (
	"net/http"

	"github.com/coleYab/erestourant/internal/dto"
	"github.com/coleYab/erestourant/internal/store"
	"github.com/coleYab/erestourant/internal/utils"
	"github.com/gin-gonic/gin"
)

type RecipeHandler struct {
	rs *store.RecipeStore
}

func NewRecipeHandler(us *store.RecipeStore) *RecipeHandler {
	return &RecipeHandler{rs: us}
}

func (a *RecipeHandler) CreateRecipe(ctx *gin.Context) {
	var payload dto.CreateRecipeDto
	if err := ctx.ShouldBind(&payload); err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, "error_create_recipe", "invalid payload")
		return
	}

	recipe, err := a.rs.CreateRecipe(payload)
	if err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, "error_create_recipe", "unable to create the recipe")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "recipe created successfully",
		"recipe": recipe,
	})
}

func (a *RecipeHandler) GetRecipeById(ctx *gin.Context) {
	id := ctx.Param("id")
	recipe, err := a.rs.GetRecipeById(id)
	if err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, "delete_error", "unable to delete the recipe")
		return
	}

	steps, err := a.rs.GetStepsByRecipeId(id)
	if err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, "delete_error", "unable to delete the recipe")
		return
	}

	ingredients, err := a.rs.GetIngredientByRecipeId(id)
	if err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, "delete_error", "unable to delete the recipe")
		return
	}

	recipeDto := dto.GetRecipeDetails(recipe, steps, ingredients)

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "recipe": recipeDto})
}

func (a *RecipeHandler) DeleteRecipe(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := a.rs.DeleteRecipe(id); err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, "delete_error", "unable to delete the recipe")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (a *RecipeHandler) GetAll(ctx *gin.Context) {
	recipes, err := a.rs.GetAllRecipes()
	if err != nil {}

	ctx.JSON(200, gin.H{
		"message": "success",
		"recipes": recipes,
	})
}

func (a *RecipeHandler) EditRecipe(ctx *gin.Context) {

}

