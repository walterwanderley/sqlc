// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const placeholder = `-- name: Placeholder :many
SELECT boo from foo
`

func (q *Queries) Placeholder(ctx context.Context) ([]sql.NullString, error) {
	rows, err := q.db.QueryContext(ctx, placeholder)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var boo sql.NullString
		if err := rows.Scan(&boo); err != nil {
			return nil, err
		}
		items = append(items, boo)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
