package controller

import (
	"github.com/salvobabani92/step1/models"
	"fmt"

	"github.com/salvobabani92/step1/config"
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
	// db.Debug().Model(&Location).Update(“first_name”, “örnek”)
}

func Delete_Location() {
	//TODO:Kaydı Sil
	//db.Debug().Delete(&Location)
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

	Location.ID = 3
	Location.Get()

	Location.Description = "Bizim Ankara"
	Location.Modify()

	Location.Description
	Location.Delete()

	fmt.Println(Location)

}



