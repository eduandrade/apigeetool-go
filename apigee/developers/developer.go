package developers

import "github.com/eduandrade/apigeetool-go/apigee/models"

type Developer struct {
	Email          string             `json:"email"`
	FirstName      string             `json:"firstName"`
	LastName       string             `json:"lastName"`
	UserName       string             `json:"userName"`
	DeveloperId    string             `json:"developerId,omitempty"`
	Organization   string             `json:"organizationName,omitempty"`
	Status         string             `json:"status,omitempty"`
	Apps           []string           `json:"apps,omitempty"`
	Companies      []string           `json:"companies,omitempty"`
	CreatedAt      int                `json:"createdAt,omitempty"`
	CreatedBy      string             `json:"createdBy,omitempty"`
	LastModifiedAt int                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy string             `json:"lastModifiedBy,omitempty"`
	Attributes     []models.Attribute `json:"attributes,omitempty"`
}
