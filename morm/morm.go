package morm

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mehrdadnekopour/go-tools/mypes"
)

// Model ...
type Model struct {
	ID        ID         `gorm:"type:uuid; primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

// BeforeCreate ...
func (m *Model) BeforeCreate(scope *gorm.Scope) error {
	s1, _ := mypes.GetString(*m.ID)
	s2, _ := mypes.GetString(mypes.NilID())

	if s1 != s2 {
		return nil
	}

	scope.SetColumn("ID", mypes.NewUUID())

	return nil
}

// ID ...
type ID = *mypes.GUID //*mypes.GUID //uuid.UUID
