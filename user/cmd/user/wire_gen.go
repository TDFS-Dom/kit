// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/goxiaoy/go-saas-kit/pkg/api"
	"github.com/goxiaoy/go-saas-kit/pkg/authn/jwt"
	"github.com/goxiaoy/go-saas-kit/pkg/authz/authorization"
	"github.com/goxiaoy/go-saas-kit/pkg/authz/casbin"
	"github.com/goxiaoy/go-saas-kit/pkg/conf"
	"github.com/goxiaoy/go-saas-kit/pkg/server"
	uow2 "github.com/goxiaoy/go-saas-kit/pkg/uow"
	"github.com/goxiaoy/go-saas-kit/user/private/biz"
	conf2 "github.com/goxiaoy/go-saas-kit/user/private/conf"
	"github.com/goxiaoy/go-saas-kit/user/private/data"
	"github.com/goxiaoy/go-saas-kit/user/private/seed"
	server2 "github.com/goxiaoy/go-saas-kit/user/private/server"
	http2 "github.com/goxiaoy/go-saas-kit/user/private/server/http"
	"github.com/goxiaoy/go-saas-kit/user/private/service"
	"github.com/goxiaoy/go-saas/common/http"
	"github.com/goxiaoy/go-saas/gorm"
	"github.com/goxiaoy/uow"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(services *conf.Services, security *conf.Security, userConf *conf2.UserConf, confData *conf2.Data, logger log.Logger, passwordValidatorConfig *biz.PasswordValidatorConfig, config *uow.Config, gormConfig *gorm.Config, webMultiTenancyOption *http.WebMultiTenancyOption) (*kratos.App, func(), error) {
	tokenizerConfig := jwt.NewTokenizerConfig(security)
	tokenizer := jwt.NewTokenizer(tokenizerConfig)
	dbOpener, cleanup := gorm.NewDbOpener()
	manager := uow2.NewUowManager(gormConfig, config, dbOpener)
	saasContributor := api.NewSaasContributor(webMultiTenancyOption)
	userContributor := api.NewUserContributor()
	option := api.NewDefaultOption(saasContributor, userContributor)
	tenantStore := data.NewTenantStore()
	decodeRequestFunc := _wireDecodeRequestFuncValue
	encodeResponseFunc := _wireEncodeResponseFuncValue
	encodeErrorFunc := _wireEncodeErrorFuncValue
	dbProvider := data.NewProvider(confData, gormConfig, dbOpener, tenantStore, logger)
	dataData, cleanup2, err := data.NewData(confData, dbProvider, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData)
	passwordHasher := biz.NewPasswordHasher()
	userValidator := biz.NewUserValidator()
	passwordValidator := biz.NewPasswordValidator(passwordValidatorConfig)
	lookupNormalizer := biz.NewLookupNormalizer()
	userTokenRepo := data.NewUserTokenRepo(dataData)
	refreshTokenRepo := data.NewRefreshTokenRepo(dataData)
	userTenantRepo := data.NewUserTenantRepo(dataData)
	userManager := biz.NewUserManager(userRepo, passwordHasher, userValidator, passwordValidator, lookupNormalizer, userTokenRepo, refreshTokenRepo, userTenantRepo, logger)
	enforcerProvider := data.NewEnforcerProvider(dbProvider)
	permissionService := casbin.NewPermissionService(enforcerProvider)
	userRoleContributor := service.NewUserRoleContributor(userRepo)
	authorizationOption := service.NewAuthorizationOption(userRoleContributor)
	subjectResolverImpl := authorization.NewSubjectResolver(authorizationOption)
	defaultAuthorizationService := authorization.NewDefaultAuthorizationService(permissionService, subjectResolverImpl, logger)
	userService := service.NewUserService(userManager, defaultAuthorizationService)
	accountService := service.NewAccountService(userManager)
	roleRepo := data.NewRoleRepo(dataData)
	roleManager := biz.NewRoleManager(roleRepo, lookupNormalizer)
	authService := service.NewAuthService(userManager, roleManager, tokenizer, tokenizerConfig, passwordValidator, refreshTokenRepo, security)
	roleService := service.NewRoleServiceService(roleRepo, defaultAuthorizationService)
	servicePermissionService := service.NewPermissionService(defaultAuthorizationService, permissionService, subjectResolverImpl)
	signInManager := biz.NewSignInManager(userManager)
	auth := http2.NewAuth(decodeRequestFunc, encodeResponseFunc, userManager, logger, signInManager)
	defaultErrorHandler := server.NewDefaultErrorHandler(encodeErrorFunc)
	httpServer := server2.NewHTTPServer(services, security, tokenizer, manager, webMultiTenancyOption, option, tenantStore, decodeRequestFunc, encodeResponseFunc, encodeErrorFunc, logger, userService, accountService, authService, roleService, servicePermissionService, auth, defaultErrorHandler)
	grpcServer := server2.NewGRPCServer(services, tokenizer, tenantStore, manager, webMultiTenancyOption, option, logger, userService, accountService, authService, roleService, servicePermissionService)
	migrate := data.NewMigrate(dataData)
	roleSeed := biz.NewRoleSeed(roleManager, permissionService)
	userSeed := biz.NewUserSeed(userManager, roleManager)
	fake := seed.NewFake(userManager)
	permissionSeeder := biz.NewPermissionSeeder(permissionService, permissionService, roleManager)
	seeder := server2.NewSeeder(userConf, manager, migrate, roleSeed, userSeed, fake, permissionSeeder)
	app := newApp(logger, httpServer, grpcServer, seeder)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}

var (
	_wireDecodeRequestFuncValue  = server.ReqDecode
	_wireEncodeResponseFuncValue = server.ResEncoder
	_wireEncodeErrorFuncValue    = server.ErrEncoder
)
