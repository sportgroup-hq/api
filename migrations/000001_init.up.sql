BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    email      VARCHAR(255) NOT NULL UNIQUE CHECK (email ~* '^.+@.+\..+$'),
    first_name VARCHAR(255) NOT NULL,
    last_name  VARCHAR(255),
    picture    VARCHAR(255)
);

COMMIT;
