package entity

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	Title       *string `gorm:"type:VARCHAR(64)"`
	Content     *string `gorm:"type:VARCHAR(2048)"`
	CandidateID *uint
	Candidate   *Candidate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments    []Comment
}

func (p *Post) AfterDelete(tx *gorm.DB) error {
	data := tx.Model(&Comment{}).Where("post_id = ?", p.ID)
	err := data.Update("deleted_at", time.Now()).Error
	return err
}
