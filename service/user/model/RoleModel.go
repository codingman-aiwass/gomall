package model

import "time"

type RoleModel struct {
	ID        uint32    `gorm:"primaryKey;autoIncrement;column:id"`
	Name      string    `gorm:"size:64;unique;not null;column:name"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;column:created_at"`
}
