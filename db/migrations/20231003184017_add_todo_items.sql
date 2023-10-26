-- migrate:up

CREATE TABLE IF NOT EXISTS todo_items (
    id uuid DEFAULT uuid_generate_v4 (),
    user_id TEXT NOT NULL,
    todo_item TEXT NOT NULL,
    is_completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    PRIMARY KEY (id)
);

-- migrate:down

DROP TABLE IF EXISTS todo_items;