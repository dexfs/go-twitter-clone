CREATE TABLE users
(
    user_id         VARCHAR(36) PRIMARY KEY, -- Sem default, deve ser fornecido manualmente, ulid : ex: 01JJYY0V9AMD9656HT4BSV0ZEK
    username   varchar(50) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
