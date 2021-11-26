package main

import (
	"fmt"
	"os"
)

type envConfig struct {
	grpcAddr     string
	tinkoffToken string
}

func loadEnvConfig() (*envConfig, error) {
	const (
		provideEnvErrorMsg = `please provide "%s" environment variable`

		grpcAddrEnv     = "GRPC_SERVER_ADDR"
		tinkoffTokenEnv = "TINKOFF_INVEST_TOKEN"
	)

	var ok bool

	cfg := &envConfig{}

	cfg.grpcAddr, ok = os.LookupEnv(grpcAddrEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, grpcAddrEnv)
	}

	cfg.tinkoffToken, ok = os.LookupEnv(tinkoffTokenEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, tinkoffTokenEnv)
	}

	return cfg, nil
}
