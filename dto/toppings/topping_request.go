package toppingsdto

type CreateToppingRequest struct {
	Title string `gorm:"type: varchar(255)" json:"title"`
	Price int    `gorm:"type: int" json:"email"`
	Image string `gorm:"type: varchar(255)" json:"image"`
	Qty   int    `gorm:"type: int" json:"qty"`
}

type UpdateToppingRequest struct {
	Title string `gorm:"type: varchar(255)" json:"title"`
	Price int    `gorm:"type: int" json:"email"`
	Image string `gorm:"type: varchar(255)" json:"image"`
	Qty   int    `gorm:"type: int" json:"qty"`
}