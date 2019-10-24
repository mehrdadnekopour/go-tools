package mypes

// City ...
type City struct {
	Model
	Name       string    `json:"name"`
	ProvinceID uint      `json:"province_id"`
	Province   *Province `json:"province"`
}

// Cities ...
type Cities []*City
