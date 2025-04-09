package property

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	Name        string `gorm:"size:255;not null"`
	Location    string `gorm:"size:255"`  // e.g., city, coordinates, area
	Description string `gorm:"type:text"` //
}
