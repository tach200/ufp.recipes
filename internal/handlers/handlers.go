package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tach200/ufp.recipes/internal/db"
)

type Handlers struct {
	DB *db.DB
}

func (h *Handlers) GetRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	recipes, err := h.DB.SelectRecipes()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := map[string]string{"error": err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	json.NewEncoder(w).Encode(recipes)
}

func (h *Handlers) GetRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	recipeID := params["id"]

	recipes, err := h.DB.SelectRecipe(recipeID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := map[string]string{"error": err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	json.NewEncoder(w).Encode(recipes)
}

func (h *Handlers) GetRecipeProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	recipeID := params["id"]

	recipes, err := h.DB.SelectRecipe(recipeID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := map[string]string{"error": err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	var allProducts []*db.Product
	for _, ingrediant := range recipes.RecipeIngredients {
		products, err := h.DB.SelectProduct(ingrediant.Name)
		if err != nil {
			fmt.Printf("warn: couldn't select product with name %s, err %s\n", ingrediant.Name, err.Error())
		}
		allProducts = append(allProducts, products...)
	}

	json.NewEncoder(w).Encode(allProducts)
}

func (h *Handlers) Test(w http.ResponseWriter, r *http.Request) {}
