package models

import (
	"gorm.io/gorm"
)

type Feed struct {
	gorm.Model
	AuthorId int
	Title    string
	Secret   string
}

type Form struct {
	gorm.Model
	FeedId int
	Feed   Feed
}
