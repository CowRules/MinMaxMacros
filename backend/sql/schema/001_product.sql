-- +goose Up
CREATE TABLE product (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    price FLOAT NOT NULL,
    weight FLOAT NOT NULL,
    calories FLOAT NOT NULL,
    protein FLOAT NOT NULL,
    fiber FLOAT NOT NULL,
    categories TEXT[],
    shops TEXT[]
);

-- +goose Down
DROP TABLE product;
