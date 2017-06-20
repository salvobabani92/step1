package controller

import (
	"github.com/salvobabani92/step1/models"
	"fmt"


)

func Post_Location() {
	//TODO:Yeni Kayıt kodu burada yazılacak

}

func Get_Location() {
	//TODO:Tüm Kayıtlar(Liste)


}

func Get_LocationBYID() {
	//TODO:Primary Key değerine göre ilgili kaydın getirilmesi.

}

func Put_Location() {
	//TODO:Kaydı Güncelle

}

func Delete_Location() {
	//TODO:Kaydı Sil

}

func Upload_Location_From_Excel() {
	//TODO:Excel Dosyasından Kayıtlatın oluşturulması
}

func Upload_Location_From_Json_Array() {
	//TODO:Json Array yüklenmesi ile gelen kayıtların açılması
}

func Add_Default_Records_Locations() {
	//TODO:Default Kayıt Eklemek

	var Location = models.Location{}
	Location.Code = "İzm"
	Location.Description = "İzmir"
	Location.Insert()

	Location.Code = "İst"
	Location.Description = "İstanbul"
	Location.Insert()

	Location.Code = "Brs"
	Location.Description = "Bursa"
	Location.Insert()

	Location.Code = "Ank"
	Location.Description = "Ankara"
	Location.Insert()

	Location.Code = "Adn"
	Location.Description = "Adana"
	Location.Insert()

	Location.ID = 1
	Location.Get()

	Location.Code = "İzm"
	Location.Modify()

	//Location.Description
	//Location.Delete()

	fmt.Println(Location)

}



