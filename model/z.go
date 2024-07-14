package model

import "time"

type Translate struct {
	ID             uint      `gorm:"primaryKey"`
	OriginalText   string    `gorm:"not null"`
	TranslatedText string    `gorm:"not null"`
	Direction      string    `gorm:"not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}

func (t *Translate) Create() error {
	return DB.Create(t).Error
}

func SearchTranslates() (ts []*Translate, err error) {
	err = DB.Find(&ts).Error
	return ts, err
}
