CREATE TABLE IF NOT EXISTS notes (
                                     id UUID PRIMARY KEY,
                                     title TEXT NOT NULL,
                                     content TEXT,
                                     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);