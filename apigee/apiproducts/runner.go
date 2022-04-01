package apiproducts

import (
	"errors"
	"flag"

	"github.com/eduandrade/apigeetool-go/apigee/options"
	"github.com/eduandrade/apigeetool-go/utils"
)

var Name *string

type Runnable struct{}

func (r *Runnable) SetFlags(flagSet *flag.FlagSet) {
	Name = flagSet.String(options.Name, "", "API Product name (Optional)")
}

func (r *Runnable) Run(opts options.Options) interface{} {
	setAdditionalOptions(opts)
	data, err := runCommand(opts)
	if err != nil {
		utils.PrintErrorAndExit(err)
	}
	return data
}

func runCommand(opts options.Options) (interface{}, error) {
	switch opts.Operation() {
	case options.GetOperation:
		if opts.Get(options.Name) != "" {
			return GetApiProductsByName(opts)
		} else {
			return GetAllApiProducts(opts)

		}
	case options.CreateOperation:
		p := ApiProduct{}
		err := utils.UnmarshalJsonFile(opts.Get(options.FilePath), &p)
		if err != nil {
			return nil, err
		}
		return CreateApiProduct(opts, p)
	case options.UpdateOperation:
		p := ApiProduct{}
		err := utils.UnmarshalJsonFile(opts.Get(options.FilePath), &p)
		if err != nil {
			return nil, err
		}
		return UpdateApiProduct(opts, p)
	case options.DeleteOperation:
		return DeleteApiProduct(opts)
	}
	return nil, errors.New("operation not implemented: " + opts.Operation())
}

func setAdditionalOptions(opts options.Options) {
	opts.Set(options.Name, *Name)
}
