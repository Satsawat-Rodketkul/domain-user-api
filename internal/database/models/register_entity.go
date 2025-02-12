package models

import "time"

type RegisterEntity struct {
	UserId     uint      `gorm:"primaryKey"`
	Username   string    `gorm:"size:50;not null"`
	Password   string    `gorm:"size:50;not null"`
	CreateDate time.Time `gorm:"autoCreateTime"`
	UpdateDate time.Time `gorm:"autoCreateTime"`
}
