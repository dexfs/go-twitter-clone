CREATE TABLE posts
(
    post_id                   VARCHAR(36) PRIMARY KEY, -- Sem default, deve ser fornecido manualmente, ulid: ex: 01JJYY0V9AMD9656HT4BSV0ZEK
    user_id                   VARCHAR(36) NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    content                   TEXT        NOT NULL,
    is_quote                  BOOLEAN     NOT NULL DEFAULT FALSE,
    is_repost                 BOOLEAN     NOT NULL DEFAULT FALSE,
    original_post_id          VARCHAR(36) NULL,
    original_post_content     TEXT        NULL,
    original_post_user_id     VARCHAR(36) NULL,
    original_post_screen_name TEXT        NULL,
    created_at                TIMESTAMPTZ NOT NULL DEFAULT NOW()
--     CONSTRAINT fk_original_post FOREIGN KEY (original_post_id) REFERENCES posts (id) ON DELETE SET NULL
);

CREATE INDEX idx_posts_user_created_at ON posts (user_id, created_at DESC);
CREATE INDEX idx_posts_original_post_id ON posts (original_post_id);