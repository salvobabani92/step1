package models

import (
	"github.com/salvobabani92/step1/config"
	"time"
)

// kullanıcı tablosudur
type User struct {
	// kullanıcı numarası
	ID uint `gorm:"primary_key" json:"ıd"`
	// olusturma tarihi
	CreatedAt time.Time `json:"-"`
	//değiştirme tarihi
	UpdateAt time.Time `json:"-"`
	// kullanıcı adı
	Username string `json:"username"`
	// şifre
	Password string `json:"password"`
	// isim
	Name     string `json:"name"`
	// soyadı
	Surname  string `json:"surname"`
}
// tablo olusturma
func (this User) CreateTable() {
config.DB.DropTable(this)
config.DB.CreateTable(this)
}


