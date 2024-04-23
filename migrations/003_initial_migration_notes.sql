CREATE TABLE notes (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    author_id INT,
    text TEXT,
    is_public BOOLEAN,
    FOREIGN KEY (author_id) REFERENCES users(id)
);