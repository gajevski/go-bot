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
    helpMessage := `Bot Commands

**!help**
- **Description**: Lists all available commands and their functionalities.
- **Usage**: !help

**!joined**
- **Description**: Displays the date when the user joined the server and the number of days since they joined.
- **Usage**: !joined
- **Example Output**: Joined at: 2023-04-25 18:19:10 (100 days ago)

**!clear-all**
- **Description**: Sends a message indicating that all messages in the current channel will be deleted and then deletes all of messages.
- **Usage**: !clear-all
- **Example Output**: Deleting all messages on this channel...`
    discord.ChannelMessageSend(message.ChannelID, helpMessage)
 case strings.Contains(message.Content, "!joined"):
  joined := message.Member.JoinedAt.Format("2006-01-02 15:04:05")
  now := time.Now()
  days := now.Sub(message.Member.JoinedAt).Hours() / 24
  response := fmt.Sprintf("Joined at: %s (%d days ago)", joined, int(days))
  discord.ChannelMessageSend(message.ChannelID, response)
 case strings.Contains(message.Content, "!dm"):
  discord.ChannelMessageSend(message.ChannelID, "I have your bb in my DMs")
 case strings.Contains(message.Content, "!clear-all"):
  discord.ChannelMessageSend(message.ChannelID, "Deleting all messages on this channel...")

  messages, err := discord.ChannelMessages(message.ChannelID, 100, "", "", "")
  if err != nil {
    log.Printf("Error getting messages: %v", err)
    discord.ChannelMessageSend(message.ChannelID, "Error deleting all messages")
    return
  }

  allMessagesIds := []string{}
    
  for _, message := range messages {
    allMessagesIds = append(allMessagesIds, message.ID)
  }

  discord.ChannelMessagesBulkDelete(message.ChannelID, allMessagesIds)
}
}
