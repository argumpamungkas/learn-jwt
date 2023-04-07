package models

import (
	"DTS/Chapter-3/sesi/sesi2-go-jwt/helpers"
	"log"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Tag valid untuk menentukan validasi terhadap field menggunakan package validator, diaktifkan pada hooks func BeforeCreate dari orm Gorm
type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) TableName() string {
	return "tb_users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u) // akan mengarah pada field valid

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password, err = helpers.HashPass(u.Password)
	if err != nil {
		log.Println("error hash password")
		return
	}

	return
}
