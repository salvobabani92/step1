package controller

import (
	"github.com/salvobabani92/step1/models"
	"fmt"


)

func Post_ExpenceTypes() {
	//TODO:Yeni Kayıt kodu burada yazılacak
}

func Get_ExpenceTypes() {
	//TODO:Tüm Kayıtlar(Liste)

}

func Get_ExpenceTypesBYID() {
	//TODO:Primary Key değerine göre ilgili kaydın getirilmesi.
}

func Put_ExpenceTypes() {
	//TODO:Kaydı Güncelle
}

func Delete_ExpenceTypes() {
	//TODO:Kaydı Sil
}

func Upload_ExpenceTypes_From_Excel() {
	//TODO:Excel Dosyasından Kayıtlatın oluşturulması
}

func Upload_ExpenceTypes_From_Json_Array() {
	//TODO:Json Array yüklenmesi ile gelen kayıtların açılması
}

func Add_Default_Records_ExpenceTypes() {
	//TODO:Default Kayıt Eklemek

	var ExpenceTypes = models.ExpenceTypes{}
	ExpenceTypes.Description = "yeme içme"
	ExpenceTypes.Code = "ymk"
	ExpenceTypes.Insert()

	ExpenceTypes.Description = "giyim"
	ExpenceTypes.Code = "gym"
	ExpenceTypes.Insert()

	ExpenceTypes.Description = "alışveriş"
	ExpenceTypes.Code = "alş"
	ExpenceTypes.Insert()

	ExpenceTypes.Description = "benzin"
	ExpenceTypes.Code = "bnz"


	ExpenceTypes.ID = 4
	ExpenceTypes.Get()

	ExpenceTypes.Code= "bnz"
	ExpenceTypes.Modify()

	//ExpenceTypes.Description
	//ExpenceTypes.Delete()

	fmt.Println(ExpenceTypes)
}
