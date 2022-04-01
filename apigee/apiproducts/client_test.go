package apiproducts

import (
	"fmt"
	"testing"

	"github.com/eduandrade/apigeetool-go/apigee/oauth"
	"github.com/eduandrade/apigeetool-go/apigee/options"
	"github.com/eduandrade/apigeetool-go/utils"
)

func getOptions() options.Options {
	opts := options.New()

	token, err := oauth.GetAccessToken(opts)
	if err != nil {
		panic(err)
	}
	opts.Set(options.AccessToken, token)
	return opts
}

func TestApiProductAllOperations(t *testing.T) {
	opts := getOptions()
	opts.Set(options.Name, "testmyapiproduct")

	DeleteApiProduct(opts)

	apiProduct := ApiProduct{}
	err := utils.UnmarshalJsonFile("../../testdata/apiproduct-create.json", &apiProduct)
	if err != nil {
		t.Errorf("Failed to load file: %v", err)
	}

	p1, err := CreateApiProduct(opts, apiProduct)
	if err != nil {
		t.Errorf("Failed to create product: %v", err)
	}
	//fmt.Printf("apiProduct = %+v\n", p1)
	if p1.Name != "testmyapiproduct" {
		t.Errorf("invalid product name: %v", p1.Name)
	}
	if p1.Description != "My API product TEST description" {
		t.Errorf("invalid product Description: %v", p1.Description)
	}
	if p1.DisplayName != "My API product TEST displayName" {
		t.Errorf("invalid product DisplayName: %v", p1.DisplayName)
	}
	if len(p1.Attributes) != 1 {
		t.Errorf("invalid product Attributes: %v", p1.Attributes)
	}
	if len(p1.Scopes) != 1 {
		t.Errorf("invalid product Scopes: %v", p1.Scopes)
	}

	apiProduct = ApiProduct{}
	err = utils.UnmarshalJsonFile("../../testdata/apiproduct-update.json", &apiProduct)
	if err != nil {
		t.Errorf("Failed to load file: %v", err)
	}

	p2, err := UpdateApiProduct(opts, apiProduct)
	if err != nil {
		t.Errorf("Failed to update product: %v", err)
	}
	if p2.Description != "My API product TEST description 2" {
		t.Errorf("invalid product Description: %v", p2.Description)
	}
	if p2.DisplayName != "My API product TEST displayName 2" {
		t.Errorf("invalid product DisplayName: %v", p2.DisplayName)
	}
	if len(p2.Attributes) != 2 {
		t.Errorf("invalid product Attributes: %v", p2.Attributes)
	}
	if len(p2.Scopes) != 2 {
		t.Errorf("invalid product Scopes: %v", p2.Scopes)
	}

	_, err = DeleteApiProduct(opts)
	if err != nil {
		t.Errorf("Failed to delete api product: %v", err)
	}
}

/* func TestGetApiProductsByName(t *testing.T) {
	opts := getOptions()
	p, err := GetApiProductsByName(opts)
	if err != nil {
		t.Errorf("Failed to get API Product by name: %v", err)
	}
	checkProperties(p, t)
} */

func TestGetAllApiProducts(t *testing.T) {
	opts := getOptions()
	products, err := GetAllApiProducts(opts)
	if err != nil {
		t.Errorf("Failed to get all API Products: %v", err)
	}
	if len(products) == 0 {
		t.Errorf("products array is empty")
	}
	fmt.Printf("products=%d\n", len(products))
	checkProperties(products[0], t)
}

func checkProperties(p ApiProduct, t *testing.T) {
	//fmt.Printf("apiProduct = %+v\n", p)
	if len(p.Name) == 0 {
		t.Error("Name was not set")
	}
	if len(p.DisplayName) == 0 {
		t.Error("DisplayName was not set")
	}
	// if len(p.Proxies) == 0 {
	// 	t.Error("Proxies was not set")
	// }
	// if len(p.Environments) == 0 {
	// 	t.Error("Environments was not set")
	// }
}
