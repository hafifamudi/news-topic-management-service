package common

import (
	"github.com/google/uuid"
	"time"
)

type NewsTopic struct {
	ID      uint      `gorm:"primary_key"`
	NewsID  uuid.UUID `gorm:"not null"`
	TopicID uuid.UUID `gorm:"not null"`
}

type News struct {
	ID        uuid.UUID `gorm:"primary_key"`
	Title     string    `gorm:"type:varchar(255)"`
	Content   string    `gorm:"type:text"`
	Status    string    `gorm:"type:varchar(50)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Topics    []Topic `gorm:"many2many:news_topics"`
}

type Topic struct {
	ID        uuid.UUID `gorm:"primary_key"`
	Name      string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	News      []News `gorm:"many2many:news_topics"`
}
