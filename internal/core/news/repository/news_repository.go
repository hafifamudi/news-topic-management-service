package repository

import (
	"github.com/google/uuid"
	"github.com/hafifamudi/news-topic-management-service/internal/core/news/model"
	"github.com/hafifamudi/news-topic-management-service/pkg/infrastructure/db"
	"gorm.io/gorm"
)

type NewsRepository interface {
	GetAll() ([]model.News, error)
	Create(news *model.News) (*model.News, error)
	Update(id uuid.UUID, news *model.News) (*model.News, error)
	Find(id uuid.UUID) (*model.News, error)
	Delete(newsID uuid.UUID) (*model.News, error)
	Preload(news *model.News) (*model.News, error)
}

type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return &newsRepository{
		db: db,
	}
}

func News() NewsRepository {
	return NewNewsRepository(db.DB)
}

func (r *newsRepository) GetAll() ([]model.News, error) {
	var news []model.News
	result := r.db.Find(&news)
	if result.Error != nil {
		return nil, result.Error
	}

	return news, nil
}

func (r *newsRepository) Create(news *model.News) (*model.News, error) {
	if news.ID == uuid.Nil {
		news.ID = uuid.New()
	}

	result := r.db.Create(&news)
	if result.Error != nil {
		return nil, result.Error
	}

	return news, nil
}

func (r *newsRepository) Find(id uuid.UUID) (*model.News, error) {
	var news model.News
	result := r.db.First(&news, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &news, nil
}

func (r *newsRepository) Update(id uuid.UUID, news *model.News) (*model.News, error) {
	result := r.db.Model(&model.News{}).Where("id = ?", id).Updates(news)
	if result.Error != nil {
		return nil, result.Error
	}

	err := r.db.Model(news).Association("Topics").Replace(news.Topics)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (r *newsRepository) Delete(newsID uuid.UUID) (*model.News, error) {
	var news model.News
	result := r.db.Where("id = ?", newsID).First(&news)
	if result.Error != nil {
		return nil, result.Error
	}

	result = r.db.Delete(&news)
	if result.Error != nil {
		return nil, result.Error
	}

	return &news, nil
}

func (r *newsRepository) Preload(news *model.News) (*model.News, error) {
	if err := r.db.Preload("Topics").Find(&news).Error; err != nil {
		return nil, err
	}

	return news, nil
}
