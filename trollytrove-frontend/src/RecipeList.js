import React, { useState, useEffect } from 'react';

function RecipeList() {
  const [recipes, setRecipes] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const [selectedRecipe, setSelectedRecipe] = useState(null);

  useEffect(() => {
    fetch('http://localhost:8080/recipes')
      .then((response) => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then((data) => {
        setRecipes(data);
        setIsLoading(false);
      })
      .catch((error) => {
        setError(error);
        setIsLoading(false);
      });
  }, []);

  const handleRecipeClick = (id) => {
    // Make an API request to fetch the specific recipe by ID
    fetch(`http://localhost:8080/recipe/${id}`)
      .then((response) => response.json())
      .then((data) => {
        // Set the selectedRecipe state to display the data
        setSelectedRecipe(data);
      })
      .catch((error) => console.error(`Error fetching recipe ${id}:`, error));
  };

  if (isLoading) {
    return <p>Loading...</p>;
  }

  if (error) {
    return <p>Error: {error.message}</p>;
  }

  return (
    <div>
      <h1>Recipe List</h1>
      <ul>
        {recipes.map((recipe) => (
          <li key={recipe.id}>
            <h2>{recipe.title}</h2>
            <p>{recipe.recipe_desc}</p>
            <p>Rating: {recipe.rating}</p>
            <p>Serves: {recipe.r_serves}</p>
            <button onClick={() => handleRecipeClick(recipe.id)}>View Recipe</button>
          </li>
        ))}
      </ul>

      {selectedRecipe && (
        <div className="modal">
          <h2>{selectedRecipe.title}</h2>
          <p>{selectedRecipe.recipe_desc}</p>
          <p>Rating: {selectedRecipe.rating}</p>
          <p>Serves: {selectedRecipe.r_serves}</p>

          <h3>Ingredients</h3>
          <ul>
            {selectedRecipe.ingredients.map((ingredient, index) => (
              <li key={index}>
                {ingredient.name} - {ingredient.uom} - {ingredient.quantity}
              </li>
            ))}
          </ul>

          <h3>Instructions</h3>
          <ol>
            {selectedRecipe.instructions.map((instruction, index) => (
              <li key={index}>{instruction.step_no}: {instruction.instructions}</li>
            ))}
          </ol>
          <button onClick={() => setSelectedRecipe(null)}>Close</button>
        </div>
      )}
    </div>
  );
}

export default RecipeList;
