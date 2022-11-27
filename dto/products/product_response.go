package productsdto

type ProductResponse struct {
	Title string `json:"title"`
	Price int `json:"price"`
	Image string `json:"image"`
	Qty   int `json:"qty"`
}
