package Model

import (
	"time"
	// "github.com/jinzhu/gorm"
)

type Event struct {
  ID int `json:"id"`
  Title string `json:"title",gorm:"type:varchar(255)"`
	Body string `json:"body",gorm:"type:text"`
	HostUserID int `json:"host_user_id",gorm:"type:integer"`
	Place string `json:"place",gorm:"type:varchar(255)"`
	Date time.Time `json:"date",gorm:"type:date"`
	StartTime time.Time `json:"start_time",gorm:"type:time"`
	EndTime time.Time `json:"end_time",gorm:"type:time"`
	HostUser *User `gorm:"foreignKey:HostUserID"`
	CreatedAt time.Time `json:"crated_at",gorm:"type:timestamp,autoCreateTimex"`
	UpdatedAt time.Time `json:"updated_at",gorm:"type:timestamp,autoUpdateTime"`
	Users []*User `gorm:"many2many:events_users;"`
}
