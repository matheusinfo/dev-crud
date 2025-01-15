package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;" json:"id"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
    base.ID = uuid.New()
    return
}

type Dev struct {
	BaseModel
	Username  string `json:"username" gorm:"not null;"`
	Email     string `json:"email" gorm:"unique;not null;"`
}

func (dev *Dev) FindById(db *gorm.DB, id string) (*Dev, error) {
    var result Dev

    err := db.First(&result, id).Error

    if err != nil {
        return nil, err
    }

    return &result, nil
}
