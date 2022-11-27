package models

type Topping struct {
	ID     int          `json:"id" gorm:"primary_key:auto_increment"`
	Title  string       `json:"title" gorm:"type: varchar(255)"`
	Price  int          `json:"price" gorm:"type: int"`
	Image  string       `json:"image" gorm:"type: varchar(255)"`
	Qty    int          `json:"qty" gorm:"type: varchar(255)"`
}