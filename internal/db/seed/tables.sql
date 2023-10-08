CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS products (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    supermarket VARCHAR(255),
    item_id VARCHAR(255) UNIQUE NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    product_desc VARCHAR(255),
    sales_unit VARCHAR(255) NOT NULL,
    price_per_uom VARCHAR(255),
    price VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS recipes (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    recipe_desc VARCHAR(255),
    rating INT,
    r_serves INT
);

CREATE TABLE IF NOT EXISTS recipe_ingredients (
    recipe_id UUID REFERENCES recipes(id),
    name VARCHAR(255) NOT NULL,
    uom VARCHAR(255) NOT NULL,
    quantity FLOAT NOT NULL
);

CREATE TABLE IF NOT EXISTS recipe_instructions (
    recipe_id UUID REFERENCES recipes(id),
    step_no INT NOT NULL,
    instructions TEXT NOT NULL
);


-- Metadata, tags such as healthy, cheap, etc...
-- worry about this later
-- CREATE TABLE IF NOT EXISTS recipe_tags (

-- )