package productdto

type CreateProductRequest struct {
	Title    string               `json:"title" form:"title" validate:"required"`
	Price    int                  `json:"price" form:"price" validate:"required"`
	Image    string               `json:"image" form:"image" validate:"required"`
}

type UpdateProductRequest struct {
	Title    string               `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price    int                  `json:"price" form:"price" gorm:"type: int"`
	Image    string               `json:"image" form:"image" gorm:"type: varchar(255)"`
}