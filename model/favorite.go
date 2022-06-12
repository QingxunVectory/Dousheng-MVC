package model

import (
	"time"
)

type Favorite struct {
	Id        int64 `json:"id,omitempty" gorm:"primaryKey;AUTO_INCREMENT"`
	UserID    int64 `json:"user_id"`
	User      User  `json:"user"`
	VideoID   int64 `json:"video_id"`
	Video     Video `json:"video"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
