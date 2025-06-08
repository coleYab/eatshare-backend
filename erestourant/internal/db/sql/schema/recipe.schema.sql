CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "recipe" (
    "id" UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "userId" UUID NOT NULL REFERENCES "user"("id"),
    "createdAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Ingredient table
CREATE TABLE "ingredient" (
    "id" UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" TEXT NOT NULL,
    "quantity" TEXT NOT NULL,
    "recipeId" UUID NOT NULL REFERENCES "recipe"("id") ON DELETE CASCADE,
    "createdAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Step table
CREATE TABLE "step" (
    "id" UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    "instruction" TEXT NOT NULL,
    "stepNumber" INT NOT NULL,
    "recipeId" UUID NOT NULL REFERENCES "recipe"("id") ON DELETE CASCADE,
    "createdAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

