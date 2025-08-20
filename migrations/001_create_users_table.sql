-- Migration: Create users table
-- Description: Initial migration to create users table with proper schema

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    age INT NOT NULL CHECK (age > 0 AND age <= 150),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Create index for email lookups (already unique due to constraint)
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- Create index for age-based queries
CREATE INDEX IF NOT EXISTS idx_users_age ON users(age);

-- Create index for created_at for time-based queries
CREATE INDEX IF NOT EXISTS idx_users_created_at ON users(created_at);

-- Add comment to table
COMMENT ON TABLE users IS 'User management table for storing user information';
COMMENT ON COLUMN users.id IS 'Unique identifier for the user';
COMMENT ON COLUMN users.name IS 'Full name of the user';
COMMENT ON COLUMN users.email IS 'Unique email address of the user';
COMMENT ON COLUMN users.age IS 'Age of the user (1-150)';
COMMENT ON COLUMN users.created_at IS 'Timestamp when the user was created';
