package model

import "time"

type User struct {
	Id        uint      `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name;type:varchar(200)"`
	Email     string    `gorm:"column:email;type:varchar(200);unique;not null"`
	Password  string    `gorm:"_;column:password;type:varchar(200);not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;not null"`
}
