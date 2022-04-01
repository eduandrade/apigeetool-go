package options

import (
	"errors"
	"fmt"
	"os"
)

type Options struct {
	//TokenApiURL   string
	//BaseURI       string
	//Organization string
	//Username      *string
	//Password      string
	//AccessToken   string
	//Operation     string
	SupportedApis []string

	optionsMap map[string]string
}

//Keys for map options
const (
	TokenApiURL  = "tokenurl"
	BaseURI      = "baseuri"
	Username     = "username"
	Password     = "password"
	Organization = "organization"
	AccessToken  = "accesstoken"
	Operation    = "operation"

	GetOperation    = "GET"
	CreateOperation = "POST"
	UpdateOperation = "PUT"
	DeleteOperation = "DELETE"

	Name     = "name"
	Email    = "email"
	FilePath = "filepath"
)

func New() Options {
	opts := Options{}
	//opts.SupportedApis = make([]string, 3)
	opts.SupportedApis = []string{"developers", "apiproducts", "apps"}

	opts.optionsMap = make(map[string]string)
	opts.optionsMap[TokenApiURL] = "https://login.apigee.com/oauth/token"
	opts.optionsMap[BaseURI] = "https://api.enterprise.apigee.com"

	return opts
}

// func (o *Options) GetUsername() string {
// 	return o.optionsMap[Username]
// }

// func (o *Options) SetUsername(value string) {
// 	o.optionsMap[Username] = value
// }

// func (o *Options) GetPassword() string {
// 	return o.optionsMap[Password]
// }

// func (o *Options) SetPassword(value string) {
// 	o.optionsMap[Password] = value
// }

// func (o *Options) GetTokenApiURL() string {
// 	return o.optionsMap[TokenApiURL]
// }

func (o *Options) Get(key string) string {
	val := os.Getenv(key)
	if val == "" {
		val = o.optionsMap[key]
	}
	//fmt.Printf("%v=%v\n", key, val)
	return val
}

func (o *Options) Set(key, value string) {
	o.optionsMap[key] = value
}

func (o *Options) OrgMgmtURL() string {
	return fmt.Sprintf("%v/v1/o/%v", o.Get(BaseURI), o.Get(Organization))
}

func (o *Options) IsSupportedApi(a string) bool {
	for _, api := range o.SupportedApis {
		if api == a {
			return true
		}
	}
	return false
}

func (o *Options) Operation() string {
	return o.Get(Operation)
}

func (o *Options) SetOperation(create, update, delete, get bool) error {
	if create {
		o.Set(Operation, CreateOperation)
	} else if update {
		o.Set(Operation, UpdateOperation)
	} else if delete {
		o.Set(Operation, DeleteOperation)
	} else if get {
		o.Set(Operation, GetOperation)
	} else {
		return errors.New("invalid operation! valid values are: get, create, update, delete")
	}

	return nil
}
