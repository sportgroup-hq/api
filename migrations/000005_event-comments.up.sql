BEGIN;

CREATE TABLE event_comments
(
    id         UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    event_id   UUID NOT NULL REFERENCES events (id) ON DELETE CASCADE,
    user_id    UUID REFERENCES users (id) ON DELETE SET NULL,
    text       TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

COMMIT;
