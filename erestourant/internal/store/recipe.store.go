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

func (s *RecipeStore) CreateRecipe(userParams dto.RegisterDto) (repository.Recipe, error) {
	ctx := context.Background()
	return s.qry.CreateRecipe(ctx, repository.CreateRecipeParams{ })
}


func (s *RecipeStore) GetRecipeById(id string) (repository.Recipe, error) {
	uid, _ := uuid.Parse(id)
	ctx := context.Background()
	return s.qry.GetRecipe(ctx, uid)
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
