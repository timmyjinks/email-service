package main

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Resend Resend
}

type Resend struct {
	KEY string `env:"RESEND_API_KEY,required"`
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	if cfg.Resend.KEY == "" {
		log.Fatal("[ERROR] Resend API key not assigned")
	}

	log.Println(cfg)

	return cfg
}
