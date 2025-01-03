CREATE TABLE IF NOT EXISTS "base"
(
    "key"   TEXT NOT NULL,
    "value" TEXT,
    PRIMARY KEY ("key")
);
CREATE TABLE IF NOT EXISTS "video"
(
    "id"         TEXT NOT NULL,
    "name"       TEXT,
    "tags"       TEXT,
    "path"       TEXT,
    "size"       INTEGER,
    "cover"      BLOB,
    "format"     TEXT,
    "duration"   INTEGER,
    "created_at" INTEGER,
    "updated_at" INTEGER,
    PRIMARY KEY ("id")
);