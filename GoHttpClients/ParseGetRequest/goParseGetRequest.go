package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("https://partnerportal-dev-ae-web-dev.azurewebsites.net/api/handbooks")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//log.Println(string(body))

	var handbooks []Handbook

	json.Unmarshal([]byte(body), &handbooks)

	for index, element := range handbooks {
		log.Println(index)
		log.Println(element)
	}

	arrayOne := [3]string{"Apple", "Mango", "Banana"}
	for index, rE := range arrayOne {
		log.Println(index)
		log.Println(rE)
	}

	var familyList []FamilyMember
	familyListJon := `[
		{
			"name":"julius",
			"position":"father"
		},
		{
			"name":"faith carmel",
			"position":"mother"
		},
		{
			"name":"Vlad",
			"position":"eldest son"
		}
	]`
	json.Unmarshal([]byte(familyListJon), &familyList)

	for _, element := range familyList {

		fmt.Println(element.Name, element.Position)
	}
}

type FamilyMember struct {
	Name     string
	Position string
}
type Handbook struct {
	name string
}
