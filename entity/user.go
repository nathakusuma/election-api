package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string  `gorm:"type:VARCHAR(64); NOT NULL"`
	Username string  `gorm:"unique; type:VARCHAR(32); NOT NULL;"`
	Password string  `gorm:"type:VARCHAR(128); NOT NULL"`
	Bio      *string `gorm:"type:VARCHAR(256)"`
	Role     string  `gorm:"type:ENUM('admin', 'candidate', 'user'); NOT NULL; DEFAULT:'user'"`
	CanVote  bool    `gorm:"NOT NULL; DEFAULT:true"`
}

func (u *User) AfterDelete(tx *gorm.DB) error {
	data := tx.Model(&Candidate{}).Where("user_id = ?", u.ID)
	err := data.Update("deleted_at", time.Now()).Error
	return err
}
