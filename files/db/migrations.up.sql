-- Put your migrations query here

-- Note: use `TIMESTAMP WITH TIMEZONE` to include timezone
CREATE TABLE IF NOT EXISTS
    "go_sample" (
--              id SERIAL,
--              uid VARCHAR(10) PRIMARY KEY NOT NULL,
--              uuid UUID PRIMARY KEY DEFAULT md5(random()::text || clock_timestamp()::text)::uuid,
                 data TEXT NULL,
                 create_time TIMESTAMP DEFAULT now(),
                 create_by VARCHAR(20) NULL,
                 update_time TIMESTAMP NULL,
                 update_by VARCHAR(20) NULL
);

-- CREATE INDEX IF NOT EXISTS go_sample_title_idx ON go_sample (uid, title);
-- CREATE INDEX IF NOT EXISTS go_sample_title_date_idx ON go_sample (uid, title);