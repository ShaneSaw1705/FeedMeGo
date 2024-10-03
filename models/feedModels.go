package models

import (
	"gorm.io/gorm"
)

type Feed struct {
	gorm.Model
	authorId string
	title    string
	secret   string
}

type Form struct {
	gorm.Model
	feedId int
	feed   Feed
}
