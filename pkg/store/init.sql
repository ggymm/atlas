CREATE TABLE IF NOT EXISTS "base"
(
    "key"   TEXT NOT NULL,
    "value" TEXT,
    PRIMARY KEY ("key")
);
CREATE TABLE IF NOT EXISTS "video"
(
    "id"        INTEGER NOT NULL,
    "name"      TEXT,
    "tags"      TEXT,
    "size"      INTEGER,
    "rel_path"  TEXT,
    "thumbnail" BLOB,
    PRIMARY KEY ("id")
);