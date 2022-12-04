package productsdto

type CreateProductRequest struct {
	Title string `gorm:"type: varchar(255)" form:"title" json:"title"`
	Price int    `gorm:"type: int" form:"price" json:"email"`
	Image string `gorm:"type: varchar(255)" form:"image" json:"image"`
	Qty   int    `gorm:"type: int" form:"qty" json:"qty"`
}

type UpdateProductRequest struct {
	Title string `gorm:"type: varchar(255)" json:"title"`
	Price int    `gorm:"type: int" json:"email"`
	Image string `gorm:"type: varchar(255)" json:"image"`
	Qty   int    `gorm:"type: int" json:"qty"`
}
