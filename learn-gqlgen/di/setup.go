package di

import "learn-gqlgen/infra/datastore/gorm"

type App struct {
	DB *gorm.GormDBWrapper
}

func NewApp(DB *gorm.GormDBWrapper) App {
	return App{
		DB: DB,
	}
}
