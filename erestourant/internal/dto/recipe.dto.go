package dto

import (
	"github.com/coleYab/erestourant/internal/db/repository"
	"github.com/google/uuid"
)

type Ingredient struct {
	Name     string `json:"name" binding:"required,min=1"`
	Quantity string `json:"quantity" binding:"required"`
}

type Step struct {
	Instruction string `json:"instruction" binding:"required,min=1"`
	StepNumber  int32  `json:"stepNumber" binding:"required,gt=0"`
}

type CreateRecipeDto struct {
	Name        string       `json:"name" binding:"required,min=1"`
	Description string       `json:"description" binding:"required,min=1"`
	UserId      uuid.UUID    `json:"userId" binding:"required"`
	Steps       []Step       `json:"steps" binding:"required,dive,required"`
	Ingredients []Ingredient `json:"ingredients" binding:"required,dive,required"`
}



type Recipe struct {
	ID          uuid.UUID        `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Steps       []Step       `json:"steps" binding:"required,dive,required"`
	Ingredients []Ingredient `json:"ingredients" binding:"required,dive,required"`
}

func GetRecipeDetails(recipe repository.Recipe, steps []repository.Step, ingredients []repository.Ingredient) Recipe {
	rec := Recipe{
		ID: recipe.ID,
		Name: recipe.Name,
		Description: recipe.Description,
	}

	for _, step := range steps {
		rec.Steps = append(rec.Steps, Step{Instruction: step.Instruction, StepNumber: step.StepNumber})
	}

	for _, ing := range ingredients {
		rec.Ingredients = append(rec.Ingredients, Ingredient{Name: ing.Name, Quantity: ing.Quantity})
	}

	return rec
}


