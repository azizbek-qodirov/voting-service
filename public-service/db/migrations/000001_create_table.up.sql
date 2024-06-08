-- First, create the custom ENUM type for gender
DO $$ BEGIN
    CREATE TYPE gender_type AS ENUM ('m', 'f');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

-- Create the 'party' table
CREATE TABLE IF NOT EXISTS party(
    id UUID PRIMARY KEY,
    name VARCHAR(255),
    slogan VARCHAR,
    opened_date TIMESTAMP,
    description VARCHAR,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

-- Create the 'public' table
CREATE TABLE IF NOT EXISTS public(
    id UUID PRIMARY KEY,
    name VARCHAR(50),
    last_name VARCHAR(50),
    phone VARCHAR(13) UNIQUE,
    email VARCHAR(100) UNIQUE,
    birthday TIMESTAMP,
    gender gender_type,
    party_id UUID REFERENCES party(id) NOT NULL,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);
