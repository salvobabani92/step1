package models

import (
	"time"
	"github.com/salvobabani92/step1/config"
	"fmt"
)

type Branches struct {
	ID          uint `json:"id"`
	// kod
	Name        string `json:"name"`
	//olusturma tarihi
	CreatedAt   time.Time `json:"-"`
	//değiştirme tarihi
	UpdatedAt   time.Time `json:"-"`
	//kod
	Code        string `json:"code"`
	//acıklama
	Description string `json:"description"`
}
// tablo olusturma
func (this Branches) CreateTable() {
	config.DB.DropTable(this)
	config.DB.CreateTable(this)
}

func (this Branches) Insert() {
	if config.DB.NewRecord(&this) {
		config.DB.Create(&this)
	}
	config.DB.NewRecord(&this)

}
func (this Branches) Get() {
	config.DB.First(&this, this.ID)
	fmt.Println(this.Description)
}

func (this *Branches) Modify() {
	// Modify Fonksiyonu burada olacak.
	config.DB.Save(&this)
}

func (this *Branches) Delete() {
	config.DB.Delete(&this)
}