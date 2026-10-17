-- Removes download rows that the current counting rules would not have
-- stored: a repeat from the same IP address for the same platform within
-- 10 minutes of the previously kept row. The first row of each burst is
-- kept, matching download.Repository.RecordUnique.
--
-- One-off cleanup of data recorded before the fix. Take a backup first
-- (make db-backup) and run it with:
--   make db-psql < deploy/sql/dedupe-downloads.sql

BEGIN;

CREATE TEMP TABLE duplicate_downloads (id bigint PRIMARY KEY) ON COMMIT DROP;

DO $$
DECLARE
    dedup_window CONSTANT interval := interval '10 minutes';
    row          record;
    group_key    text;
    last_key     text;
    last_kept    timestamptz;
BEGIN
    FOR row IN
        SELECT id, ip_address, os, created_at FROM downloads
        ORDER BY ip_address, os, created_at, id
    LOOP
        group_key := row.ip_address || '|' || row.os;
        IF group_key IS DISTINCT FROM last_key OR row.created_at >= last_kept + dedup_window THEN
            last_key  := group_key;
            last_kept := row.created_at;
        ELSE
            INSERT INTO duplicate_downloads VALUES (row.id);
        END IF;
    END LOOP;
END $$;

SELECT (SELECT count(*) FROM downloads)           AS rows_before,
       (SELECT count(*) FROM duplicate_downloads) AS duplicates_removed;

DELETE FROM downloads WHERE id IN (SELECT id FROM duplicate_downloads);

SELECT os, count(*) AS downloads_after FROM downloads GROUP BY os ORDER BY os;

COMMIT;
