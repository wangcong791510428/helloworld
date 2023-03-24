package main

import (
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2/registry"
	"net/url"
	"os"

	"helloworld/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	fmt.Println("flagconf=", flagconf)
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server, registry registry.Registrar) *kratos.App {
	if Name == "" {
		Name = "用户中心"
	}
	if Version == "" {
		Version = "v1.0"
	}
	// http://127.0.0.1:8000?isSecure=false
	// grpc://127.0.0.1:9000?isSecure=false
	urlHttp, _:= url.Parse("http://127.0.0.1:8000?isSecure=true")
	urlGrpc, _:= url.Parse("grpc://127.0.0.1:9000?isSecure=true")
	return kratos.New(
		kratos.ID(id),
		kratos.Endpoint(urlHttp,urlGrpc),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{
			"key": "name",
			"value": "value",
		}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(registry),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(any(err))
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(any(err))
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(any(err))
	}
	defer cleanup()

	// start and wait for stop signal
	if err = app.Run(); err != nil {
		panic(any(err))
	}
}
