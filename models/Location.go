package models

import (
	"time"
	"github.com/salvobabani92/step1/config"
)

// Şubeler
type Location struct {
	// kod
	Code        string `json:"code"`
	// olusturma tarihi
	CreatedAt   time.Time `json:"-"`
	//değiştirme tarihi
	UpdateAt    time.Time `json:"-"`
	// açıklama girişi
	Description string `json:"description"`
}
// tablo oluşturma
func (this Location)CreateTable() {
	config.DB.DropTable(this)
	config.DB.CreateTable(this)
}
