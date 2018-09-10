package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/nosnibor89/irsales/controllers"
	"github.com/nosnibor89/irsales/model"

	"github.com/labstack/echo"
	"github.com/nosnibor89/irsales/test"
)

type mockCompanyService struct {
}

// CreateCompany creates a company
func (mc mockCompanyService) Create(company *model.Company) (model.Company, error) {

	company.ID = bson.NewObjectId()
	company.CreatedAt = time.Now().String()

	created := *company

	return created, nil
}

func (mc mockCompanyService) Find() []model.Company {

	return []model.Company{
		model.Company{
			ID:       bson.NewObjectId(),
			Name:     "Tests company 1",
			Location: "Caracas",
		},
		model.Company{
			ID:       bson.NewObjectId(),
			Name:     "Tests company 2",
			Location: "Valencia",
		},
	}
}

func TestSaveCompany(t *testing.T) {

	companyJson := `{"name":"deglotec","location":"Satiago"}`

	app := test.TestApp

	req := httptest.NewRequest(echo.POST, "/api/company", strings.NewReader(companyJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := app.NewContext(req, rec)

	companyController := controllers.CompanyController{ServiceHandler: mockCompanyService{}}
	companyController.SaveCompany(c)
	resJson := rec.Body.String()

	company := model.Company{}
	json.Unmarshal([]byte(resJson), &company)

	if company.ID.String() == "" || company.Name != "deglotec" {
		t.Errorf("Could not save company, got: %v, expected: %v.", resJson, companyJson)
	}

	if http.StatusCreated != rec.Code {
		t.Errorf("Could not save company, wrong http status code receive, got: %v, expected: %v.", rec.Code, http.StatusCreated)
	}
}

func TestFindCompanies(t *testing.T) {
	app := test.TestApp

	req := httptest.NewRequest(echo.GET, "/api/company", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := app.NewContext(req, rec)

	companyController := controllers.CompanyController{ServiceHandler: mockCompanyService{}}
	companyController.GetCompanies(c)
	resJson := rec.Body.String()

	foundCompanies := []model.Company{}
	json.Unmarshal([]byte(resJson), &foundCompanies)

	first := foundCompanies[0]

	if first.Name != "Tests company 1" {
		t.Errorf("Could not fetch companies correctly, got: %v, expected: %v.", first.Name, "Tests company 1")
	}

	if http.StatusOK != rec.Code {
		t.Errorf("Could not fetch companies, wrong http status code receive, got: %v, expected: %v.", rec.Code, http.StatusOK)
	}
}
