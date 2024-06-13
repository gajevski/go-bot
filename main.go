package main

import (
  "go-bot/bot"
)

const (
  TOKEN_ID = "DISCORD_ID"
)

func main() {
  bot.BotToken = TOKEN_ID
  bot.Start()
}
