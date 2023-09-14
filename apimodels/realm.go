package apimodels

import "github.com/r35krag0th/datasorcerer/models"

type SpecificRealm struct {
	RealmID string `path:"id" doc:"Realms unique identifier."`
}

type ListRealmsRequest struct{}
type ListRealmsResponse struct {
	CountOfObjects
	Realms []models.Realm `json:"realms"`
}
type ShowRealmRequest struct {
	SpecificRealm
}

type ShowRealmResponse struct {
	Realm models.Realm `json:"realm"`
}

type CreateRealmRequest struct {
	Body struct {
		Name string `json:"name"`
	}
}

type CreateRealmResponse struct {
	Realm models.Realm `json:"realm"`
}

type UpdateRealmRequest struct {
	SpecificRealm
	Body struct {
		Name string `json:"name"`
	}
}

type UpdateRealmResponse struct {
	Realm models.Realm `json:"realm"`
}

type DeleteRealmRequest struct {
	SpecificRealm
}

type DeleteRealmResponse struct {
}
