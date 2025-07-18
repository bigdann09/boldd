CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    uuid TEXT NOT NULL,
    name VARCHAR(30) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);