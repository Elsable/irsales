package controllers

import (
	"net/http"

	"github.com/nosnibor89/irsales/repository"

	"github.com/labstack/echo"
)

// GetCompanies return all the companies. Using query string as filters
func GetCompanies(c echo.Context) error {
	companyRepo := repository.CompanyRepository{}

	companyRepo.GetCompanies()

	return c.String(http.StatusOK, "you're in "+c.Path())
}
