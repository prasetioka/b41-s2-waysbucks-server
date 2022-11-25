package profiledto

import "waysbucks-api/models"

type ProfileResponse struct {
	ID     int                         `json:"id" gorm:"primary_key:auto_increment"`
	Name   string                      `json:"phone" gorm:"type: varchar(255)"`
	Email  string                      `json:"gender" gorm:"type: varchar(255)"`
	UserID int                         `json:"user_id"`
	User   models.UsersProfileResponse `json:"user"`
}
