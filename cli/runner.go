package cli

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/eduandrade/apigeetool-go/apigee/apiproducts"
	"github.com/eduandrade/apigeetool-go/apigee/developers"
	"github.com/eduandrade/apigeetool-go/apigee/oauth"
	"github.com/eduandrade/apigeetool-go/apigee/options"
	"github.com/eduandrade/apigeetool-go/utils"
)

type Runnable interface {
	SetFlags(flagSet *flag.FlagSet)
	Run(opts options.Options) interface{}
}

//type stubMapping map[string]interface{}
//var funcs = stubMapping{}

var runners = map[string]Runnable{}

func init() {
	// funcs = map[string]interface{}{
	// 	"developers-SetFlags": developers.SetFlags,
	// 	"developers-Run":      developers.Run,
	// }

	runners = make(map[string]Runnable)
	runners["developers"] = &developers.Runnable{}
	runners["apiproducts"] = &apiproducts.Runnable{}
}

func Run(api string) {

	opts := options.New()

	flagSet := flag.NewFlagSet(api, flag.ExitOnError)
	//global flags
	baseUri := flagSet.String("baseuri", opts.Get(options.BaseURI), "The base URI for your organization on Apigee Edge")
	tokenUrl := flagSet.String("tokenurl", opts.Get(options.TokenApiURL), "Apigee Token API URL")
	username := flagSet.String("u", "", "Your Apigee account username (Required if a token is not provided)")
	password := flagSet.String("p", "", "Your Apigee account password (Required if a token is not provided)")
	token := flagSet.String("t", "", "Your Apigee access token. Use this in lieu of -u / -p")
	organization := flagSet.String("o", "", "Organization name (Required)")
	filePath := flagSet.String("f", "", "Path of the file with contents for the API payload (Required if operation is create or update)")
	//output := flagSet.String("out", "", "Output format (console, file)")

	//CRUD operations
	get := flagSet.Bool("get", true, "Get operation - Retrieve the record(s) from the selected API")
	create := flagSet.Bool("create", false, "Create operation - Insert a new record in the selected API")
	update := flagSet.Bool("update", false, "Update operation - Update an existing record in the selected API")
	delete := flagSet.Bool("delete", false, "Delete operation - Delete a record from the selected API")

	//set API specific flags
	//call(api+"-SetFlags", flagSet)
	r := runners[api]
	r.SetFlags(flagSet)
	flagSet.Parse(os.Args[2:])

	opts.Set(options.BaseURI, *baseUri)
	opts.Set(options.TokenApiURL, *tokenUrl)
	opts.Set(options.Username, *username)
	opts.Set(options.Password, *password)
	opts.Set(options.AccessToken, *token)
	opts.Set(options.Organization, *organization)
	opts.Set(options.FilePath, *filePath)

	checkRequiredOptions(opts, flagSet)
	setOperation(opts, create, update, delete, get)
	getAccessTokenIfNotSet(opts)

	//call(api+"-Run", opts)
	data := r.Run(opts)
	val := utils.PrettyJsonString(data)
	println(string(val))
}

func checkRequiredOptions(opts options.Options, flagSet *flag.FlagSet) {
	errOptions := checkOptions(opts)
	if errOptions != nil {
		//utils.PrintErrorAndExit(errOptions)
		fmt.Print("Error: " + errOptions.Error() + "\n")
		flagSet.PrintDefaults()
		os.Exit(1)
	}
}

func setOperation(opts options.Options, create *bool, update *bool, delete *bool, get *bool) {
	err := opts.SetOperation(*create, *update, *delete, *get)
	if err != nil {
		utils.PrintErrorAndExit(err)
	}
}

func getAccessTokenIfNotSet(opts options.Options) {
	if opts.Get(options.AccessToken) == "" {
		tok, err := oauth.GetAccessToken(opts)
		if err != nil {
			utils.PrintErrorAndExit(err)
		}
		opts.Set(options.AccessToken, tok)
	}
}

func checkOptions(opts options.Options) error {
	if opts.Get(options.AccessToken) == "" && opts.Get(options.Username) == "" && opts.Get(options.Password) == "" {
		return errors.New("missing required parameter: you need to provide a username/password or a valid access token")
	} else if opts.Get(options.AccessToken) == "" && opts.Get(options.Username) == "" {
		return errors.New("missing required parameter: username")
	} else if opts.Get(options.AccessToken) == "" && opts.Get(options.Password) == "" {
		return errors.New("missing required parameter: password")
	} else if len(opts.Get(options.Organization)) == 0 {
		return errors.New("missing required parameter: organization")
	}
	return nil
}

// func call(funcName string, params ...interface{}) (result interface{}, err error) {
// 	f := reflect.ValueOf(funcs[funcName])
// 	if len(params) != f.Type().NumIn() {
// 		err = errors.New("the number of params is out of index")
// 		return
// 	}
// 	in := make([]reflect.Value, len(params))
// 	for k, param := range params {
// 		in[k] = reflect.ValueOf(param)
// 	}
// 	//var res []reflect.Value
// 	f.Call(in)
// 	//result = res[0].Interface()
// 	return
// }
