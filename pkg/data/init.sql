CREATE TABLE IF NOT EXISTS "video"
(
    "id"         TEXT NOT NULL,
    "name"       TEXT,
    "size"       INTEGER,
    "path"       TEXT,
    "star"       INTEGER,
    "tags"       TEXT,
    "cover"      BLOB,
    "format"     TEXT,
    "duration"   INTEGER,
    "created_at" INTEGER,
    "updated_at" INTEGER,
    PRIMARY KEY ("id")
);