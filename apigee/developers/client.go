package developers

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/eduandrade/apigeetool-go/apigee/options"
	"github.com/eduandrade/apigeetool-go/utils"
)

type developers struct {
	Developer []Developer
}

func CreateDeveloper(opts options.Options, dev Developer) (Developer, error) {
	url := fmt.Sprintf("%v/developers", opts.OrgMgmtURL())
	return createOrUpdateDeveloper(opts, dev, url, options.CreateOperation)
}

func UpdateDeveloper(opts options.Options, dev Developer) (Developer, error) {
	email := opts.Get(options.Email)
	if email == "" {
		return Developer{}, errors.New("developer email is required")
	}
	url := fmt.Sprintf("%v/developers/%v", opts.OrgMgmtURL(), email)
	return createOrUpdateDeveloper(opts, dev, url, options.UpdateOperation)
}

func createOrUpdateDeveloper(opts options.Options, dev Developer, url, method string) (Developer, error) {
	payload := utils.PrettyJsonString(dev)
	jsonStr, err := utils.CallApi(opts, method, url, strings.NewReader(payload))
	if err != nil {
		return Developer{}, err
	}
	newDev := Developer{}
	err = utils.UnmarshalJsonString(jsonStr, &newDev)
	return newDev, err
}

func DeleteDeveloper(opts options.Options) (string, error) {
	email := opts.Get(options.Email)
	if email == "" {
		return "", errors.New("developer email is required")
	}

	url := fmt.Sprintf("%v/developers/%v", opts.OrgMgmtURL(), email)
	_, err := utils.CallDeleteApi(opts, url)
	if err != nil {
		return "", err
	}
	return "developer deleted: " + email, nil
}

func GetDeveloperByEmail(opts options.Options) (Developer, error) {
	email := opts.Get(options.Email)
	if email == "" {
		return Developer{}, errors.New("developer email is required")
	}

	url := fmt.Sprintf("%v/developers/%v?expand=true", opts.OrgMgmtURL(), email)
	jsonStr, err := utils.CallGetApi(opts, url)
	if err != nil {
		return Developer{}, err
	}

	dev := Developer{}
	err = utils.UnmarshalJsonString(jsonStr, &dev)
	return dev, err
}

// func toDeveloper(jsonStr string) (Developer, error) {
// 	d := Developer{}
// 	err := json.Unmarshal([]byte(jsonStr), &d)
// 	if err != nil {
// 		return Developer{}, err
// 	}
// 	return d, nil
// }

func GetAllDevelopers(opts options.Options) ([]Developer, error) {
	developers := []Developer{}
	count := 100
	lastKey := ""

	for {
		url := fmt.Sprintf("%v/developers?expand=true&count=%v&startKey=%v", opts.OrgMgmtURL(), count, url.QueryEscape(lastKey))
		jsonStr, err := utils.CallGetApi(opts, url)
		if err != nil {
			return []Developer{}, err
		}
		devs, err := toDevelopers(jsonStr)
		if err != nil {
			return []Developer{}, err
		}
		if len(devs.Developer) == 0 {
			return developers, nil
		}

		if len(lastKey) == 0 {
			developers = append(developers, devs.Developer...) //first loop add all elements
		} else {
			developers = append(developers, devs.Developer[1:]...) //after first loop skip the first element to not duplicate it
		}

		last := developers[len(developers)-1]
		if last.Email != lastKey {
			lastKey = last.Email
		} else {
			break
		}

	}

	return developers, nil
}

func toDevelopers(jsonStr string) (developers, error) {
	d := developers{}
	err := utils.UnmarshalJsonString(jsonStr, &d)
	if err != nil {
		return developers{}, err
	}
	return d, nil
}
