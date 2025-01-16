CREATE TABLE IF NOT EXISTS "video"
(
    "id"         TEXT NOT NULL,
    "path"       TEXT,
    "size"       INTEGER,
    "star"       INTEGER,
    "tags"       TEXT,
    "title"      TEXT,
    "cover"      BLOB,
    "format"     TEXT,
    "duration"   INTEGER,
    "created_at" INTEGER,
    "updated_at" INTEGER,
    PRIMARY KEY ("id")
);
CREATE VIRTUAL TABLE IF NOT EXISTS video_index USING fts5
(
    id,
    tags,
    title,
    tokenize = 'simple'
);