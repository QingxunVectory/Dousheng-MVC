package model

import "time"

type UserRelation struct {
	Id        int64 `json:"id,omitempty" gorm:"primaryKey;AUTO_INCREMENT"`
	UserID    int64 `json:"user_id"`
	User      User  `json:"user"`
	ToUserID  int64 `json:"to_user_id"`
	ToUser    User  `json:"to_user"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
