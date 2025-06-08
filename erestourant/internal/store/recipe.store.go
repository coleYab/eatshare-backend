package store

import (
	"context"

	"github.com/coleYab/erestourant/internal/db/repository"
	"github.com/coleYab/erestourant/internal/dto"
	"github.com/google/uuid"
)

type RecipeStore struct {
	qry *repository.Queries
}

func NewRecipeStore(qry *repository.Queries) *RecipeStore {
	return &RecipeStore{qry: qry}
}

func (s *RecipeStore) CreateRecipe(recipePayload dto.CreateRecipeDto) (repository.Recipe, error) {
	ctx := context.Background()
	recipe, err := s.qry.CreateRecipe(ctx, repository.CreateRecipeParams{
		Name: recipePayload.Name,
		Description: recipePayload.Description,
			UserId: recipePayload.UserId,
	})

	if err != nil {
		return repository.Recipe{}, err
	}

	for _, step := range recipePayload.Steps {
		s.qry.AddStep(ctx, repository.AddStepParams{
			StepNumber: step.StepNumber,
			Instruction: step.Instruction,
			RecipeId: recipe.ID,
		})
	}

	for _, ingredient := range recipePayload.Ingredients {
		s.qry.AddIngredient(ctx, repository.AddIngredientParams{
			Name: ingredient.Name,
			Quantity: ingredient.Quantity,
			RecipeId: recipe.ID,
		})

	}

	return recipe, nil
}

func (s *RecipeStore) GetRecipeById(id string) (repository.Recipe, error) {
	uid, _ := uuid.Parse(id)
	ctx := context.Background()
	return s.qry.GetRecipe(ctx, uid)
}


func (s *RecipeStore) GetIngredientByRecipeId(id string) ([]repository.Ingredient, error) {
	uid, _ := uuid.Parse(id)
	ctx := context.Background()
	return s.qry.ListIngredientsByRecipe(ctx, uid)
}

func (s *RecipeStore) GetStepsByRecipeId(id string) ([]repository.Step, error) {
	uid, _ := uuid.Parse(id)
	ctx := context.Background()
	return s.qry.ListStepsByRecipe(ctx, uid)
}

func (s *RecipeStore) GetAllRecipes() ([]repository.Recipe, error) {
	ctx := context.Background()
	return s.qry.ListRecipes(ctx)
}

func (s *RecipeStore) DeleteRecipe(id string) error {
	uid, _ := uuid.Parse(id)
	ctx := context.Background()
	return s.qry.DeleteRecipeById(ctx, uid)
}
