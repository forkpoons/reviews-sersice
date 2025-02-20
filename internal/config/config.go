package config

import (
	"github.com/forkpoons/library/pg"
	"github.com/forkpoons/library/probes"
	"github.com/forkpoons/library/tracing"
	"github.com/forkpoons/library/yamlenv"
	"github.com/forkpoons/library/zerohook"
)

type ApiConfig struct {
	Port *yamlenv.Env[int] `yaml:"port"`
	test string
}

type AppConfig struct {
	Probe  probes.ProbeConfig `yaml:"probe"`
	Secret string             `yaml:"secret"`
}

type Config struct {
	Log      zerohook.LoggerConfig `yaml:"log"`
	App      AppConfig             `yaml:"app"`
	Api      ApiConfig             `yaml:"api"`
	Postgres pg.PostgresConfig     `yaml:"postgres"`
	Kafka    struct {
		Broker     *yamlenv.Env[string] `yaml:"broker"`
		Topic      *yamlenv.Env[string] `yaml:"topic"`
		GroupID    *yamlenv.Env[string] `yaml:"group_id"`
		MaxRetries *yamlenv.Env[int]    `yaml:"max_retries"`
	} `yaml:"kafka"`
	Jaeger tracing.JaegerConfig `yaml:"jaeger"`
}
