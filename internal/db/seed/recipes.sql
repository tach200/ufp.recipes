-- Insert the recipe details and retrieve the generated UUID.
WITH inserted_recipe AS (
    INSERT INTO recipes (title, recipe_desc, rating, r_serves)
    VALUES (
        'Squash & Chicken Couscous One-Pot',
        'A delicious one-pot meal with chicken and squash.',
        4,
        4
    )
    RETURNING id
),

inserted_ingredients AS (
    INSERT INTO recipe_ingredients (recipe_id, name, uom, quantity)
    SELECT
        id, name, uom, quantity
    FROM (

        VALUES
            ((SELECT id FROM inserted_recipe), 'Chicken thighs', 'grams', 500),
            ((SELECT id FROM inserted_recipe), 'Olive oil', 'tablespoons', 2),
            ((SELECT id FROM inserted_recipe), 'Onion', 'medium', 1),
            ((SELECT id FROM inserted_recipe), 'Garlic cloves', 'cloves', 2),
            ((SELECT id FROM inserted_recipe), 'Cumin', 'teaspoons', 1),
            ((SELECT id FROM inserted_recipe), 'Coriander', 'teaspoons', 1),
            ((SELECT id FROM inserted_recipe), 'Ground cinnamon', 'teaspoon', 0.5),
            ((SELECT id FROM inserted_recipe), 'Butternut squash', 'grams', 500),
            ((SELECT id FROM inserted_recipe), 'Chicken stock', 'milliliters', 500),
            ((SELECT id FROM inserted_recipe), 'Couscous', 'grams', 250),
            ((SELECT id FROM inserted_recipe), 'Lemon', 'zest and juice', '1'),
            ((SELECT id FROM inserted_recipe), 'Coriander leaves', 'handful', '1')
    ) AS ingredients(id, name, uom, quantity)
)

INSERT INTO recipe_instructions (recipe_id, step_no, instructions)
SELECT
    id, step_no, instructions
FROM (
    VALUES
        ((SELECT id FROM inserted_recipe), 1, 'Heat the oven to 200C/180C fan/gas 6.'),
        ((SELECT id FROM inserted_recipe), 2, 'In a large ovenproof dish, mix the chicken thighs, olive oil, onion, garlic cloves, cumin, coriander, and cinnamon.'),
        ((SELECT id FROM inserted_recipe), 3, 'Season, then roast for 20 mins until the chicken is golden.'),
        ((SELECT id FROM inserted_recipe), 4, 'Stir the squash into the tin, then roast for a further 20 mins until the squash is tender and the chicken cooked.'),
        ((SELECT id FROM inserted_recipe), 5, 'Stir in the couscous, then pour over the chicken stock and lemon zest.'),
        ((SELECT id FROM inserted_recipe), 6, 'Cover with a tightly fitting lid or foil and leave for 10 mins. Squeeze over the lemon juice and fluff up the couscous with a fork, mixing in the coriander leaves before serving.')
) AS instructions(id, step_no, instructions);
