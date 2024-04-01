CREATE TABLE posts (
    id UUID NOT NULL,
    owner_id UUID NOT NULL,
    content TEXT NOT NULL,
    title TEXT NOT NULL,
    likes INT,
    dislikes INT,
    views INT, 
    category VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
