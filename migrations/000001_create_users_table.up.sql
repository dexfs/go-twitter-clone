CREATE TABLE users
(
    id         VARCHAR(36) PRIMARY KEY, -- Sem default, deve ser fornecido manualmente, uuid V7: ex: 0194b7b8-1110-79c0-8d41-cff79b880911
    username   varchar(50) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
