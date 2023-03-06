package models

type Tag struct {
	Id       uint     `json:"id" gorm:"primaryKey"`
	Name     string   `json:"tag"`
	Hash     string   `json:"hash"`
	TagPosts PostTags `gorm:"foreignKey:TagId"`
}
