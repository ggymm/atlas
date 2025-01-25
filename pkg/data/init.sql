CREATE TABLE IF NOT EXISTS "event"
(
    "id"         TEXT NOT NULL,
    "content"    TEXT,
    "service"    TEXT,
    "created_at" INTEGER,
    PRIMARY KEY ("id")
);
CREATE TABLE IF NOT EXISTS "video"
(
    "id"         TEXT NOT NULL,
    "path"       TEXT,
    "size"       INTEGER,
    "tags"       TEXT,
    "title"      TEXT,
    "stars"      INTEGER,
    "cover"      BLOB,
    "format"     TEXT,
    "duration"   INTEGER,
    "created_at" INTEGER,
    "updated_at" INTEGER,
    PRIMARY KEY ("id")
);