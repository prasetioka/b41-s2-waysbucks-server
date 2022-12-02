package authdto

type LoginResponse struct {
	Name  string `gorm:"type: varchar(255)" json:"fullname"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	Token string `gorm:"type: varchar(255)" json:"token"`
	Role string `gorm:"type: varchar(50)"  json:"role"`
}

type RegisterResponse struct {
	Name  string `json:"fullname"`
	Email string `json:"email"`
}

type CheckAuthResponse struct {
	Id     int    `gorm:"type: int" json:"id"`
	Name   string `gorm:"type: varchar(255)" json:"name"`
	Email  string `gorm:"type: varchar(255)" json:"email"`
	Role string `gorm:"type: varchar(50)"  json:"role"`
	Token string `gorm:"type: varchar(255)" json:"token"`
}