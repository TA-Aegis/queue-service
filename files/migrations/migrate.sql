CREATE TABLE IF NOT EXISTS tenants (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(75) unique not null,
    password VARCHAR(155) not null,
    name VARCHAR(155) not null,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp
);