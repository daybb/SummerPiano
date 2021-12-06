// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"helloworld/internal/biz"
	"helloworld/internal/conf"
	"helloworld/internal/dao"
	"helloworld/internal/server"
	"helloworld/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, dao.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
