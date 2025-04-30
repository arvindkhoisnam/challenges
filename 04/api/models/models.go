package models

import (
	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)


type UserModel struct {
	Id		  string  `gorm:"primaryKey" json:"id"`
	Username  string  `json:"username" gorm:"unique"`
	Password  string  `json:"password"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	CreateAt  string  `json:"createAt"`
	Role	  string  `json:"role"`
}
func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = cuid.New()
	return
}

func MigrateDB(client *gorm.DB)error{
	err := client.AutoMigrate(&UserModel{})
	return err
}