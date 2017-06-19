package controller

import (
	"github.com/salvobabani92/step1/models"
	"fmt"
	"github.com/salvobabani92/step1/config"
	"os/user"
)

func Post_Expences() {
	//TODO:Yeni Kayıt kodu burada yazılacak
}

func Get_Expences() {
	//TODO:Tüm Kayıtlar(Liste)

}

func Get_ExpencesBYID() {
	//TODO:Primary Key değerine göre ilgili kaydın getirilmesi.
}

func Put_Expences() {
	//TODO:Kaydı Güncelle
}

func Delete_Expences() {
	//TODO:Kaydı Sil
}

func Upload_Expences_From_Excel() {
	//TODO:Excel Dosyasından Kayıtlatın oluşturulması
}

func Upload_Expences_From_Json_Array() {
	//TODO:Json Array yüklenmesi ile gelen kayıtların açılması
}

func Add_Default_Records_Expences() {
	//TODO:Default Kayıt Eklemek

	var Expences = models.Expences{}
	Expences.Amount = "%18"
	Expences.Description = "Clothes"
	Expences.VAT = "20"
	Expences.Expencetype = "Visa"
	Expences.Insert()

	Expences.Amount = "%18"
	Expences.Description = "Tecnology"
	Expences.VAT = "15"
	Expences.Expencetype = "Master Card"
	Expences.Insert()

	Expences.Amount = "%8"
	Expences.Description = "Food"
	Expences.VAT = "%2"
	Expences.Expencetype = "Cash"
	Expences.Insert()

	Expences.ID = 3
	Expences.Get()

	Expences.VAT = "%20"
	Expences.Modify()

	Expences.Description
	Expences.Delete()

	fmt.Println(Expences)
}
