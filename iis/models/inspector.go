package models

import (
	"time"

	"github.com/mehrdadnekopour/go-tools/mypes"
)

// Inspector ...
type Inspector struct {
	mypes.Model
	FirstName    string        `json:"first_name"`
	LastName     string        `json:"last_name"`
	Tel          string        `json:"tel"`
	Mobile       string        `json:"mobile"`
	Address      string        `json:"address"`
	Email        string        `json:"email"`
	NationalCode string        `json:"national_code"`
	PersonalCode string        `json:"personal_code"`
	Online       bool          `json:"online"`
	LastLogin    *time.Time    `json:"last_loign"`
	AvatarID     uint          `json:"avatar_id"`
	Avatar       *mypes.Mimage `json:"avatar"`
	// Username     string        `json:"username"`
	// Password     string        `json:"password"`
	// AccessSI     bool          `json:"access_si"`
	// AccessASI    bool          `json:"access_asi"`
}

// Inspectors ...
type Inspectors []*Inspector

// AccessSI     bool   `json:"access_si" gorm:"default:'true'"`
