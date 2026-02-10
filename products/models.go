package products

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        string `gorm:"primaryKey;type:text"`
	Name      string `gorm:"not null"`
	Price     int64  `gorm:"not null"`
	Stock     int    `gorm:"not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()
	return nil
}
