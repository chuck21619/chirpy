-- name: GetRefreshTokenFromRefreshToken :one
SELECT * FROM refresh_tokens WHERE token=$1;
