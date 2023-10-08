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

	fmt.Print(recipeID)

	recipes, err := h.DB.SelectRecipe(recipeID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := map[string]string{"error": err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	json.NewEncoder(w).Encode(recipes)
}
