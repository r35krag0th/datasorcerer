package models

import "github.com/google/uuid"

type Weight struct {
	BaseModel
	Build   Build     `json:"build"`
	BuildID uuid.UUID `json:"build_id"`
	Slot    Slot      `json:"slot,omitempty"`
	SlotID  uuid.UUID `json:"slot_id"`
	Rank    int       `json:"rank"`
	Value   float64   `json:"value"`
}
