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

const configFilename = "config.json"

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
	Address string `json:"address" default:":8080"`
}

type GRPC struct {
	ApiAddress string `json:"api_address" default:":8180"`
}

type JWT struct {
	Secret string `json:"secret" default:"my-totally-secret-key"`
}

type Postgres struct {
	Host     string `json:"host" default:"localhost"`
	Port     int    `json:"port" default:"5432"`
	Database string `json:"database" default:"oss"`
	User     string `json:"user" default:"oss"`
	Password string `json:"password"`
	Log      bool   `json:"log" default:"true"`
}

type Redis struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}

type Logger struct {
	Level string `json:"level" default:"info"`
}

type AWS struct {
	Region          string `json:"region" default:"us-east-1"`
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"access_key_secret"`
	S3              S3     `json:"s3"`
}

type S3 struct {
	Bucket string `json:"bucket"`
}

func New() (*Config, error) {
	var cfg Config

	filename := configFilename

	flag.Parse()

	if flag.Arg(0) != "" {
		filename = flag.Arg(0)
	}

	if err := goconfig.Init(&cfg, filename); err != nil {
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
