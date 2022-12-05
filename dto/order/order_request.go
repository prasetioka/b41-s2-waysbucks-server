package orderdto

type CreateOrder struct {
	ID        int   `json:"id"`
	Price     int   `json:"price"`
	ProductID int   `json:"product_id"`
	ToppingID []int `json:"topping_id"`
	BuyerID   int   `json:"buyer_id"`
	Qty       int   `json:"qty"`
}

type UpdateOrder struct {
	Qty int `json:"qty" validate:"required"`
}
