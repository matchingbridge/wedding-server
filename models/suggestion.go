package models

import "time"

type Suggestion struct {
	UserID      string    `json:"user_id"      gorm:"not null;type:varchar(36)"                        form:"user_id" binding:"required"`
	PartnerID   string    `json:"partner_id"   gorm:"not null;type:varchar(36)"                        form:"partner_id" binding:"required"`
	SuggestedAt time.Time `json:"suggested_at" gorm:"not null;type:date;default:(CURRENT_TIMESTAMP())" form:"suggested_at" binding:"required"`
}
