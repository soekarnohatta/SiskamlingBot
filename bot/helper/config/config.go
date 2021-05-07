package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

type TgConfig struct {
	WebhookURL    string `env:"WEBHOOK_URL"`
	WebhookPath   string `env:"WEBHOOK_PATH"`
	WebhookListen string `env:"WEBHOOK_LISTEN"`
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

var Config *TgConfig

func NewConfig() *TgConfig {
	returnConfig := new(TgConfig)

	err := godotenv.Load("data/.env")
	if err != nil {
		log.Println("using declared Env vars!")

		returnConfig.WebhookURL = os.Getenv("WEBHOOK_URL")
		returnConfig.WebhookPath = os.Getenv("WEBHOOK_PATH")

		port, _ := strconv.Atoi(os.Getenv("PORT"))
		if port != 0 {
			returnConfig.WebhookPort = port
		} else {
			returnConfig.WebhookPort, _ = strconv.Atoi(os.Getenv("WEBHOOK_PORT"))
		}

		returnConfig.WebhookListen = os.Getenv("WEBHOOK_LISTEN")
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
		returnConfig.SudoUsers = strToIntSlice(strings.Split(os.Getenv("SUDO_USERS"), ":"))

		returnConfig.DatabaseURL = os.Getenv("DATABASE_URL")
		returnConfig.RedisAddress = os.Getenv("REDIS_ADDRESS")
		returnConfig.RedisPassword = os.Getenv("REDIS_PASSWORD")

		Config = returnConfig
		return returnConfig
	}

	err = env.Parse(returnConfig)
	if err != nil {
		panic(err.Error())
	}
	log.Println("configurations have been parsed succesfully!")

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
