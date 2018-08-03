package repository

// CompanyRepository holds the methods for getting company data from DB
type CompanyRepository struct {
}

// GetCompanies fetch companies from DB
func (cr *CompanyRepository) GetCompanies() {
	testUser := struct {
		Name, Lastname string
	}{
		Name:     "Robi",
		Lastname: "Marquez",
	}
	db.C("users").Insert(&testUser)
}
