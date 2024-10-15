-- name: Revoke :exec
UPDATE refresh_tokens SET updated_at=$2, revoked_at=$2 WHERE user_id=$1;
