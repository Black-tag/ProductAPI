-- +goose Up
ALTER TABLE products 
ADD COLUMN posted_by UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE;

-- +goose Down
ALTER TABLE products
DROP COLUMN posted_by;

