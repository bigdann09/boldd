CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    UUID TEXT NOT NULL DEFAULT gen_random_uuid(),
    fullname VARCHAR(50) NOT NULL,
    email VARCHAR(30) NOT NULL,
    phone_number TEXT,
    password TEXT,
    google_id TEXT,
    email_verified bool NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);