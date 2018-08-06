package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetCompanies return all the companies. Using query string as filters
func GetCompanies(c echo.Context) error {
	// companyRepo := repository.CompanyRepository{}

	// companyRepo.GetCompanies()

	return c.JSON(http.StatusOK, Alive(c))
}

func SaveCompany(c echo.Context) error {
	// companyRepo := repository.CompanyRepository{}

	// companyRepo.GetCompanies()

	return c.JSON(http.StatusOK, Alive(c))
}
