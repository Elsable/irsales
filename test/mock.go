package test

import (
	"github.com/nosnibor89/irsales/model"
	"gopkg.in/mgo.v2/bson"
)

func MockCompanyData() []model.Company {
	return []model.Company{
		model.Company{
			ID:       bson.ObjectIdHex("507f191e810c19729de860ea"),
			Name:     "Tests company 1",
			Location: "Caracas",
		},
		model.Company{
			ID:       bson.ObjectIdHex("507f1f77bcf86cd799439011"),
			Name:     "Tests company 2",
			Location: "Valencia",
		},
	}
}
