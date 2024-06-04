BEGIN;

CREATE TABLE events
(
    id          UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    group_id    UUID                     NOT NULL REFERENCES groups (id) ON DELETE CASCADE,
    title       varchar(255)             NOT NULL,
    description varchar(255)             NOT NULL,

    start_at    TIMESTAMP WITH TIME ZONE NOT NULL,
    end_at      TIMESTAMP WITH TIME ZONE NOT NULL,
    location    TEXT,

    records     JSONB,

    assign_type VARCHAR(255)             NOT NULL,

    created_by  UUID                     REFERENCES users (id) ON DELETE SET NULL,

    created_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TRIGGER update_event_updated_at
    BEFORE UPDATE
    ON events
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TABLE event_assignees
(
    id         UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    event_id   UUID NOT NULL REFERENCES events (id) ON DELETE CASCADE,
    user_id    UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE group_records
(
    id                  UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    group_id            UUID         NOT NULL REFERENCES groups (id) ON DELETE CASCADE,
    title               text         NOT NULL,
    type                VARCHAR(255) NOT NULL,
    read_access_scopes  VARCHAR(255)[],
    write_access_scopes VARCHAR(255)[],
    created_at          TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE group_records_default
(
    id                  UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title               text         NOT NULL,
    type                VARCHAR(255) NOT NULL,
    read_access_scopes  VARCHAR(255)[],
    write_access_scopes VARCHAR(255)[]
);

INSERT INTO group_records_default (title, type, read_access_scopes, write_access_scopes)
VALUES ('Присутність', 'string', '{"student", "coach"}', '{"student", "coach"}');
INSERT INTO group_records_default (title, type)
VALUES ('Вага', 'number');
INSERT INTO group_records_default (title, type)
VALUES ('Зріст', 'number');

CREATE TABLE event_record_values
(
    id         UUID NOT NULL            DEFAULT uuid_generate_v4(),
    event_id   UUID NOT NULL REFERENCES events (id) ON DELETE CASCADE,

    user_id    UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    record_id  UUID NOT NULL,

    value      JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE unique index event_record_values_event_id_user_id_record_id_uindex
    ON event_record_values (event_id, user_id, record_id);

CREATE TRIGGER update_event_record_values_updated_at
    BEFORE UPDATE
    ON event_record_values
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

COMMIT;
