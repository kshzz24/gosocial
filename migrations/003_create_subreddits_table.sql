-- Migration: Create subreddits table
-- Date: 2025-11-06
-- Description: Creates subreddits table with advanced features (JSONB rules, flairs, images)

CREATE TABLE subreddits (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,              -- URL-friendly name (e.g., "programming")
    display_name VARCHAR(100) NOT NULL,            -- Display name (e.g., "Programming")
    description TEXT,                              -- About section
    rules JSONB DEFAULT '[]'::jsonb,               -- Array of rule objects: [{title, description}]
    banner_image_url TEXT,                         -- Top cover image URL
    icon_image_url TEXT,                           -- Small subreddit icon URL
    is_nsfw BOOLEAN DEFAULT FALSE,                 -- Age-restricted content flag
    is_private BOOLEAN DEFAULT FALSE,              -- Private/public subreddit
    created_by INTEGER NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    members_count INTEGER DEFAULT 1,               -- Total members who joined
    active_users INTEGER DEFAULT 0,                -- Currently online users (real-time)
    flairs JSONB DEFAULT '[]'::jsonb,              -- Available post flairs: [{id, name, color}]
    rules_updated_at TIMESTAMP,                    -- Last time rules were modified
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for performance
CREATE INDEX idx_subreddits_name ON subreddits(name);
CREATE INDEX idx_subreddits_created_by ON subreddits(created_by);
CREATE INDEX idx_subreddits_created_at ON subreddits(created_at DESC);
CREATE INDEX idx_subreddits_members_count ON subreddits(members_count DESC);

-- Add check constraints
ALTER TABLE subreddits ADD CONSTRAINT check_name_format 
    CHECK (name ~ '^[a-z0-9_]+$');  -- Only lowercase, numbers, underscores

ALTER TABLE subreddits ADD CONSTRAINT check_name_length
    CHECK (length(name) >= 3 AND length(name) <= 50);

-- Comments for documentation
COMMENT ON TABLE subreddits IS 'Communities where users can create and share posts';
COMMENT ON COLUMN subreddits.name IS 'URL-friendly lowercase name without spaces (3-50 chars)';
COMMENT ON COLUMN subreddits.display_name IS 'Human-readable display name with proper capitalization';
COMMENT ON COLUMN subreddits.rules IS 'JSON array of subreddit rules: [{"title": "...", "description": "..."}]';
COMMENT ON COLUMN subreddits.flairs IS 'JSON array of post flairs: [{"id": "...", "name": "...", "color": "..."}]';
COMMENT ON COLUMN subreddits.members_count IS 'Cached count of members (updated via triggers/app logic)';
COMMENT ON COLUMN subreddits.active_users IS 'Real-time count of active users (updated via WebSocket connections)';