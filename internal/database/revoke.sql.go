// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: revoke.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const revoke = `-- name: Revoke :exec
UPDATE refresh_tokens SET updated_at=$2, revoked_at=$2 WHERE user_id=$1
`

type RevokeParams struct {
	UserID    uuid.UUID `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) Revoke(ctx context.Context, arg RevokeParams) error {
	_, err := q.db.ExecContext(ctx, revoke, arg.UserID, arg.UpdatedAt)
	return err
}
