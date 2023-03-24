package server

import (
	v1 "helloworld/api/helloworld/v1"
	student_v1 "helloworld/api/student/v1"
	"helloworld/internal/conf"
	"helloworld/internal/service"
	"helloworld/pkg/encoder"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, student *service.StudentService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		// 返回code
		http.ResponseEncoder(encoder.RespEncoder),

		// 请求的中间健
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	student_v1.RegisterStudentHTTPServer(srv, student)
	return srv
}
