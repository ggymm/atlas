CREATE TABLE IF NOT EXISTS "index"
(
    "id"    TEXT NOT NULL,
    "term"  TEXT NOT NULL,
    "score" INTEGER,
    PRIMARY KEY ("id", "term")
);
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