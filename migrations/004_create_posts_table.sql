-- Migration: Create posts table
-- Date: 2025-11-07
-- Description: Creates posts table with voting system and foreign keys to users and subreddits

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(300) NOT NULL,
    content TEXT,                               -- Text content (optional for link posts)
    post_type VARCHAR(20) DEFAULT 'text',       -- 'text', 'link', 'image'
    link_url TEXT,                              -- URL for link posts
    image_url TEXT,                             -- Image URL (for future image uploads)
    author_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    subreddit_id INTEGER NOT NULL REFERENCES subreddits(id) ON DELETE CASCADE,
    upvotes INTEGER DEFAULT 0,
    downvotes INTEGER DEFAULT 0,
    score INTEGER DEFAULT 0,                    -- Calculated: upvotes - downvotes
    comment_count INTEGER DEFAULT 0,            -- Total comments (updated by triggers/app)
    is_locked BOOLEAN DEFAULT FALSE,            -- Locked posts can't receive new comments
    is_nsfw BOOLEAN DEFAULT FALSE,              -- NSFW content flag
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for performance
CREATE INDEX idx_posts_author ON posts(author_id);
CREATE INDEX idx_posts_subreddit ON posts(subreddit_id);
CREATE INDEX idx_posts_score ON posts(score DESC);
CREATE INDEX idx_posts_created ON posts(created_at DESC);
CREATE INDEX idx_posts_subreddit_score ON posts(subreddit_id, score DESC);
CREATE INDEX idx_posts_subreddit_created ON posts(subreddit_id, created_at DESC);

-- Check constraints
ALTER TABLE posts ADD CONSTRAINT check_post_type 
    CHECK (post_type IN ('text', 'link', 'image'));

ALTER TABLE posts ADD CONSTRAINT check_title_length
    CHECK (length(title) >= 3 AND length(title) <= 300);

-- Comments for documentation
COMMENT ON TABLE posts IS 'User posts within subreddits';
COMMENT ON COLUMN posts.post_type IS 'Type of post: text (self-post), link (external URL), or image (image upload)';
COMMENT ON COLUMN posts.score IS 'Cached score value (upvotes - downvotes) for sorting';
COMMENT ON COLUMN posts.comment_count IS 'Cached count of comments (updated by application logic)';
COMMENT ON COLUMN posts.is_locked IS 'Locked posts prevent new comments (moderation tool)';