package models

import (
	"time"
	"github.com/salvobabani92/step1/config"
)

type ExpenceTypes struct{
	// kod
	Code string `json:"code"`
	//olusturma tarihi
	CreatedAt time.Time `json:"-"`
	//değiştirme tarihi
	UpdatedAt time.Time `json:"-"`
	//acıklama
	Description string `json:"description"`
}
// tablo olusturma
func (this ExpenceTypes) CreateTable() {
	config.DB.DropTable(this)
	config.DB.CreateTable(this)
}