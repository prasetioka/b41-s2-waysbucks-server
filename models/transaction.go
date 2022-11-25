package models

type Transaction struct {
	ID int `json:"id" gorm:"primary_key:auto_increment"`
	UserID int `json:"user_id" form:"user_id"`
	User UsersProfileResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductID int `json:"product_id" form:"product_id"`
	Product ProductUserResponse `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}