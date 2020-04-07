package models

import "time"

type User struct {
	ID        uint32     `gorm:"AUTO_INCREMENT" json:"id"`
	FirstName string     `gorm:"type:varchar(100)" json:"first_name" binding="required"`
	LastName  string     `gorm:"type:varchar(100)" json:"last_name" binding="required"`
	Email     string     `gorm:"type:varchar(100)" json:"email"`
	CreatedAt time.Time  `gorm:"DEFAULT:now()" json:"created_at"`
	UpdatedAt *time.Time `gorm:"DEFAULT:now()" json:"updated_at"`
	DeletedAt *time.Time `gorm:"DEFAULT:NULL" json:"deleted_at"`
}
