CREATE TABLE IF NOT EXISTS products (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    description text,
    categories text[] NOT NULL,
    quantity integer NOT NULL DEFAULT 0,
    price decimal(10, 2) NOT NULL,
    version integer NOT NULL DEFAULT 1
);
