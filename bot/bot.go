package bot

import (
  "fmt"
  "time"
  "log"
  "os"
  "os/signal"
  "strings"
  "github.com/bwmarrin/discordgo"
)

var BotToken string

func Start()  {
 discord, err := discordgo.New("Bot " + BotToken)
 if err != nil {
   log.Fatal("Error")
 }
 discord.AddHandler(newMessage)
 discord.Open()
 defer discord.Close()
 fmt.Println("Bot running....")
 c := make(chan os.Signal, 1)
 signal.Notify(c, os.Interrupt)
 <-c
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate)  {
  // Prevent bot from respoding to its own messages
  if message.Author.ID == discord.State.User.ID {
    return
  }
 switch {
 case strings.Contains(message.Content, "!help"):
  discord.ChannelMessageSend(message.ChannelID, "Hello World😃")
 case strings.Contains(message.Content, "!bye"):
  discord.ChannelMessageSend(message.ChannelID, "Good Bye👋")
 case strings.Contains(message.Content, "!joined"):
  joined := message.Member.JoinedAt.Format("2006-01-02 15:04:05")
  now := time.Now()
  days := now.Sub(message.Member.JoinedAt).Hours() / 24
  response := fmt.Sprintf("Joined at: %s (%d days ago)", joined, int(days))
  discord.ChannelMessageSend(message.ChannelID, response)
 case strings.Contains(message.Content, "!clear-all"):
   messagesToDelete := []string{"1251536094705614910"}
   discord.ChannelMessageSend(message.ChannelID, "clear-all")
   discord.ChannelMessagesBulkDelete(message.ChannelID, messagesToDelete)
}
}
