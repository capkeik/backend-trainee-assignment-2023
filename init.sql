BEGIN;

CREATE TABLE IF NOT EXISTS users
(
    id INT NOT NULL,
    CONSTRAINT "pk_user_id" PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS segments
(
    id   SERIAL         NOT NULL,
    slug VARCHAR UNIQUE NOT NULL,
    CONSTRAINT "pk_segment_id" PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user_segments
(
    user_id    INTEGER REFERENCES users (id) ON DELETE CASCADE,
    segment_id INTEGER REFERENCES segments (id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, segment_id)
);

CREATE TABLE records
(
    id        SERIAL PRIMARY KEY,
    user_id    INTEGER REFERENCES users (id) ON DELETE CASCADE,
    slug      VARCHAR(255),
    action    VARCHAR(10) CHECK (action IN ('add', 'remove')),
    action_timestamp TIMESTAMPTZ DEFAULT NOW()
);

COMMIT;