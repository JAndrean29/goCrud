package model

type User struct {
	ID     int64  `db:"id" json:"id"`
	Name   string `db:"name" json:"Name"`
	Age    int    `db:"age" json:"Age"`
	Gender string `db:"age" json:"Gender"`
}
