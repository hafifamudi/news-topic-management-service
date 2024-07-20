package model

import (
	"github.com/google/uuid"
	"github.com/hafifamudi/news-topic-management-service/internal/general/model/common"
	"time"
)

type News struct {
	ID        uuid.UUID `gorm:"primary_key"`
	Title     string    `gorm:"type:varchar(255);unique"`
	Content   string    `gorm:"type:text"`
	Status    string    `gorm:"type:varchar(50)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Topics    []common.Topic `gorm:"many2many:news_topics;constraint:OnDelete:CASCADE;"`
}
