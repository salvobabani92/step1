package models

import (
	"time"
	"github.com/salvobabani92/step1/config"

	"fmt"
)

// Şubeler
type Location struct {
	ID          uint `json:"id"`
	// kod
	Code        string `json:"code"`
	// olusturma tarihi
	CreatedAt   time.Time `json:"-"`
	//değiştirme tarihi
	UpdateAt    time.Time `json:"-"`
	// açıklama girişi
	Description string `json:"description"`
}
/*

	Model için default CRUD işlemleri kendisi içinde tanımlanabilir.

	C:	Create
	R:	Read  	==> GET
	U:	Update	==> Modify
	D:	Delete


 */



// Database üzerinde tablo yeniden oluşturulur.
func (this Location)CreateTable() {
	config.DB.DropTable(this)
	config.DB.CreateTable(this)
}

func (this *Location) Modify()  {
	// Modify Fonksiyonu buraa olacak.
	config.DB.Save(&this)
}

func (this Location) Insert() {
	if config.DB.NewRecord(&this) {
		config.DB.Create(&this)
	}
	config.DB.NewRecord(&this)
}

func (this *Location) Get() {
	config.DB.First(&this, this.ID)
}

func (this *Location) Delete() {
	config.DB.Delete(&this)
}