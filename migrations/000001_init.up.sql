BEGIN;

CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(50) PRIMARY KEY,
    customer_id VARCHAR(50) NOT NULL,
    currency VARCHAR(10) NOT NULL,
    amount BIGINT NOT NULL CHECK (amount > 0),
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_orders_customer_id ON orders(customer_id);

COMMIT;