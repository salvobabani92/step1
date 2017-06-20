package controller

import (
	"github.com/salvobabani92/step1/models"
	"fmt"
)

func Post_Branches() {
	//TODO:Yeni Kayıt kodu burada yazılacak
}

func Get_Branches() {
	//TODO:Tüm Kayıtlar(Liste)

}

func Get_BranchesBYID() {
	//TODO:Primary Key değerine göre ilgili kaydın getirilmesi.
}

func Put_Branches() {
	//TODO:Kaydı Güncelle
}

func Delete_Branches() {
	//TODO:Kaydı Sil
}

func Upload_Branches_From_Excel() {
	//TODO:Excel Dosyasından Kayıtlatın oluşturulması
}

func Upload_Branches_From_Json_Array() {
	//TODO:Json Array yüklenmesi ile gelen kayıtların açılması
}

func Add_Default_Records_Branches() {
	//TODO:Default Kayıt Eklemek

	var Branches = models.Branches{}
	Branches.Name = "Beyoğlu"
	Branches.Description = "Merkez Bankası"
	Branches.Code = "1101"
	Branches.Insert()

	Branches.Name = "Levent"
	Branches.Description = "Yapı ve Kredi"
	Branches.Code = "1102"
	Branches.Insert()

	Branches.Name = "Şişli"
	Branches.Description = "Akbank"
	Branches.Code = "1103"
	Branches.Insert()

	Branches.Name = "Göktürk"
	Branches.Description = "Denizbank"
	Branches.Code = "1104"
	Branches.Insert()

	Branches.ID = 4
	Branches.Get()

	Branches.Code = "1104"
	Branches.Modify()

	//Branches.Description
	//Branches.Delete()

	fmt.Println(Branches)
}
