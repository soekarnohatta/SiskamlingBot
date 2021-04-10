package bot

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type TgConfig struct {
	WebhookURL    string `env:"WEBHOOK_URL"`
	WebhookPath   string `env:"WEBHOOK_PATH"`
	WebhookServe  string `env:"WEBHOOK_SERVE"`
	BotAPIKey     string `env:"BOT_API_KEY,required"`
	DatabaseURL   string `env:"DATABASE_URL,required"`
	RedisAddress  string `env:"REDIS_ADDRESS,required"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	BotVer        string `env:"BOT_VERSION,required"`
	WebhookPort   int    `env:"WEBHOOK_PORT"`
	OwnerID       int    `env:"OWNER_ID,required"`
	SudoUsers     []int  `env:"SUDO_USERS,required" envSeparator:":"`
	LogEvent      int64  `env:"LOG_EVENT,required"`
	LogBan        int64  `env:"LOG_BAN,required"`
	MainGrp       int64  `env:"MAIN_GRP,required"`
	IsDebug       bool   `env:"IS_DEBUG"`
	CleanPolling  bool   `env:"CLEAN_POLLING,required"`
}

var Config TgConfig

func NewConfig() TgConfig {
	returnConfig := TgConfig{}

	err := godotenv.Load("data/.env")
	if err != nil {
		panic(err.Error())
	}

	errParse := env.Parse(&returnConfig)
	if errParse != nil {
		panic(errParse)
	}

	Config = returnConfig
	return returnConfig
}

