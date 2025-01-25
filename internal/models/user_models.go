package models

import (
	"time"
)

type Role struct {
	Id   uint   `gorm:"primaryKey;not null"`
	Name string `gorm:"not null;unique;size:32"`
}

type User struct {
	Id        uint      `gorm:"primaryKey;not null"`
	Login     string    `gorm:"not null;unique;size:32"`
	Password  string    `gorm:"not null"`
	RoleId    uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime;autoUpdateTime"`
	Role      Role      `gorm:"foreignKey:RoleId"`
}
