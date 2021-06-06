CREATE TABLE product  (
    id uuid PRIMARY KEY NOT NULL,
    code varchar(128) NOT NULL,
    name varchar(256) NOT NULL,
    value int4 NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz
);