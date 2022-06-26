package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	uow2 "github.com/go-saas/uow"
	"github.com/go-saas/saas"
	sapi "github.com/go-saas/kit/pkg/api"
	"github.com/go-saas/kit/pkg/authn/jwt"
	"github.com/go-saas/kit/pkg/authn/session"
	"github.com/go-saas/kit/pkg/conf"
	"github.com/go-saas/kit/pkg/localize"
	"github.com/go-saas/kit/pkg/logging"
	"github.com/go-saas/kit/pkg/server"
	"github.com/go-saas/kit/pkg/uow"
	uapi "github.com/go-saas/kit/user/api"
	"github.com/go-saas/kit/user/i18n"
	shttp "github.com/go-saas/saas/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Services,
	sCfg *conf.Security,
	tokenizer jwt.Tokenizer,
	uowMgr uow2.Manager,
	mOpt *shttp.WebMultiTenancyOption,
	apiOpt *sapi.Option,
	ts saas.TenantStore,
	reqDecoder khttp.DecodeRequestFunc,
	resEncoder khttp.EncodeResponseFunc,
	errEncoder khttp.EncodeErrorFunc,
	logger log.Logger,
	userTenant *uapi.UserTenantContrib,
	validator sapi.TrustedContextValidator,
	refreshProvider session.RefreshTokenProvider,
	register HttpServerRegister,
) *khttp.Server {
	var opts []khttp.ServerOption
	opts = server.PatchHttpOpts(logger, opts, uapi.ServiceName, c, sCfg, reqDecoder, resEncoder, errEncoder,
		session.Auth(sCfg, validator),
		session.Refresh(errEncoder, refreshProvider, validator),
	)
	middlewares := middleware.Chain(server.Recovery(),
		tracing.Server(),
		logging.Server(logger),
		metrics.Server(),
		validate.Validator(),
		//TODO combine i18n
		localize.I18N(i18n.Files...),
		jwt.ServerExtractAndAuth(tokenizer, logger),
		sapi.ServerPropagation(apiOpt, validator, logger),
		server.Saas(mOpt, ts, validator, func(o *saas.TenantResolveOption) {
			o.AppendContribs(userTenant)
		}),
		uow.Uow(logger, uowMgr))
	opts = append(opts, []khttp.ServerOption{
		khttp.Middleware(middlewares),
	}...)

	srv := khttp.NewServer(opts...)

	register.Register(srv, middlewares)
	return srv
}
