package models

import "github.com/google/uuid"

type Build struct {
	BaseModel
	CharacterClass   CharacterClass `json:"class,omitempty"`
	CharacterClassID uuid.UUID      `json:"class_id"`
	Name             string         `json:"name"`
}
