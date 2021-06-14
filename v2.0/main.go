package main

import (
	"context"
	"flag"
	"net"
	"os"

	"myDiscover2/endpoints"
	"myDiscover2/logging"
	"myDiscover2/services"
	"myDiscover2/transports"

	pb "myDiscover2/pd"

	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	ctx := context.Background()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	var svc services.Service
	svc = services.StringService{}

	// add logging middleware
	svc = logging.LoggingMiddleware(logger)(svc)
	endpoint := endpoints.MakeStringEndpoint(svc)

	//创建健康检查的Endpoint
	healthEndpoint := endpoints.MakeHealthCheckEndpoint(svc)
	//把算术运算Endpoint和健康检查Endpoint封装至StringEndpoints
	endpts := endpoints.StringEndpoints{
		StringEndpoint:      endpoint,
		HealthCheckEndpoint: healthEndpoint,
	}

	handler := transports.NewStringServer(ctx, endpts)
	ls, _ := net.Listen("tcp", "127.0.0.1:8080")
	gRPCServer := grpc.NewServer()
	pb.RegisterStringServiceServer(gRPCServer, handler)
	gRPCServer.Serve(ls)
}
