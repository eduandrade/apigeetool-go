package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/eduandrade/apigeetool-go/apigee/options"
)

type oAuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var accessToken string

func GetAccessToken(opts options.Options) (string, error) {
	if accessToken == "" {
		oauthResponse, err := passwordGrantType(opts)
		if err != nil {
			return "", err
		}
		accessToken = oauthResponse.AccessToken
	}
	return accessToken, nil
}

func passwordGrantType(opts options.Options) (oAuthResponse, error) {

	if len(opts.Get(options.Username)) == 0 || len(opts.Get(options.Password)) == 0 {
		return oAuthResponse{}, errors.New("username and password are required for OAuth request")
	}

	url := opts.Get(options.TokenApiURL)
	payload := strings.NewReader(fmt.Sprintf("username=%v&password=%v&grant_type=password", opts.Get(options.Username), opts.Get(options.Password)))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, payload)

	if err != nil {
		return oAuthResponse{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic ZWRnZWNsaTplZGdlY2xpc2VjcmV0") //https://docs.apigee.com/api-platform/system-administration/management-api-tokens

	res, err := client.Do(req)
	if err != nil {
		return oAuthResponse{}, err
	} else if res.StatusCode != 200 {
		return oAuthResponse{}, errors.New("OAuth API response code is invalid: " + fmt.Sprint(res.StatusCode))
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return oAuthResponse{}, err
	}

	return toOAuthResponse(string(body))
}

func toOAuthResponse(jsonStr string) (oAuthResponse, error) {
	respObj := oAuthResponse{}
	err := json.Unmarshal([]byte(jsonStr), &respObj)
	if err != nil {
		return oAuthResponse{}, err
	}
	return respObj, nil
}
