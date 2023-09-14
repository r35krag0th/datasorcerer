package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

func (base *BaseModel) BeforeCreate(_ *gorm.DB) error {
	var err error
	base.ID, err = uuid.NewUUID()
	return err
}

func (base *BaseModel) Exists(result *gorm.DB) bool {
	return !errors.Is(result.Error, gorm.ErrRecordNotFound) && result.RowsAffected == 1
}
