package model

import "time"

type Branch struct {
	Name      string     `dbq:"name"`
	ID        string     `dbq:"id"`
	Address   string     `dbq:"address"`
	CreatedAt time.Time  `dbq:"created_at"`
	UpdatedAt *time.Time `dbq:"updated_at"`
	DeletedAt *time.Time `dbq:"deleted_at"`
}

func (Branch) GetTableName() string {
	return "branchs"
}
