package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type CompanyService struct {}

type Company struct {
	Name string `json:"name"`
	Website string `json:"website"`
}

type Companies struct {
	Companies []Company `json:"companies"`
}

func NewCompany(name string, website string) Company {
	return Company{
		Name: name,
		Website: website,
	}
}

type CompanyMethods interface {
	FindAll() []Company
}


func (companyService CompanyService) FindAll() Companies {
	jsonFile, err := os.Open("mock/companies.json")
	if err != nil {
		fmt.Println(err)

	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var companies Companies

	json.Unmarshal(byteValue, &companies)

	return companies
}