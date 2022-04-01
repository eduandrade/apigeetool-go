package apiproducts

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/eduandrade/apigeetool-go/apigee/options"
	"github.com/eduandrade/apigeetool-go/utils"
)

type apiproducts struct {
	ApiProduct []ApiProduct
}

func CreateApiProduct(opts options.Options, product ApiProduct) (ApiProduct, error) {
	url := fmt.Sprintf("%v/apiproducts", opts.OrgMgmtURL())
	return createOrUpdateApiProduct(opts, product, url, options.CreateOperation)
}

func UpdateApiProduct(opts options.Options, product ApiProduct) (ApiProduct, error) {
	name := opts.Get(options.Name)
	if name == "" {
		return ApiProduct{}, errors.New("api product name is required")
	}
	url := fmt.Sprintf("%v/apiproducts/%v", opts.OrgMgmtURL(), name)
	return createOrUpdateApiProduct(opts, product, url, options.UpdateOperation)
}

func createOrUpdateApiProduct(opts options.Options, product ApiProduct, url, method string) (ApiProduct, error) {
	payload := utils.PrettyJsonString(product)
	jsonStr, err := utils.CallApi(opts, method, url, strings.NewReader(payload))
	if err != nil {
		return ApiProduct{}, err
	}
	p := ApiProduct{}
	err = utils.UnmarshalJsonString(jsonStr, &p)
	return p, err
}

func DeleteApiProduct(opts options.Options) (string, error) {
	name := opts.Get(options.Name)
	if name == "" {
		return "", errors.New("api product name is required")
	}

	url := fmt.Sprintf("%v/apiproducts/%v", opts.OrgMgmtURL(), name)
	_, err := utils.CallDeleteApi(opts, url)
	if err != nil {
		return "", err
	}
	return "api product deleted: " + name, nil
}

func GetApiProductsByName(opts options.Options) (ApiProduct, error) {
	url := fmt.Sprintf("%v/apiproducts/%v?expand=true", opts.OrgMgmtURL(), opts.Get(options.Name))
	jsonStr, err := utils.CallGetApi(opts, url)
	if err != nil {
		return ApiProduct{}, err
	}
	return toApiProduct(jsonStr)
}

func toApiProduct(jsonStr string) (ApiProduct, error) {
	p := ApiProduct{}
	err := utils.UnmarshalJsonString(jsonStr, &p)
	if err != nil {
		return ApiProduct{}, err
	}
	return p, nil
}

func GetAllApiProducts(opts options.Options) ([]ApiProduct, error) {
	prods := []ApiProduct{}
	count := 100
	lastKey := ""

	for {
		url := fmt.Sprintf("%v/apiproducts?expand=true&count=%v&startKey=%v", opts.OrgMgmtURL(), count, url.QueryEscape(lastKey))
		jsonStr, err := utils.CallGetApi(opts, url)
		if err != nil {
			return []ApiProduct{}, err
		}
		p, err := toApiProducts(jsonStr)
		if err != nil {
			return []ApiProduct{}, err
		}
		if len(p.ApiProduct) == 0 {
			return prods, nil
		}

		if len(lastKey) == 0 {
			prods = append(prods, p.ApiProduct...) //first loop add all elements
		} else {
			prods = append(prods, p.ApiProduct[1:]...) //after first loop skip the first element to not duplicate it
		}

		last := prods[len(prods)-1]
		if last.Name != lastKey {
			lastKey = last.Name
		} else {
			break
		}
	}

	return prods, nil
}

func toApiProducts(jsonStr string) (apiproducts, error) {
	p := apiproducts{}
	err := utils.UnmarshalJsonString(jsonStr, &p)
	if err != nil {
		return apiproducts{}, err
	}
	return p, nil
}
