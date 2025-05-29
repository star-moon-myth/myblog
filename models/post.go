package models

import (
    "time"
    _"gorm.io/gorm"
)

type Post struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Title     string    `gorm:"type:varchar(255);not null" json:"title"`
    Content   string    `gorm:"type:text;not null" json:"content"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}