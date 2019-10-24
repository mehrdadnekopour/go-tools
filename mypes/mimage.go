package mypes

// Mimage ...
type Mimage struct {
	Model
	Title       string `json:"title"`
	Caption     string `json:"caption"`
	Description string `json:"description"`
	Alias       string `json:"alias"`
	URL         string `json:"url"`
}

// Mimages ...
type Mimages []*Mimage

// TableName ...
func (Mimage) TableName() string {
	return "images"
}

// Mimagable ...
type Mimagable struct {
	Model
	ImageID   uint    `json:"image_id"`
	Image     *Mimage `json:"image"`
	OwnerID   uint    `json:"owner_id"`
	OwnerType string  `json:"owner_type"`
}

// Mimagables ....
type Mimagables []*Mimagable

// TableName ...
func (Mimagable) TableName() string {
	return "imagables"
}

// Dimention ...
type Dimention struct {
	Title  string
	Width  uint
	Height uint
}

// Dimentions ...
type Dimentions []*Dimention
