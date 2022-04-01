package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/eduandrade/apigeetool-go/apigee/options"
)

func PrintErrorAndExit(err error) {
	log.Fatalf("Error: %v", err)
}

func PrintMessageAndExit(str string, v ...interface{}) {
	log.Fatalf(str, v...)
}

func ReadFileContents(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		PrintErrorAndExit(err)
	}
	return string(content)
}

func PrettyJsonString(data interface{}) string {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		PrintErrorAndExit(err)
	}
	return string(val)
}

func UnmarshalJsonFile(filePath string, v interface{}) error {
	contents := ReadFileContents(filePath)
	return UnmarshalJsonString(contents, v)
}

func UnmarshalJsonString(jsonStr string, v interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), v)
	if err != nil {
		return err
	}
	return nil
}

func CallGetApi(opts options.Options, url string) (string, error) {
	return CallApi(opts, options.GetOperation, url, nil)
}

func CallDeleteApi(opts options.Options, url string) (string, error) {
	return CallApi(opts, options.DeleteOperation, url, nil)
}

// func CallPostApi(opts options.Options, url string, payload string) (string, error) {
// 	if payload == "" {
// 		return "", errors.New("payload is empty")
// 	}
// 	return CallApi(opts, options.CreateOperation, url, strings.NewReader(payload))
// }

func CallApi(opts options.Options, method, url string, payload io.Reader) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+opts.Get(options.AccessToken))
	req.Header.Set("Accept", "application/json")
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode > 299 {
		fmt.Println("response=" + string(body))
		return "", errors.New("API response code is invalid: " + fmt.Sprint(res.StatusCode))
	}

	//fmt.Println("response=" + string(body))
	return string(body), nil
}
