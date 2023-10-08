package db

import (
	sqlmapper "github.com/jackskj/carta"
)

// Recipe define a struct for the 'recipes' table.
type Recipe struct {
	ID                 string              `json:"id" db:"recipe_id"`
	Title              string              `json:"title" db:"recipe_title"`
	RecipeDesc         string              `json:"recipe_desc" db:"recipe_desc"`
	Rating             int                 `json:"rating" db:"recipe_rating"`
	ReccServes         int                 `json:"r_serves" db:"recipe_serves"`
	RecipeIngredients  []RecipeIngredient  `json:"ingredients"`
	RecipeInstructions []RecipeInstruction `json:"instructions"`
}

// RecipeIngredient defines a struct for the 'recipe_ingredients' table.
type RecipeIngredient struct {
	Name     string  `json:"name" db:"ingredient_name"`
	UOM      string  `json:"uom" db:"ingredient_unit_of_measure"`
	Quantity float64 `json:"quantity" db:"ingredient_quantity"`
}

// RecipeInstruction defines a struct for the 'recipe_instructions' table.
type RecipeInstruction struct {
	StepNo       int    `json:"step_no" db:"instruction_step_number"`
	Instructions string `json:"instructions" db:"instruction_desc"`
}

const selectRecipeSQL = `
	SELECT 
		id AS recipe_id, 
		title AS recipe_title, 
		recipe_desc AS recipe_desc, 
		rating AS recipe_rating,
		r_serves AS recipe_serves
	FROM recipes;
`

const selectRecipesSQL = `
    SELECT
		r.id AS recipe_id,
        r.title AS recipe_title,
        r.recipe_desc AS recipe_desc,
        r.rating AS recipe_rating,
        r.r_serves AS recipe_serves,
        ri.name AS ingredient_name,
        ri.uom AS ingredient_unit_of_measure,
        ri.quantity AS ingredient_quantity,
        i.step_no AS instruction_step_number,
        i.instructions AS instruction_desc
    FROM recipes AS r
    LEFT JOIN recipe_ingredients AS ri ON r.id = ri.recipe_id
    LEFT JOIN recipe_instructions AS i ON r.id = i.recipe_id
    WHERE r.id = $1
    ORDER BY ri.recipe_id, i.step_no;
`

// SelectRecipes returns some basic information about all recipes
func (db *DB) SelectRecipes() ([]*Recipe, error) {
	rows, err := db.Client.Query(selectRecipeSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recipes := []*Recipe{}
	err = sqlmapper.Map(rows, &recipes)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}

// SelectRecipe returns all information about a single recipe
func (db *DB) SelectRecipe(id string) (*Recipe, error) {

	rows, err := db.Client.Query(selectRecipesSQL, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recipes := &Recipe{}
	err = sqlmapper.Map(rows, recipes)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}
