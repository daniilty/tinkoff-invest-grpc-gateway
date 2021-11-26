package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	invest "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/daniilty/tinkoff-invest-grpc-gateway/internal/server"
	schema "github.com/daniilty/tinkoff-invest-grpc-schema"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func run() error {
	cfg, err := loadEnvConfig()
	if err != nil {
		return err
	}

	client := invest.NewRestClient(cfg.tinkoffToken)

	loggerCfg := zap.NewProductionConfig()

	logger, err := loggerCfg.Build()
	if err != nil {
		return err
	}

	sugaredLogger := logger.Sugar()

	grpcControllers := server.NewGRPC(client, server.WithZapLogger(sugaredLogger))
	grpcServer := grpc.NewServer()
	schema.RegisterTinkoffInvestAPIGatewayServer(grpcServer, grpcControllers)

	listener, err := net.Listen("tcp", cfg.grpcAddr)
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	sugaredLogger.Infow("GRPC server is starting.", "addr", listener.Addr())
	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			sugaredLogger.Errorw("Server failed to start.", "err", err)
		}
	}()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-termChan

	sugaredLogger.Info("Gracefully stopping GRPC server.")
	grpcServer.GracefulStop()

	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
