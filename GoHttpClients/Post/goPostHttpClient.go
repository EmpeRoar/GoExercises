package main

/*

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)
*/

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	/*
		requestBody, err := json.Marshal(map[string]string{
			"name": "julius: Go Post Request",
			"description": "this is a Post Request from a Go Application",
			"categoryId": 1,
			"statusId": 1,
			"countryId": [13]
		})

		if err != nil {
			log.Fatalln(err)
		}

		resp. err := http.Post("https://partnerportal-dev-ae-web-dev.azurewebsites.net/api/handbooks",
							   "application/json",
							   bytes.NewBuffer(requestBody))
	*/

	jsonStr := []byte(`{
		"name": "julius: Go Post Request",
		"description": "this is a Post Request from a Go Application",
		"categoryId": 1,
		"statusId": 1,
		"countryId": [13] 
	}`)

	resp, err := http.Post(
		"https://partnerportal-dev-ae-web-dev.azurewebsites.net/api/handbooks",
		"application/json",
		bytes.NewBuffer(jsonStr))

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}
