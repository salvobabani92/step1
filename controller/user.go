package controller

import (
	"github.com/salvobabani92/step1/models"
	"fmt"
	"github.com/salvobabani92/step1/config"
)

func Post_User() {
	//TODO:Yeni Kayıt kodu burada yazılacak
}

func Get_User() {
	//TODO:Tüm Kayıtlar(Liste)

}

func Get_UserBYID() {
	//TODO:Primary Key değerine göre ilgili kaydın getirilmesi.


}


func Put_User() {
	//TODO:Kaydı Güncelle
}

func Delete_User() {
	//TODO:Kaydı Sil
}

func Upload_User_From_Excel() {
	//TODO:Excel Dosyasından Kayıtlatın oluşturulması
}

func Upload_User_From_Json_Array() {
	//TODO:Json Array yüklenmesi ile gelen kayıtların açılması
}

func Add_Default_Records_User() {
	//TODO:Default Kayıt Eklemek

	var user = models.User{}
	user.FirstName = "ibrahim"
	user.LastName= "cobani"
	user.LocationID = 1
	user.Username = "icobani"
	user.Password = "12345"

	fmt.Print(user)

	if config.DB.NewRecord(&user) {
		config.DB.Create(&user)
	}
	fmt.Println(user)
	config.DB.NewRecord(&user)
	fmt.Println(user)


	user = models.User{}
	user.FirstName= "salvator"
	user.LastName = "babani"
	user.LocationID = 1
	user.Username = "sbabani"
	user.Password = "123453"

	if config.DB.NewRecord(&user) {
		config.DB.Create(&user)
	}
	config.DB.NewRecord(&user)


	fmt.Println("Users added")
}
