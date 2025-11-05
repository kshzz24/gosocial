-- Migration: Add password reset functionality to users table
-- Date: 2025-11-06
-- Description: Adds reset_token and reset_token_expires columns for forgot password feature

ALTER TABLE users 
ADD COLUMN reset_token VARCHAR(255),
ADD COLUMN reset_token_expires TIMESTAMP;

-- Add index for faster token lookups
CREATE INDEX idx_users_reset_token ON users(reset_token);

-- Comments for documentation
COMMENT ON COLUMN users.reset_token IS 'Secure token for password reset, expires after 1 hour';
COMMENT ON COLUMN users.reset_token_expires IS 'Timestamp when reset token expires';