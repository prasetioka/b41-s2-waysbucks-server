package usersdto

type UpdateUserRequest struct {
	Name     string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	Image    string `gorm:"type: varchar(255)" json:"image"`
}
