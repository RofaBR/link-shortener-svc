CREATE TABLE links (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    original TEXT NOT NULL UNIQUE,
    shortened TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT now()
);