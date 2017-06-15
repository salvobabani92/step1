package models

import (
	"time"
	"github.com/salvobabani92/step1/config"
)

type Expences struct {
	//Harcama Tarihi
	Expencedate time. `json:"expencedate"`
	//oluşturma Tarihi
	CreatedAt time.Time `json:"-"`
	//değişim tarihi
	UpdatedAt time.Time `json:"-"`
	//harcama tipi
	Expencetype string `json:"expencetype"`
	// kullanıcı kimliği
	UserID uint `json:"user_ıd"`
	//açıklama
	Description string `json:"description"`
	//Yüzdesi
	Amount string `json:"amount"`
	//KDV
	VAT string `json:"vat"`


}
// tablo olusturma
func (this Expences) CreateTable() {
	config.DB.DropTable(this)
	config.DB.CreateTable(this)
}