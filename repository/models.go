// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID     int64
	Name   pgtype.Text
	Age    pgtype.Int4
	Gender pgtype.Text
}
