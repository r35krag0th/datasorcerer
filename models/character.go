package models

import "github.com/google/uuid"

type Character struct {
	BaseModel
	Name             string         `json:"name"`
	Level            int            `json:"level"`
	Realm            Realm          `json:"realm,omitempty"`
	RealmID          uuid.UUID      `json:"realm_id"`
	CharacterClass   CharacterClass `json:"class,omitempty"`
	CharacterClassID uuid.UUID      `json:"class_id"`
	Build            Build          `json:"build,omitempty"`
	BuildID          uuid.UUID      `json:"build_id"`
}
