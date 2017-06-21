package main

import (
	"testing"
	"fmt"
	"github.com/salvobabani92/step1/models"
	"github.com/salvobabani92/step1/config"
)

func TestInitDB(t *testing.T) {

	fmt.Println("Test Init DB Process")
	config.InitDB()
	models.User{}.CreateTable()
	models.Location{}.CreateTable()
	models.Expense{}.CreateTable()
	models.ExpenseType{}.CreateTable()
	models.Branche{}.CreateTable()

}



/*

	Model testinde hazırlanacak standart test prosedürü

	1. Toplam 5 sahte Kayıt ekle. Toplam kayıt sayısı 5 değilse hata ver.
	2. 3. kaydı çağır eğer gelen kayıt eklediğin 3 numaralı kayıt değilse error ver.
	3. 3 numaralı Kaydı değiştir.
	4. 3 numaralı kaydı tekrar çağır ve son değiştirdiğin gibi değilse error ver.
	5. 3 numaralı kaydı tekrar çağır ve sil. Sonrasında 3 numaralı kaydı tekrar çağır kayıt geliyorsa hata ver.



 */
func Test_Branche(t *testing.T) {
	fmt.Println("Branche Test Process")
	var Branche = models.Branche{}
	// 5 Sahte kayıt açıldı.
	Branche.Name = "Esentepe"
	Branche.Code = "112"
	Branche.Insert()

	Branche.Name = "Levent"
	Branche.Code = "122"
	Branche.Insert()

	Branche.Name = "Şişli"
	Branche.Code = "131"
	Branche.Insert()

	Branche.Name = "Göktürk"
	Branche.Code = "141"
	Branche.Insert()

	Branche.Name = "Eyüp"
	Branche.Code = "153"
	Branche.Insert()
	//Toplam kayıt sayısı 5 değilse error ver.
	//if Branche(count != 5) {
		//t.Error("Expected 5, got", count)
	//}
	//3 numaralı kaydı çağır. 3 numaralı kayıt gelmezse error ver.
	Branche.ID = 3
	Branche.Get()

	if Branche.Name != "Şişli" {
		t.Error("Şişli çıkmadı ya.")
	}

	// 3 numaralı kaydı değiştir.
	Branche.Code = "131"
	Branche.Modify()

	//3 numaralı kaydı tekrar çagır ve değiştirdiğin gibi değilse error ver.

	//3 numaralı kaydı tekrar çağır ve sil. Sonra 3 numaralı kaydı tekrar çağır kayıt geliyorsa hata ver.
	//Branche.Code
	//Branche.Delete()

}

func Test_ExpenseTypes(t *testing.T) {

	fmt.Println("ExpenseTypes Test Process")
	var ExpenseTypes = models.ExpenseType{}
	// 5 sahte kayıt açıldı.
	ExpenseTypes.Description = "Yeme İçme"
	ExpenseTypes.Code = "Ymk"
	ExpenseTypes.Insert()

	ExpenseTypes.Description = "Giyim"
	ExpenseTypes.Code = "Gym"
	ExpenseTypes.Insert()

	ExpenseTypes.Description = "Alışveriş"
	ExpenseTypes.Code = "Alş"
	ExpenseTypes.Insert()

	ExpenseTypes.Description = "Kişisel"
	ExpenseTypes.Code = "Kşl"
	ExpenseTypes.Insert()

	ExpenseTypes.Description = "Akaryakıt"
	ExpenseTypes.Code = "bnz"
	ExpenseTypes.Insert()
	// toplam kayıt sayısı 5 değilse hata ver.
	//if ExpenseTypes(count != 5) {
		//t.Error("Expected 5, got", count)
	//}

	// 3 numaralı kaydı çağır.Gelmezse error ver.
	ExpenseTypes.ID = 3
	ExpenseTypes.Get()

	if ExpenseTypes.Code != "Alş" {
		t.Error("Alş çıkmadı.")
	}

	// 3 numaralı kaydı değiştir.
	ExpenseTypes.Code = "Alş"
	ExpenseTypes.Modify()

	//3 numaralı kaydı tekrar çağır değiştiği gibi gelmezse error ver.

	//3 numaralı kaydı tekrar çağır ve sil. Sonra 3 numaralı kaydı tekrar çağır kayıt geliyorsa hata ver.
	ExpenseTypes.Code = "Alş"
	ExpenseTypes.Delete()

}

