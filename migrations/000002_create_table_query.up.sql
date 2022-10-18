CREATE TABLE IF NOT EXISTS query (
    id                  VARCHAR(36)     NOT NULL DEFAULT uuid_generate_v4(),
    "method"            VARCHAR(50)     NOT NULL,
    address             VARCHAR(50)     NOT NULL,
    code                INTEGER         NOT NULL,
    "time"              NUMERIC(13, 8)  DEFAULT 0 NOT NULL,
    created_at          TIMESTAMP       DEFAULT NOW(),
    updated_at          TIMESTAMP,
    CONSTRAINT query_id_pk PRIMARY KEY (id)
);