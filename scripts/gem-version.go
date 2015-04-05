package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Version struct {
	Number string `json:"number"`
	Date   string `json:"built_at"`
}

func getUrlContents(url string) ([]byte, error) {
	var err error
	var resp *http.Response
	var body []byte

	resp, err = http.Get(url)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return body, errors.New("Server responded with non-200 status")
	}

	body, err = ioutil.ReadAll(resp.Body)
	return body, err
}

func parseJson(body []byte) {
	var versions []Version

	err := json.Unmarshal(body, &versions)
	if err != nil {
		fmt.Println("Failed to parse JSON")
		return
	}

	if len(versions) == 0 {
		fmt.Println("Gem not found")
		return
	}

	fmt.Printf(
		"Latest version: %s\nReleased: %s\n",
		versions[0].Number,
		versions[0].Date,
	)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Name required")
		os.Exit(1)
	}

	name := os.Args[1]
	url := fmt.Sprintf("https://rubygems.org/api/v1/versions/%s.json", name)
	result, err := getUrlContents(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	parseJson(result)
}
