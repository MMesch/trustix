// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package db

import (
	"context"
)

const getDerivation = `-- name: GetDerivation :one
SELECT drv, system FROM derivation
WHERE drv = ? LIMIT 1
`

func (q *Queries) GetDerivation(ctx context.Context, drv string) (Derivation, error) {
	row := q.db.QueryRowContext(ctx, getDerivation, drv)
	var i Derivation
	err := row.Scan(&i.Drv, &i.System)
	return i, err
}

const getEval = `-- name: GetEval :one
SELECT commit_sha, timestamp FROM evaluation
WHERE commit_sha = ? LIMIT 1
`

func (q *Queries) GetEval(ctx context.Context, commitSha string) (Evaluation, error) {
	row := q.db.QueryRowContext(ctx, getEval, commitSha)
	var i Evaluation
	err := row.Scan(&i.CommitSha, &i.Timestamp)
	return i, err
}
