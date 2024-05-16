BEGIN;

ALTER TABLE users
    RENAME COLUMN picture TO picture_url;

alter table users
    add created_at timestamp DEFAULT current_timestamp NOT NULL;
alter table users
    add updated_at timestamp DEFAULT current_timestamp NOT NULL;
alter table users
    add phone varchar(255);
alter table users
    add date_of_birth date;

CREATE OR REPLACE FUNCTION update_updated_at()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_user_updated_at
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at();

CREATE TABLE groups
(
    id         uuid PRIMARY KEY default uuid_generate_v4(),
    name       VARCHAR(255)                               NOT NULL,
    sport      VARCHAR(255)                               NOT NULL,
    owner_id   uuid                                       REFERENCES users (id) ON DELETE SET NULL,
    created_at TIMESTAMP        DEFAULT current_timestamp NOT NULL,
    updated_at TIMESTAMP        DEFAULT current_timestamp NOT NULL
);

CREATE TABLE group_invites
(
    id         uuid PRIMARY KEY default uuid_generate_v4(),
    group_id   uuid REFERENCES groups (id) ON DELETE CASCADE,
    code       VARCHAR(255)                               NOT NULL,
    active     BOOLEAN          DEFAULT TRUE,
    created_at TIMESTAMP        DEFAULT current_timestamp NOT NULL,
    updated_at TIMESTAMP        DEFAULT current_timestamp NOT NULL
);

CREATE TYPE group_member_type AS ENUM ('student', 'admin', 'owner');

CREATE TABLE group_members
(
    id         uuid UNIQUE default uuid_generate_v4(),
    group_id   uuid REFERENCES groups (id) ON DELETE CASCADE,
    user_id    uuid REFERENCES users (id) ON DELETE CASCADE,
    type       group_member_type                     NOT NULL,
    created_at TIMESTAMP   DEFAULT current_timestamp NOT NULL,
    updated_at TIMESTAMP   DEFAULT current_timestamp NOT NULL,
    PRIMARY KEY (group_id, user_id)
);

CREATE TRIGGER update_group_updated_at
    BEFORE UPDATE
    ON groups
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at();

COMMIT;
