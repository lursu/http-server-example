-- migrate:up
-- #TODO Huh?
CREATE EXTENSION IF NOT EXISTS "uuid-o
ssp";

CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT uuid_generate_v4 (),
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    PRIMARY KEY (id)
); 

-- migrate:down

DROP TABLE IF EXISTS users;