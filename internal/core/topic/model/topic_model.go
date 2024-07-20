package model

import (
	"github.com/google/uuid"
	"news-topic-management-service/internal/general/model/common"
	"time"
)

type Topic struct {
	ID        uuid.UUID `gorm:"primary_key"`
	Name      string    `gorm:"type:varchar(255);unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	News      []common.News `gorm:"many2many:news_topics;constraint:OnDelete:CASCADE;"`
}
