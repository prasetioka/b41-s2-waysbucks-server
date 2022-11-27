package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Role     string `json:"role" gorm:"type: varchar(255)"`
	Image    string `json:"image" gorm:"type: varchar(255)"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"fullname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Image    string `json:"image"`
}

func (UserResponse) TableName() string {
	return "users"
}