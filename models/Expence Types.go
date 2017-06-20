package models

import (
	"time"
	"github.com/salvobabani92/step1/config"
	"fmt"
)

type ExpenceTypes struct {
	ID          uint `json:"id"`
	// kod
	Code        string `json:"code"`
	//olusturma tarihi
	CreatedAt   time.Time `json:"-"`
	//değiştirme tarihi
	UpdatedAt   time.Time `json:"-"`
	//acıklama
	Description string `json:"description"`
}
// tablo olusturma
func (this ExpenceTypes) CreateTable() {
	config.DB.DropTable(this)
	config.DB.CreateTable(this)
}

func (this ExpenceTypes) Insert() {
	if config.DB.NewRecord(&this) {
		config.DB.Create(&this)
	}
	config.DB.NewRecord(&this)

}
func (this ExpenceTypes) Get() {
	config.DB.First(&this, this.ID)
	fmt.Println(this.Description)
}

func (this *ExpenceTypes) Modify() {
	// Modify Fonksiyonu burada olacak.
	config.DB.Save(&this)
}

func (this *ExpenceTypes) Delete() {
	config.DB.Delete(&this)
}