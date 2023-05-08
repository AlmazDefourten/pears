package services

import (
	"github.com/AlmazDefourten/goapp/infrastructure/logger_instance"
	"github.com/AlmazDefourten/goapp/models"
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

type PostService struct {
}

func NewPostService() *PostService {
	return &PostService{}
}

func (postService *PostService) ListPosts() ([]models.Post, error) {
	var posts []models.Post
	var db gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		return nil, err
	}
	err = db.Find(&posts).Error // TODO: after friends and community-projects subscribe realization add where with filter on it
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (postService *PostService) CreatePost(post models.Post) error {
	var db gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		return err
	}
	err = db.Create(&post).Error
	if err != nil {
		// logging here or on another level after return // TODO: think about logging level
		logger_instance.ServiceLogger.Error(err)
		return err
	}
	return nil
}

func (postService *PostService) GetPost(id int) (models.Post, error) {
	var db gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		return models.Post{}, err
	}
	var post models.Post
	err = db.Where("id = ?", id).First(post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}
