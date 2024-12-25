CREATE TABLE IF NOT EXISTS "video"
(
    "id"        INTEGER NOT NULL,
    "name"      TEXT,
    "tags"      TEXT,
    "rel_path"  TEXT,
    "thumbnail" BLOB,
    PRIMARY KEY ("id")
);