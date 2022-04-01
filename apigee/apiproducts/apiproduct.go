package apiproducts

import "github.com/eduandrade/apigeetool-go/apigee/models"

type ApiProduct struct {
	Name           string             `json:"name"`
	DisplayName    string             `json:"displayName"`
	Description    string             `json:"description"`
	ApprovalType   string             `json:"approvalType,omitempty"`
	ApiResources   []string           `json:"apiResources,omitempty"`
	Scopes         []string           `json:"scopes,omitempty"`
	Proxies        []string           `json:"proxies,omitempty"`
	Environments   []string           `json:"environments,omitempty"`
	CreatedAt      int                `json:"createdAt,omitempty"`
	CreatedBy      string             `json:"createdBy,omitempty"`
	LastModifiedAt int                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy string             `json:"lastModifiedBy,omitempty"`
	Attributes     []models.Attribute `json:"attributes,omitempty"`
}
