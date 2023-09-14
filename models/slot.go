package models

type Slot struct {
	BaseModel
	Name    string            `json:"name"`
	Classes []*CharacterClass `json:"classes" gorm:"many2many:character_class_slots;"`
}

// const (
// 	HelmSlot                  string = "Helm"
// 	ChestSlot                 string = "Chest"
// 	GloveSlot                 string = "Gloves"
// 	LegSlot                   string = "Legs"
// 	BootSlot                  string = "Boots"
// 	TwoHandedBludgeoningSlot  string = "2H (Bludgeoning)"
// 	TwoHandedSlashingSlot     string = "2H (Slashing)"
// 	DualWieldingLeftHandSlot  string = "Dual Wielding (Left)"
// 	DualWieldingRightHandSlot string = "Dial Wielding (Right)"
// 	MainHandSlot              string = "Main-Hand"
// 	OffHandSlot               string = "Off-Hand"
// 	AmuletSlot                string = "Amulet"
// 	RingTopSlot               string = "Ring (Top)"
// 	RingBottomSlot            string = "Ring (Bottom)"
// )
