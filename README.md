[![codecov](https://codecov.io/gh/soekarnohatta/SiskamlingBot/branch/main/graph/badge.svg?token=M4U97ZU3N2)](https://codecov.io/gh/soekarnohatta/SiskamlingBot)

# SiskamlingBot Go

Official repository of SiskamlingBot, written in Golang

# Preparation

1. Golang 1.15+
2. MongoDB
3. Nginx or OpenLiteSpeed for reverse proxy (Optional)
4. Bot Token

# Run Development

- Clone this repo.
- Install MongoDB and create database e.g. `test`.
- Copy .example.env to .env and fill all required fields.
- Type `go run .` in your CLI.
- Your bot has succesfully working local as Development using Poll mode.

# Run Production

- Clone this repo.
- Install MongoDB and create database e.g. `test`.
- Copy .example.env to .env and fill all required fields.
- Server with domain name must include HTTPS support (e.g https://yoursite.co.id) for using webhook mode.
- Run `go build .` and place your binary somewhere.
- Setup reverse proxy for Web
  Server, [here example](https://www.google.com/search?client=firefox-b-d&q=nginx+reverse+proxy+example).
- Launch bot with `./SiskamlingBot`, your bot will run using poll or webhook.