package gorm

import (
	"gorm.io/gorm"
)

type GormDBWrapper struct {
	conn *gorm.DB
}

func NewGormDBWapper(conn *gorm.DB) *GormDBWrapper {
	return &GormDBWrapper{
		conn: conn,
	}
}
