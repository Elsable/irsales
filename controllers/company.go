package controllers

import (
	"net/http"

	"github.com/nosnibor89/irsales/model"

	"github.com/nosnibor89/irsales/services"

	"github.com/labstack/echo"
)

// CompanyController have all the handler for company resources
type CompanyController struct {
	ServiceHandler services.CompanyServiceHandler
}

// NewCompanyController creates a new controller
func NewCompanyController() CompanyController {
	return CompanyController{ServiceHandler: services.NewCompanyService()}
}

// GetCompanies return all the companies. Using query string as filters
func (cc CompanyController) GetCompanies(c echo.Context) error {
	companies := cc.ServiceHandler.Find()
	return c.JSON(http.StatusOK, companies)
}

// SaveCompany is the handler for saving a company
func (cc CompanyController) SaveCompany(c echo.Context) error {
	var created model.Company

	company := model.Company{}

	err := c.Bind(&company)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	created, err = cc.ServiceHandler.Create(&company)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, created)
}
