package mypes

import "time"

// Model ...
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id" groups:"compact"`
	CreatedAt time.Time  `json:"created_at" groups:"compact"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
