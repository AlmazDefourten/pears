package tag_models

import "github.com/AlmazDefourten/goapp/models/post_models"

type Tag struct {
	Id       uint                 `json:"id" gorm:"primaryKey"`
	Name     string               `json:"tag"`
	Hash     string               `json:"hash"`
	TagPosts post_models.PostTags `gorm:"foreignKey:TagId"`
}
