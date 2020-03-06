package entities

import "time"

type Bike struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Make      string
	Model     string
	Reg       string
	Price     float64
	Active    bool
}
