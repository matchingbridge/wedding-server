package models

import (
	"time"
)

type Auth struct {
	AuthID       int        `json:"auth_id"       gorm:"not null;primary_key;AUTO_INCREMENT"`
	UserID       string     `json:"user_id"       gorm:"not null;type:varchar(36);unique"`
	Password     string     `json:"password"      gorm:"not null;type:varchar(60)"`
	RefreshToken string     `json:"refresh_token" gorm:"type:varchar(64)"`
	CreatedAt    time.Time  `json:"created_at"    gorm:"not null;default:(CURRENT_TIMESTAMP())"`
	ReviewedAt   *time.Time `json:"reviewed_at"`
}
