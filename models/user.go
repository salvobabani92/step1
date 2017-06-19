package models

import (
	"github.com/salvobabani92/step1/config"
	"time"

)

// kullanıcı tablosudur
type User struct {
	ID         uint `json:"id"`
	// olusturma tarihi
	CreatedAt  time.Time `json:"-"`
	//değiştirme tarihi
	UpdateAt   time.Time `json:"-"`
	// kullanıcı adı
	Username   string `json:"username"`
	// şifre
	Password   string `json:"password"`
	// isim
	FirstName  string `json:"first_name"`
	// soyadı
	LastName   string `json:"last_name"`
	// lokasyonla birebir ilişkilendirme
	Location   Location `json:"location"`
	//lokasyon kimliği
	LocationID uint `json:"location_ıd"`
}
// tablo olusturma
func (this User) CreateTable() {
	config.DB.DropTable(this)
	config.DB.CreateTable(this)
}
func (this *User) Modify() {
	// Modify Fonksiyonu burada olacak.
	config.DB.Save(&this)
}

func (this User) Insert() {
	if config.DB.NewRecord(&this) {
		config.DB.Create(&this)
	}
	config.DB.NewRecord(&this)
}

func (this User) Get() {
	config.DB.First(&this, this.ID)
}

func (this *User) Delete() {
	config.DB.Delete(&this)
}



