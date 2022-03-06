// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/goxiaoy/go-saas-kit/pkg/api"
	"github.com/goxiaoy/go-saas-kit/pkg/authn/jwt"
	"github.com/goxiaoy/go-saas-kit/pkg/conf"
	api2 "github.com/goxiaoy/go-saas-kit/saas/api"
	"github.com/goxiaoy/go-saas-kit/saas/remote"
	api3 "github.com/goxiaoy/go-saas-kit/user/api"
	remote2 "github.com/goxiaoy/go-saas-kit/user/remote"
	"github.com/goxiaoy/go-saas/common/http"
)

import (
	_ "github.com/goxiaoy/go-saas-kit/gateway/apisix/cmd/go-runner/plugins"
	_ "github.com/goxiaoy/go-saas/gateway/apisix"
)

// Injectors from wire.go:

func initApp(services *conf.Services, security *conf.Security, clientName api.ClientName, logger log.Logger, arg ...grpc.ClientOption) (*App, func(), error) {
	option := NewSelfClientOption(logger)
	tokenizerConfig := jwt.NewTokenizerConfig(security)
	tokenizer := jwt.NewTokenizer(tokenizerConfig)
	inMemoryTokenManager := api.NewInMemoryTokenManager(tokenizer, logger)
	grpcConn, cleanup := api2.NewGrpcConn(clientName, services, option, inMemoryTokenManager, logger, arg...)
	tenantServiceClient := api2.NewTenantGrpcClient(grpcConn)
	tenantStore := remote.NewRemoteGrpcTenantStore(tenantServiceClient)
	apiGrpcConn, cleanup2 := api3.NewGrpcConn(clientName, services, option, inMemoryTokenManager, logger, arg...)
	userServiceClient := api3.NewUserGrpcClient(apiGrpcConn)
	userTenantContributor := remote2.NewUserTenantContributor(userServiceClient)
	webMultiTenancyOption := http.NewDefaultWebMultiTenancyOption()
	app, err := newApp(tenantStore, userTenantContributor, tokenizer, inMemoryTokenManager, services, clientName, webMultiTenancyOption, security, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
