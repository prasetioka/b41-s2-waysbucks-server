package usersdto

type UserResponse struct {
	Name     string `json:"fullname"`
	Email    string `json:"email"`
	Image    string `json:"image"`
}