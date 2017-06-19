package main

import (
	"github.com/salvobabani92/step1/config"
	"github.com/salvobabani92/step1/models"
	"github.com/salvobabani92/step1/controller"
)

func main() {
	config.InitDB()
	models.User{}.CreateTable()
	// Add Default Users
	controller.Add_Default_Records_User()

	models.Location{}.CreateTable()
	// Add Default Locations
	controller.Add_Default_Records_Locations()

	models.Expences{}.CreateTable()
	// Add Default Expences
	controller.Add_Default_Records_Expences()

	models.ExpenceTypes{}.CreateTable()
	// Add Default ExpenceTypes
	controller.Add_Default_Records_ExpenceTypes()

	models.Branches{}.CreateTable()
	// Add Default Branches
	controller.Add_Default_Records_Branches()

}






