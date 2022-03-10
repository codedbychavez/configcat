package services

import (
	"fmt"
)

type Company struct {
	Name string
	Website string
}

func NewCompany(name string, website string) Company {
	return Company{
		Name: name,
		Website: website,
	}
}

type CompanyService interface {
	FindAll() []Company
}


func (company Company) FindAll() string {
	fmt.Println("Calling FindAll() in company_service.go")
	// os.Open("mock/companies.json")
	return "c1, c2, c3"
}