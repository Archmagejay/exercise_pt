-- name: GetGoalsAchievedByUser :many
-- Get all the goals achieved by a specified user id from 'user_goals'
SELECT *
FROM user_goals
WHERE user_id = $1;

-- name: AddGoalAchievedByUser :one
-- Add a goal achieved by a specified user id to 'user_goals'
INSERT INTO user_goals (
    id,
    user_id,
    goal_id
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;