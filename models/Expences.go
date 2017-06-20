package models

import (
	"time"
	"github.com/salvobabani92/step1/config"
	"fmt"
)

type Expences struct {
	ID          uint `json:"id"`
	//Harcama Tarihi
	Expencedate time.Time `json:"expencedate"`
	//oluşturma Tarihi
	CreatedAt   time.Time `json:"-"`
	//değişim tarihi
	UpdatedAt   time.Time `json:"-"`
	//harcama tipi
	Expencetype string `json:"expencetype"`
	// kullanıcı kimliği
	UserID      uint `json:"user_ıd"`
	//açıklama
	Description string `json:"description"`
	//Yüzdesi
	Amount      string `json:"amount"`
	//KDV
	VAT         string `json:"vat"`
	//Lokasyon baglantısında kimlik
	LocationID  uint `json:"location_ıd"`
	//many to many baglantısı için
	Attendees   []User `json:"many2many:expences_user"`
}
// tablo olusturma
func (this Expences) CreateTable() {
	config.DB.DropTable(this)
	config.DB.CreateTable(this)
}

func (this Expences) Insert() {
	if config.DB.NewRecord(&this) {
		config.DB.Create(&this)
	}
	config.DB.NewRecord(&this)
}

func (this Expences) Get() {
	config.DB.First(&this, this.ID)
	fmt.Println(this.Description)
}

func (this *Expences) Modify() {
	// Modify Fonksiyonu burada olacak.
	config.DB.Save(&this)
}

func (this *Expences) Delete() {
	config.DB.Delete(&this)
}