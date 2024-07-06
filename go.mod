module SiskamlingBot

// +heroku goVersion go1.17.1
go 1.22

toolchain go1.22.4

require (
	github.com/PaulSonOfLars/gotgbot/v2 v2.0.0-rc.28
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/joho/godotenv v1.5.1
	github.com/shirou/gopsutil v3.21.11+incompatible
	go.mongodb.org/mongo-driver v1.16.0
)

require (
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	github.com/tklauser/go-sysconf v0.3.14 // indirect
	github.com/tklauser/numcpus v0.8.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240424034433-3c2c7870ae76 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	golang.org/x/crypto v0.25.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
)

replace github.com/PaulSonOfLars/gotgbot/v2 => github.com/PaulSonOfLars/gotgbot/v2 v2.0.0-rc.7.0.20220508170057-17a08aaccb58
