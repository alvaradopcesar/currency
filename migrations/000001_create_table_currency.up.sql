CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS currency(
    id                  VARCHAR(36)     NOT NULL DEFAULT uuid_generate_v4(),
    customer_id         SERIAL          NOT NULL,
    code                VARCHAR(5)      NOT NULL,
    "value"             NUMERIC(13, 8)  DEFAULT 0 NOT NULL,
    created_at          TIMESTAMP       DEFAULT NOW(),
    updated_at          TIMESTAMP,
    CONSTRAINT currency_id_pk PRIMARY KEY (id)
);