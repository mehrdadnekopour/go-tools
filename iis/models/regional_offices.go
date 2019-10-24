package models

import (
	"github.com/mehrdadnekopour/go-tools/mypes"
)

// RegionalOffice ...
type RegionalOffice struct {
	mypes.Model
	Code   int           `json:"code"`
	LogoID uint          `json:"logo_id"`
	Logo   *mypes.Mimage `json:"logo"`
}
