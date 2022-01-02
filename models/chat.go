package models

import "time"

type Chat struct {
	UserID    string    `json:"user_id"    gorm:"not null;type:varchar(36)"`
	IsTalking bool      `json:"is_talking" gorm:"not null;type:bool"`
	Message   string    `json:"message"    gorm:"not null;type:text"`
	ChattedAt time.Time `json:"chatted_at" gorm:"not null;default:(CURRENT_TIMESTAMP())"`
}
