package property

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	Name        string `json:"name" gorm:"size:255;not null"`
	Location    string `json:"location" gorm:"size:255"`  // e.g., city, coordinates, area
	Description string `json:"description" gorm:"type:text"` //
}
