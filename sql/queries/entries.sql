-- name: AddEntry :one
-- Add a new entry to 'entries
INSERT INTO entries (
    id,
    user_id,
    date,
    cardio,
    cardio_type,
    plate_count,
    plank_dur,
    weight,
    waist,
    park_run
) VALUES (
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10
)
RETURNING *;

-- name: ResetTable :exec
-- Delete all rows from 'entries'
DELETE FROM entries;

-- name: GetLatestEntryTimestampForUser :one
-- Get the timestamp for the latest entry by the specified user from 'entries
SELECT date FROM entries
WHERE user_id = $1
ORDER BY date DESC
LIMIT 1;

-- name: GetLatestWeeklyDataTimestampForUser :one
-- Get the timestamp for the latest entry that has a non null weekly value (weight, waist)
SELECT date
FROM entries
WHERE user_id = $1
AND weight IS NOT NULL
AND waist IS NOT NULL
LIMIT 1;

-- name: GetLatestPlateCountForUser :one
-- Get the latest plate count for a specified user (Bench Press, L)
SELECT plate_count
FROM entries
WHERE user_id = $1
ORDER BY date DESC
LIMIT 1;

-- name: GetAllEntriesForUser :many
-- Get all the entries for a specified user ordered by ascending date
SELECT *
FROM entries
WHERE user_id = $1;