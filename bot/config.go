package bot

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
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
	MainGroup     int64  `env:"MAIN_GRP,required"`
	IsDebug       bool   `env:"IS_DEBUG"`
	CleanPolling  bool   `env:"CLEAN_POLLING,required"`
}

var Config TgConfig

func NewConfig() TgConfig {
	returnConfig := TgConfig{}

	err := godotenv.Load("data/.env")
	if err != nil {
		returnConfig.WebhookURL = os.Getenv("WEBHOOK_URL")
		returnConfig.WebhookPath = os.Getenv("WEBHOOK_PATH")
		returnConfig.WebhookPort, _ = strconv.Atoi(os.Getenv("WEBHOOK_PORT"))
		returnConfig.WebhookServe = os.Getenv("WEBHOOK_SERVE")

		returnConfig.BotAPIKey = os.Getenv("BOT_API_KEY")
		_, cleanPolling := os.LookupEnv("CLEAN_POLLING")
		returnConfig.CleanPolling = cleanPolling
		_, isDebug := os.LookupEnv("IS_DEBUG")
		returnConfig.IsDebug = isDebug

		logEvent, _ := strconv.Atoi(os.Getenv("LOG_EVENT"))
		returnConfig.LogEvent = int64(logEvent)
		logBan, _ := strconv.Atoi(os.Getenv("LOG_BAN"))
		returnConfig.LogBan = int64(logBan)
		returnConfig.OwnerID, _ = strconv.Atoi(os.Getenv("OWNER_ID"))
		returnConfig.SudoUsers = strToIntSlice(strings.Split(os.Getenv("SUDO_USERS"), ","))

		returnConfig.DatabaseURL = os.Getenv("DATABASE_URL")
		returnConfig.RedisAddress = os.Getenv("REDIS_ADDRESS")
		returnConfig.RedisPassword = os.Getenv("REDIS_PASSWORD")

		return returnConfig
	}

	err = env.Parse(&returnConfig)
	if err != nil {
		panic(err.Error())
	}

	Config = returnConfig
	return returnConfig
}

func strToIntSlice(s []string) []int {
	var newIntSlice []int
	for _, val := range s {
		newInt, _ := strconv.Atoi(val)
		newIntSlice = append(newIntSlice, newInt)
	}
	return newIntSlice
}
