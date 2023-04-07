package models

import (
	"log"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Title       string `json:"title" form:"title" valid:"required-Title of your product is required"`
	Desceiption string `json:"description" form:"description" valid:"required-Description of your product is required"`
	UserID      uint
	User        *User
}

func (u *Product) TableName() string {
	return "tb_products"
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		log.Fatal("error pada before create product")
		return
	}

	err = nil
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		log.Fatal("error pada before update product")
		return
	}

	err = nil
	return
}
