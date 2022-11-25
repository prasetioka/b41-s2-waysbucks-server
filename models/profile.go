package models

type Profile struct {
	ID     int                  `json:"id"`
	Name   string               `json:"name" gorm:"type:varchar(255)"`
	Email  string               `json:"email" gorm:"type:varchar(255)"`
	UserID int                  `json:"user_id"`
	User   UsersProfileResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// for association relation with another table (user)
type ProfileResponse struct {
	Name   string               `json:"name"`
	Email  string               `json:"email"`
	UserID int                  `json:"-"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
