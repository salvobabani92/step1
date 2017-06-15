package main

import (
	"github.com/salvobabani92/step1/config"
	"github.com/salvobabani92/step1/models"
)

func main() {
	config.InitDB()
	models.User{}.CreateTable()

	var User = models.User{}
	User.Name = "Salvator"
	User.Password = "123456"
	User.Surname = "Babani"
	User.Username = "Salvo92"

	config.DB.Save(User)

	models.Location{}.CreateTable()

	var Location = models.Location{
		Code:"ANK",
		Description:"Ankara",
	}
	config.DB.Save(Location)

	Location = models.Location{
		Code:"IST",
		Description:"İstanbul",
	}

	config.DB.Save(Location)

	models.ExpenceTypes{}.CreateTable()

	var ExpencesTypes = models.ExpenceTypes{
		Code:"TXI",
		Description:"Taksi",
	}
	config.DB.Save(ExpencesTypes)

	ExpencesTypes = models.ExpenceTypes{
		Code:"Hotel",
		Description:"Otel",
	}
	config.DB.Save(ExpencesTypes)

	ExpencesTypes = models.ExpenceTypes{
		Code:"Ymk",
		Description:"Yeme İçme",
	}
	config.DB.Save(ExpencesTypes)

	models.Expences{}.CreateTable()
	var Expences = models.Expences{
		Expencetype:"Credit Card",
		Description:"Ödeme Detayları",
		Amount:"%18",
		VAT:"%18",
	}
	config.DB.Save(Expences)
}