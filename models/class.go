package models

type CharacterClass struct {
	BaseModel
	Name  string  `json:"name"`
	Slots []*Slot `json:"slots" gorm:"many2many:character_class_slots;"`
}
