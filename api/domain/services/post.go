package services

import (
	"github.com/AlmazDefourten/goapp/infra/logger_instance"
	"github.com/AlmazDefourten/goapp/models/post_models"
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

type PostService struct {
}

func NewPostService() *PostService {
	return &PostService{}
}

func (postService *PostService) ListPosts() ([]post_models.Post, error) {
	var posts []post_models.Post
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

func (postService *PostService) CreatePost(post post_models.Post) error {
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

func (postService *PostService) GetPost(id int) (post_models.Post, error) {
	var db gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		return post_models.Post{}, err
	}
	var post post_models.Post
	err = db.Where("id = ?", id).First(post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}
