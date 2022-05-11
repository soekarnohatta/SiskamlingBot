module SiskamlingBot

// +heroku goVersion go1.17.1
go 1.17

require (
	github.com/PaulSonOfLars/gotgbot/v2 v2.0.0-rc.7.0.20220502141358-79f30975edeb
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/joho/godotenv v1.3.0
	github.com/shirou/gopsutil v3.21.11+incompatible
	go.mongodb.org/mongo-driver v1.5.3
)

require (
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
)

replace github.com/PaulSonOfLars/gotgbot/v2 => github.com/PaulSonOfLars/gotgbot/v2 v2.0.0-rc.7.0.20220508170057-17a08aaccb58
