-- name: CreateRecipe :one
INSERT INTO "recipe" ("name", "description", "userId")
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetRecipe :one
SELECT * FROM "recipe" WHERE "id" = $1;

-- name: ListRecipes :many
SELECT * FROM "recipe"
ORDER BY "createdAt" DESC;

-- name: AddIngredient :one
INSERT INTO "ingredient" ("name", "quantity", "recipeId")
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListIngredientsByRecipe :many
SELECT * FROM "ingredient"
WHERE "recipeId" = $1
ORDER BY "createdAt" ASC;

-- name: AddStep :one
INSERT INTO "step" ("instruction", "stepNumber", "recipeId")
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListStepsByRecipe :many
SELECT * FROM "step"
WHERE "recipeId" = $1
ORDER BY "stepNumber" ASC;

-- name: DeleteRecipeById :exec
DELETE FROM "recipe" WHERE "id"=$1;
