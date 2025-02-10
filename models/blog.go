package models

import (
    "time"
)

type BlogPost struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Title       string    `json:"title" validate:"required"`
    Description string    `json:"description" validate:"required"`
    Body        string    `json:"body" validate:"required"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}