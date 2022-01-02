package models

import "time"

type Match struct {
	MatchID      int        `json:"match_id"      gorm:"not null;primary_key;AUTO_INCREMENT"`
	SenderID     string     `json:"sender_id"     gorm:"not null;type:varchar(36)"`
	ReceiverID   string     `json:"receiver_id"   gorm:"not null;type:varchar(36)"`
	AskedAt      *time.Time `json:"asked_at"      gorm:"not null;default:(CURRENT_TIMESTAMP())"`
	MatchedAt    *time.Time `json:"matched_at"`
	TerminatedAt *time.Time `json:"terminated_at"`
}
