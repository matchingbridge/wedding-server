package models

import "time"

type Review struct {
	MatchID   int       `json:"match_id"   gorm:"not null"`
	TargetID  string    `json:"target_id"  gorm:"not null;type:varchar(36)"`
	WrittenAt time.Time `json:"written_at" gorm:"not null;default:(CURRENT_TIMESTAMP())"`
	Content   string    `json:"content"    gorm:"not null;type:text"`
}
