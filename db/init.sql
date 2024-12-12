CREATE TABLE IF NOT EXISTS clicks (
    id SERIAL PRIMARY KEY,
    banner_id INT NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    count INT NOT NULL
);
