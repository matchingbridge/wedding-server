package models

import "time"

type Suggestion struct {
	UserID      string    `json:"user_id"      gorm:"not null;type:varchar(36);primary_key"                        form:"user_id" binding:"required"`
	PartnerID   string    `json:"partner_ids"  gorm:"not null;type:varchar(36)"                                    form:"partner_ids"`
	SuggestedAt time.Time `json:"suggested_at" gorm:"not null;type:date;default:(CURRENT_TIMESTAMP());primary_key" form:"suggested_at" binding:"required"`
}
