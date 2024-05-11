package config

import (
	"flag"
	"fmt"
	"sync"

	goconfig "github.com/Yalantis/go-config"
)

var (
	config        *Config
	isInitialised bool
	mutex         = new(sync.Mutex)
)

type Config struct {
	HTTP     HTTP     `json:"http"`
	GRPC     GRPC     `json:"grpc"`
	JWT      JWT      `json:"jwt"`
	Postgres Postgres `json:"postgres"`
	Redis    Redis    `json:"redis"`
	Log      Logger   `json:"logger"`
	AWS      AWS      `json:"aws"`
}

type HTTP struct {
	Address string `json:"address" envconfig:"HTTP_ADDRESS" default:":8080"`
}

type GRPC struct {
	Address string `json:"address" envconfig:"GRPC_ADDRESS" default:":8081"`
}

type JWT struct {
	Secret string `json:"secret" envconfig:"JWT_SECRET" default:"my-totally-secret-key"`
}

type Postgres struct {
	Host     string `json:"host"     envconfig:"POSTGRES_HOST"     default:"localhost"`
	Port     int    `json:"port"     envconfig:"POSTGRES_PORT"     default:"5432"`
	Database string `json:"database" envconfig:"POSTGRES_DATABASE" default:"sportgroup_api"`
	User     string `json:"user"     envconfig:"POSTGRES_USER"     default:"sportgroup_api_user"`
	Password string `json:"password" envconfig:"POSTGRES_PASSWORD"`
	Log      bool   `json:"log"      envconfig:"POSTGRES_LOG"      default:"true"`
}

type Redis struct {
	Address  string `json:"address"  envconfig:"REDIS_ADDRESS"`
	Password string `json:"password" envconfig:"REDIS_PASSWORD"`
}

type Logger struct {
	Level string `json:"level" envconfig:"LOGGER_LEVEL" default:"info"`
}

type AWS struct {
	Region          string `json:"region"            envconfig:"AWS_REGION" default:"us-east-1"`
	AccessKeyID     string `json:"access_key_id"     envconfig:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `json:"access_key_secret" envconfig:"AWS_SECRET_ACCESS_KEY"`
	S3              S3     `json:"s3"`
}

type S3 struct {
	Bucket string `json:"bucket" envconfig:"AWS_S3_BUCKET"`
}

func New() (*Config, error) {
	var cfg Config

	flag.Parse()

	if err := goconfig.Init(&cfg, flag.Arg(0)); err != nil {
		return nil, fmt.Errorf("init config: %w", err)
	}

	config = &cfg
	isInitialised = true

	return &cfg, nil
}

func Get() *Config {
	mutex.Lock()
	if !isInitialised {
		cfg, err := New()
		if err != nil {
			panic(err)
		}
		config = cfg
	}
	mutex.Unlock()

	return config
}
