package models

type Realm struct {
	BaseModel
	Name     string `json:"name"`
	Seasonal bool   `json:"seasonal"`
}
