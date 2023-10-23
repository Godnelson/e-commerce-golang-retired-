package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string  `gorm:"not null"`
	Price float32 `gorm:"not null"`
	Image string
}
