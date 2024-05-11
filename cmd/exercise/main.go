package main

import (
	"exercise/internal/config"
	handlergrpc "exercise/internal/handler/grpc"
	"exercise/internal/repo"
	"exercise/internal/usecase"
	"exercise/pkg/database"
	"exercise/pkg/grpc"
	"exercise/pkg/logger"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := logger.New()
	defer func() {
		if recErr := recover(); recErr != nil {
			switch v := recErr.(type) {
			case error:
				log.Error().Stack().Err(v).Msg("fatal error")
			case string:
				log.Error().Stack().Err(fmt.Errorf("%s", v)).Msg("fatal error")
			default:
				log.Error().Stack().Err(fmt.Errorf("unsupported recover type: %v", v))
			}
			os.Exit(1)
		}
	}()

	cfg, err := config.Setup()
	if err != nil {
		log.Panic().Stack().Err(err).Msg("failed to read config")
	}

	db, err := database.NewPool(cfg.PG)
	if err != nil {
		log.Panic().Stack().Err(err).Msg("failed to init sql")
	}
	defer db.Close()

	grpcSrv, err := grpc.NewServer(cfg.GRPCServer)
	if err != nil {
		log.Panic().Stack().Err(err).Msg("failed to create grpc server")
	}

	sqlRepo := repo.NewRepo(db)
	uc := usecase.NewUC(sqlRepo)

	handler := handlergrpc.NewHandler(uc)
	handler.Register(grpcSrv)

	grpcErrCh := grpcSrv.Start()
	defer grpcSrv.Stop()
	log.Info().Stack().Msgf("grpc server started port: %d", cfg.GRPCServer.Port)

	log.Info().Stack().Msg("service User started")

	// on stop section
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	// wait for stop signal
	select {
	case <-signals:
		log.Info().Msgf("shutting down by signal")
	case err = <-grpcErrCh:
		log.Panic().Stack().Err(err).Msg("grpc server")
	}
}
