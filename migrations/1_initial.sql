-- +migrate Up
CREATE TABLE merchant (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    url VARCHAR
);

-- +migrate Down
DROP TABLE merchant;

-- +migrate Up
CREATE TYPE tea_type AS ENUM ('black', 'green', 'oolong', 'white', 'puerh', 'yellow', 'other');

-- +migrate Down
DROP TYPE tea_type;

-- +migrate Up
CREATE TABLE tea (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    type tea_type NOT NULL,
    merchant_id int NOT NULL REFERENCES merchant
);

-- +migrate Down
DROP TABLE tea;