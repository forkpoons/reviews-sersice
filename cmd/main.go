package main

import (
	"context"
	"flag"
	"github.com/ItemCloudShopping/library/pg"
	"github.com/ItemCloudShopping/promo/internal/config"
	"github.com/ItemCloudShopping/promo/internal/repository"
	apiService "github.com/ItemCloudShopping/promo/internal/service/api"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"

	"github.com/ItemCloudShopping/library/probes"
	"github.com/ItemCloudShopping/library/tracing"
	"github.com/ItemCloudShopping/library/yamlreader"
	"github.com/ItemCloudShopping/library/zerohook"
)

func main() {
	ctx := context.Background()
	cfg := MustNewConfig(parseFlags(), zerohook.Logger)

	pgConn, err := pg.NewPG(ctx, cfg.Postgres.Conn.Value, zerohook.Logger)
	if err != nil {
		zerohook.Logger.Fatal().Msgf("Error initializing PostgreSQL connection: %v", err)
	}
	promoRepo := repository.NewPromo(ctx, pgConn.Pool(), zerohook.Logger)
	probes := probes.NewProbe(ctx, &cfg.App.Probe)
	go probes.Start()
	api := apiService.NewService(zerohook.Logger, promoRepo, cfg.App.Secret)

	go api.Start(ctx)
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
