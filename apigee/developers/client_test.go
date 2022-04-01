package developers

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

func TestDeveloperAllOperations(t *testing.T) {
	opts := getOptions()
	opts.Set(options.Email, "test_dev@email.com")
	DeleteDeveloper(opts)

	devToCreate := Developer{}
	err := utils.UnmarshalJsonFile("../../testdata/dev-create.json", &devToCreate)
	if err != nil {
		t.Errorf("Failed to load file: %v", err)
	}
	dev1, err := CreateDeveloper(opts, devToCreate)
	if err != nil {
		t.Errorf("Failed to create developer: %v", err)
	}
	checkProperties(dev1, t)

	devToUpdate := Developer{}
	err = utils.UnmarshalJsonFile("../../testdata/dev-update.json", &devToUpdate)
	if err != nil {
		t.Errorf("Failed to load file: %v", err)
	}
	dev2, err := UpdateDeveloper(opts, devToUpdate)
	if err != nil {
		t.Errorf("Failed to update developer: %v", err)
	}
	checkProperties(dev2, t)

	dev, err := GetDeveloperByEmail(opts)
	if err != nil {
		t.Errorf("Failed to get developer by email: %v", err)
	}
	checkProperties(dev, t)
	if dev.Email != "test_dev@email.com" {
		t.Errorf("invalid email: %v", dev.Email)
	}
	if dev.FirstName != "first_name2" {
		t.Errorf("invalid FirstName: %v", dev.FirstName)
	}
	if len(dev.Attributes) != 4 {
		t.Errorf("invalid Attributes: %v", dev.Attributes)
	}

	_, err = DeleteDeveloper(opts)
	if err != nil {
		t.Errorf("Failed to delete developer: %v", err)
	}
}

// func TestGetDeveloperByEmail(t *testing.T) {
// 	opts := getOptions()
// 	dev, err := GetDeveloperByEmail(opts)
// 	if err != nil {
// 		t.Errorf("Failed to get developer by email: %v", err)
// 	}
// 	checkProperties(dev, t)
// }

func TestGetAllDevelopers(t *testing.T) {
	opts := getOptions()
	devs, err := GetAllDevelopers(opts)
	if err != nil {
		t.Errorf("Failed to get all developers: %v", err)
	}
	if len(devs) == 0 {
		t.Errorf("developers array is empty")
	}
	fmt.Printf("devs=%d\n", len(devs))
	checkProperties(devs[0], t)
}

func checkProperties(dev Developer, t *testing.T) {
	if len(dev.Email) == 0 {
		t.Error("Email was not set")
	}
	if len(dev.FirstName) == 0 {
		t.Error("FirstName was not set")
	}
	if len(dev.LastName) == 0 {
		t.Error("LastName was not set")
	}
	if dev.CreatedAt == 0 {
		t.Error("CreatedAt was not set")
	}
}
