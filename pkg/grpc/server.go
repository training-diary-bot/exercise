package grpc

import (
	"fmt"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	cfg        ServerConfig
	grpcServer *grpc.Server
}

func NewServer(cfg ServerConfig, addOpts ...grpc.ServerOption) (*Server, error) {
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			otelgrpc.UnaryServerInterceptor(),
			validator.UnaryServerInterceptor(),
			TimeoutInterceptor(cfg.Timeout),
		),
		grpc.ChainStreamInterceptor(
			grpc_recovery.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			otelgrpc.StreamServerInterceptor(),
			validator.StreamServerInterceptor(),
		),
	}
	opts = append(opts, addOpts...)

	s := &Server{
		cfg:        cfg,
		grpcServer: grpc.NewServer(opts...),
	}
	reflection.Register(s.grpcServer)

	return s, nil
}

func (s *Server) GRPCServer() *grpc.Server {
	return s.grpcServer
}

func (s *Server) Start() <-chan error {
	errCh := make(chan error)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Port))
	if err != nil {
		go func() {
			errCh <- errors.WithStack(err)
			close(errCh)
		}()
		return errCh
	}

	go func() {
		errCh <- errors.WithStack(s.grpcServer.Serve(lis))
		close(errCh)
	}()
	return errCh
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}
