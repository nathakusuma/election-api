package entity

import (
	"gorm.io/gorm"
	"time"
)

type ElectionPeriod struct {
	gorm.Model
	Start *time.Time
	End   *time.Time
}
