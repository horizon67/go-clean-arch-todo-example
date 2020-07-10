package domain

import "time"

type Todos []Todo

type Todo struct {
	ID int `gorm:"primary_key;auto_increment"`
	Content string `gorm:"type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