func Test_Expense(t *testing.T) {

	fmt.Println("Expense Test Process")
	var Expense = models.Expense{}
	// 5 sahte kayıt açıldı.
	Expense.Amount = 30
	Expense.Description = "Giyim alışveriş harcamaları"
	Expense.VAT = 3
	Expense.Insert()

	Expense.Amount = 44
	Expense.Description = "Teknolojik digital eşya harcamaları"
	Expense.VAT = 3
	Expense.ExpenseTypeID = 2
	Expense.Insert()

	Expense.Amount = 46
	Expense.Description = "Yeme içme gıda harcamaları"
	Expense.VAT = 4
	Expense.Insert()

	Expense.Amount = 55
	Expense.Description = "Günlük kişisel harcamalar"
	Expense.VAT = 4
	Expense.Insert()

	Expense.Amount = 100
	Expense.VAT = 18
	Expense.Description = "Okul harcamaları"
	Expense.Insert()
	// toplam kayıt sayısı 5 değilse error ver.
	//if Expense(count != 5) {
		//t.Error("Expected 5, got", count)
	//}

	// 3 numaralı kaydı çağır gelmezse error ver.
	Expense.ID = 3
	Expense.Get()

	//3 numaralı kaydı çağır ve değiştir.
	Expense.VAT = 32
	Expense.Modify()

	//3 numaralı kaydı çağır değişmemişse error ver.

	//3 numaralı kaydı çağır ve sil.Sonra 3 numaralı kaydı tekrar çağır gelirse error ver.
	Expense.VAT = 32
	Expense.Delete()

}

func Test_Location(t *testing.T) {

	fmt.Println("Location Test Process")
	var Location = models.Location{}
	//5 sahte kayıt ekle.
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
	//toplam kayıt sayısı 5 değilse error ver.
	//if Location(count != 5) {
		//t.Error("Expected 5, got", count)
	//}

	// 3 numaralı kaydı çağır gelmezse error ver.
	Location.ID = 3
	Location.Get()

	//3 numaralı kaydı çağır ve değiştir.
	Location.Code = "Brs"
	Location.Modify()
	//3 numaralı kaydı çağır değişmemişse error ver.

	//3 numaralı kaydı çağır ve sil.Sonrasında 3 numaralı kaydı çağır gelirse error ver.
	Location.Description = "Bursa"
	Location.Delete()
}

func Test_user(t *testing.T) {

	fmt.Println("user Test Process")
	var user = models.User{}
	// 5 sahte kayıt ekle.
	user.FirstName = "ibrahim"
	user.LastName = "cobani"
	user.LocationID = 1
	user.UserName = "icobani"
	user.Password = "12345"
	user.Insert()

	user.FirstName = "salvator"
	user.LastName = "babani"
	user.LocationID = 1
	user.UserName = "sbabani"
	user.Password = "123453"
	user.Insert()

	user.FirstName = "ismail"
	user.LastName = "ünsal"
	user.LocationID = 1
	user.UserName = "iünsal"
	user.Password = "12342"
	user.Insert()

	user.FirstName = "hakkı"
	user.LastName = "sağlam"
	user.LocationID = 1
	user.UserName = "hsağlam"
	user.Password = "12315"
	user.Insert()

	user.FirstName = "faruk"
	user.LastName = "bayar"
	user.LocationID = 1
	user.UserName = "fbayar"
	user.Password = "13245"
	user.Insert()

	//toplam kayıt sayısı 5 değilse error ver.
	//if user(count != 5) {
		//t.Error("Expected 5, got", count)
	//}

	// 3 numaralı kaydı çağır.Gelmezse error ver.
	user.ID = 3
	user.Get()

	//3 numaralı kaydı çağır ve değiştir.
	user.LastName = "ünsal"
	user.Modify()
	//3 numaralı kaydı çağır değişmemişse error ver.

	//3 numaralı kaydı çağır ve sil.Sonrasında 3 numaralı kaydı çağır gelirse error ver.
	user.UserName = "iünsal"
	user.Delete()
}