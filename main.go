package main

import (
  "os"
  "log"

  "go-bot/bot"
  "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
      log.Fatalf("Error loading .env file")
  }
  bot.BotToken = os.Getenv("BOT_TOKEN")
  bot.Start()
}
