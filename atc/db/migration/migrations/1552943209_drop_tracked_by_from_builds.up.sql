BEGIN;

  DROP MATERIALIZED VIEW transition_builds_per_job;
  DROP MATERIALIZED VIEW next_builds_per_job;
  DROP MATERIALIZED VIEW latest_completed_builds_per_job;

  ALTER TABLE builds DROP COLUMN tracked_by;

COMMIT;
