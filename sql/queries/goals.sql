-- name: GetAllGoals :many
-- Get all goals registered in 'goals'
SELECT * FROM goals;

-- name: AddGoal :exec
-- Register a goal to 'goals'
INSERT INTO goals
(
    id,
    goal_type,
    goal_plate_count,
    goal_dur,
    goal_decimal,
    goal_number,
    goal_tier
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7
);

-- name: GetGoalsByType :many
-- Get all registered goal of a certain type from 'goals'
SELECT *
FROM goals
WHERE goal_type = $1;

-- name: GetGoalsByTier :many
-- Get all goals of a specified tier from 'goals'
SELECT *
FROM goals
WHERE goal_tier = $1;

-- name: DeleteAllGoals :exec
-- Remove all entries from 'goals'
DELETE FROM goals;

-- name: GetAllGoalIDs :many
-- Get all goal ids from 'goals'
SELECT id FROM goals;

-- name: GetNextGoalTier :one
-- Get the next tier of goals for a specified type and tier
SELECT *
FROM goals
WHERE goal_type = $1
AND goal_tier = $2
LIMIT 1;

-- name: GetGoalTierByID :one
-- Get the tier of the specified goal
SELECT goal_tier 
FROM goals
WHERE id = $1;