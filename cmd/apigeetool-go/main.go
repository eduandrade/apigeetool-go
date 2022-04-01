package main

import (
	"os"

	"github.com/eduandrade/apigeetool-go/apigee/options"
	"github.com/eduandrade/apigeetool-go/cli"
	"github.com/eduandrade/apigeetool-go/utils"
)

func main() {
	opts := options.New()
	if len(os.Args) < 2 || !opts.IsSupportedApi(os.Args[1]) {
		utils.PrintMessageAndExit("An Apigee Management API must be defined as the first argument. Valid values are: %v\nFor more info check: https://apidocs.apigee.com/apis", opts.SupportedApis)
	}

	cli.Run(os.Args[1])
}
