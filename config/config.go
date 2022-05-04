package config

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const (
	EnvProduction          = "production"
	EnvStaging             = "staging"
	EnvDevelopment         = "development"
	EnvLocalhost           = "localhost"
	AdminUserId            = int64(-1)
	RedisDefaultExpireTime = time.Second * 60 * 60 * 24 * 30 // 預設一個月
)

var EnvShortName = map[string]string{
	EnvProduction:  "prod",
	EnvStaging:     "stag",
	EnvDevelopment: "dev",
	EnvLocalhost:   "local",
}

// Environment
func GetEnvironment() string {
	return os.Getenv("ENVIRONMENT")
}

func GetShortEnv() string {
	return EnvShortName[GetEnvironment()]
}

// Base path
var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)

func GetBasePath() string {
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func InitEnv() {
	remoteBranch := os.Getenv("REMOTE_BRANCH")
	if remoteBranch == "" {
		// load env from .env file
		path := GetBasePath() + "/.env"
		err := godotenv.Load(path)
		if err != nil {
			log.Panicln(err)
		}
	}
}
