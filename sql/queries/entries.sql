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
    waist
) VALUES (
    $1,$2,$3,$4,$5,$6,$7,$8,$9
)
RETURNING *;

-- name: ResetTable :exec
-- Delete all rows from 'entries'
DELETE FROM entries;
