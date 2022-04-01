package developers

import (
	"errors"
	"flag"

	"github.com/eduandrade/apigeetool-go/apigee/options"
	"github.com/eduandrade/apigeetool-go/utils"
)

var email *string

type Runnable struct{}

func (r *Runnable) SetFlags(flagSet *flag.FlagSet) {
	email = flagSet.String(options.Email, "", "Developer email (Optional)")
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
		if opts.Get(options.Email) != "" {
			return GetDeveloperByEmail(opts)
		} else {
			return GetAllDevelopers(opts)
		}
	case options.CreateOperation:
		dev := Developer{}
		err := utils.UnmarshalJsonFile(opts.Get(options.FilePath), &dev)
		if err != nil {
			return nil, err
		}
		return CreateDeveloper(opts, dev)
	case options.UpdateOperation:
		dev := Developer{}
		err := utils.UnmarshalJsonFile(opts.Get(options.FilePath), &dev)
		if err != nil {
			return nil, err
		}
		return UpdateDeveloper(opts, dev)
	case options.DeleteOperation:
		return DeleteDeveloper(opts)
	}
	return nil, errors.New("operation not implemented: " + opts.Operation())
}

func setAdditionalOptions(opts options.Options) {
	opts.Set(options.Email, *email)
}
