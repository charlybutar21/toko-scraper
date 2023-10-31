CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name TEXT,
    description TEXT,
    image_link TEXT,
    price TEXT,
    rating TEXT,
    merchant TEXT
);
