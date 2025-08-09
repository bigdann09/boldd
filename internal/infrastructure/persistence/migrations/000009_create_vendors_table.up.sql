CREATE TABLE IF NOT EXISTS vendors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID UNIQUE FOREIGN KEY REFERENCES users(id) ON DELETE RESTRICT,
    name VARCHAR(255) NOT NULL,
    business_email VARCHAR(255) NOT NULL,
    business_address TEXT NOT NULL,
    business_phone VARCHAR(50) NOT NULL,
    banner_url TEXT,
    logo_url TEXT,
    description TEXT,
    status VARCHAR(50) NOT NULL CHECK (status IN ('pending', 'processing', 'verified')) DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);