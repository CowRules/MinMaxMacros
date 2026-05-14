-- +goose Up
ALTER TABLE product
ADD COLUMN user_id UUID NOT NULL,
ADD CONSTRAINT fk_user_id
FOREIGN KEY (user_id)
REFERENCES users(id)
ON DELETE CASCADE;

-- +goose Down
ALTER TABLE product
DROP COLUMN user_id,
DROP CONSTRAINT fk_user_id;