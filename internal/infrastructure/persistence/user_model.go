package persistence

import "time"

type UserModel struct {
	ID        string    `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name;not null"`
	Email     string    `gorm:"column:email;uniqueIndex;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}

func (UserModel) TableName() string { return "users" }
