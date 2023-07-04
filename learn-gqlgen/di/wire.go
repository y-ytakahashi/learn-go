//go:build wireinject
// +build wireinject

package di

import (
	"learn-gqlgen/domain/service"
	"learn-gqlgen/infra/datastore"
	"learn-gqlgen/infra/datastore/gorm"

	"github.com/google/wire"
)

func InitializeApp() App {
	wire.Build(

		datastore.NewMySqlDB,
		gorm.NewGormDBWapper,
		NewApp,
		service.WireSet,
	)
	return App{}
}
