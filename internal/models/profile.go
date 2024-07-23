package models

import "time"

type Profile struct {
	ID        string    `gorm:"type:uuid;primary_key" json:"profile_id"`
	FirstName string    `gorm:"column:first_name; type:text; not null" json:"first_name"`
	LastName  string    `gorm:"column:last_name; type:text;not null" json:"last_name"`
	Phone     string    `gorm:"type:varchar(255)" json:"phone"`
	AvatarURL string    `gorm:"type:varchar(255)" json:"avatar_url"`
	Userid    string    `gorm:"type:uuid;" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at; not null; autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at; null; autoUpdateTime" json:"updated_at"`
}
