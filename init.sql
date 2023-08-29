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
    user_id INTEGER REFERENCES users (id),
    seg_id  INTEGER REFERENCES segments (id)
);

COMMIT;