package usersdto

type UserResponse struct {
	Name     string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Image    string `json:"image"`
}