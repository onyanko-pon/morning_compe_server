package Model

import (
	"time"
)

type User struct {
  ID int `json:"id"`
  Username string `json:"username" gorm:"type:varchar(255)"`
	Name string `json:"name" gorm:"type:varchar(255)"`
	TwitterUserID int `json:"twitter_user_id"`
	// Email string `json:"email" gorm:"type:varchar(255)"`
	// PasswordHash string `gorm:"->:false;<-:create" json:"-"`
	// Password string `gorm:"-" json:"-"`
	Description string `json:"description" gorm:"type:text"`
	Image string `json:"image" gorm:"type:text"`
	// Status int `json:"status"`
	AuthorizeTokenHash string `gorm:"->:false;<-:create" json:"-"`
	AuthorizeToken string `gorm:"-" json:"-"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp,autoCreateTime"`
	Updated_at time.Time `json:"updated_at" gorm:"type:timestamp,autoUpdateTime"`
	// DeletedAt time.Time `json:"deleted_at" gorm:"type:datetime"`
}

type UserBindingJson struct {
	Username string `json:"username"`
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	AuthorizeToken string `json:"authorize_token"`
}

func BindUser(user *User, user_bind_json UserBindingJson) {

	if user_bind_json.Username != "" {
		user.Username = user_bind_json.Username
	}

	if user_bind_json.Name != "" {
			user.Name = user_bind_json.Name
	}

	if user_bind_json.Description != "" {
		user.Description = user_bind_json.Description
	}

	if user_bind_json.Image != "" {
		user.Image = user_bind_json.Image
	}
}
