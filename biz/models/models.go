package models

import (
	"time"
)

type Users struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	CreatedAt time.Time `gorm:"autoCreateTime"` //创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"` //改密时间
	Level     int       //授权等级
	Password  string
}
