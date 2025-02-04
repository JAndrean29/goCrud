package model

type User struct {
	ID     int64  `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Age    int    `db:"age" json:"age"`
	Gender string `db:"gender" json:"gender"`
}
