package main

import (
	"context"
	"flag"
	"github.com/forkpoons/reviews-sersice/internal/config"
	"github.com/forkpoons/reviews-sersice/internal/repository"
	apiService "github.com/forkpoons/reviews-sersice/internal/service/api"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	_ "regexp"
	"syscall"

	_ "github.com/prometheus/client_golang/prometheus"
	_ "github.com/prometheus/client_golang/prometheus/collectors"

	"github.com/forkpoons/library/pg"
	"github.com/forkpoons/library/probes"
	"github.com/forkpoons/library/yamlreader"
	"github.com/forkpoons/library/zerohook"
)

func main() {
	ctx := context.Background()
	cfg := MustNewConfig(parseFlags(), zerohook.Logger)

	pgConn, err := pg.NewPG(ctx, cfg.Postgres.Conn.Value, zerohook.Logger)
	if err != nil {
		zerohook.Logger.Fatal().Msgf("Error initializing PostgreSQL connection: %v", err)
	}
	promoRepo := repository.NewPromo(ctx, pgConn.Pool(), zerohook.Logger)
	probe := probes.NewProbe(ctx, &cfg.App.Probe)
	go probe.Start()
	if err != nil {
		zerohook.Logger.Fatal().Msgf("Error starting probe: %v", err)
	}
	zerohook.Logger.Debug().Msgf("Starting server")
	api := apiService.NewService(zerohook.Logger, promoRepo, cfg.App.Secret)
	go api.Start(ctx)
	if err != nil {
		zerohook.Logger.Fatal().Msgf("Error starting API: %v", err)
	}
	zerohook.Logger.Debug().Msgf("Starting server2")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func MustNewConfig(path string, lgr zerolog.Logger) *config.Config {
	cfg, err := yamlreader.NewConfig[config.Config](path)

	if err != nil {
		lgr.Fatal().Str("path", path).Err(err).Msg("ошибка чтения конфигурации приложения")
		return nil
	}

	return cfg
}

func parseFlags() string {
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()
	return configPath
}
