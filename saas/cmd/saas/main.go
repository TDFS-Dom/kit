package main

import (
	"context"
	"flag"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-saas/kit/pkg/event"
	"github.com/go-saas/kit/pkg/job"
	"github.com/go-saas/kit/pkg/logging"
	"github.com/go-saas/kit/pkg/tracers"
	"github.com/go-saas/saas/seed"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-saas/kit/saas/private/conf"

	_ "github.com/go-saas/kit/pkg/event/kafka"
	_ "github.com/go-saas/kit/pkg/event/pulsar"
)

// go build -buildvcs=false -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "SAAS"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
	ifSeed   bool
	id, _    = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "./configs", "config path, eg: -conf config.yaml")
	flag.BoolVar(&ifSeed, "seed", true, "run seeder or not")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, js *job.Server, es *event.ConsumerFactoryServer, seeder seed.Seeder) *kratos.App {
	if ifSeed {
		if err := seeder.Seed(context.Background(), seed.AddHost()); err != nil {
			panic(err)
		}
	}
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
			js,
			es,
		),
	)
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	l, lc, err := logging.NewLogger(bc.Logging)
	if err != nil {
		panic(err)
	}
	defer lc()
	logger := log.With(l,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
	shutdown, err := tracers.SetTracerProvider(context.Background(), bc.Tracing, Name)
	if err != nil {
		log.Error(err)
	}
	defer shutdown()
	app, cleanup, err := initApp(bc.Services, bc.Security, bc.Data, bc.Saas, logger, bc.App)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
