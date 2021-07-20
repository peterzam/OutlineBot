package controller

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var CHANNEL_ID = os.Getenv("CHANNEL_ID")

func StartBot() {
	dg, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Println(":::Error session <discordgo.New>:::\n", err)
		return
	}
	dg.AddHandler(messageCreate)
	err = dg.Open()
	if err != nil {
		log.Println(":::Error connection <dg.Open>:::\n", err)
		return
	}
	fmt.Println("Bot is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.ChannelID == CHANNEL_ID && (m.Content == "!req" || m.Content == "!request" || m.Content == "!vpn") {
		u, err := s.UserChannelCreate(m.Author.ID)
		if err != nil {
			log.Println(":::User Channel create error <s.UserChannelCreate>:::\n", err)
			s.ChannelMessageSend(m.ChannelID, "Sorry, Message Error. Contact developer to fix")
			return
		}
		name := GetName(m)

		key, status := CreateNewUser(name)

		if status != true {
			s.ChannelMessageSend(m.ChannelID, key+"\nSorry, Please contact developer")
		} else {
			s.ChannelMessageSend(u.ID, key)
			s.ChannelMessageSendReply(m.ChannelID, name+"\nVPN key sent.\nPlease check direct messages.", m.Reference())
		}
	}
	if m.Content == "!about" {
		s.ChannelMessageSend(m.ChannelID, "https://github.com/peterzam/OutlineBot")
	}
}

func GetName(m *discordgo.MessageCreate) string {

	for i := 0; i < len(m.Member.Roles); i++ {
		switch m.Member.Roles[i] {
		case os.Getenv("ROLE_CALL"):
			return "call-" + m.Author.Username + "#" + m.Author.Discriminator

		case os.Getenv("ROLE_COM"):
			return "com-" + m.Author.Username + "#" + m.Author.Discriminator

		case os.Getenv("ROLE_DATA"):
			return "data-" + m.Author.Username + "#" + m.Author.Discriminator

		case os.Getenv("ROLE_FB"):
			return "fb-" + m.Author.Username + "#" + m.Author.Discriminator

		case os.Getenv("ROLE_WEB"):
			return "web-" + m.Author.Username + "#" + m.Author.Discriminator

		default:
			return "other-" + m.Author.Username + "#" + m.Author.Discriminator
		}
	}
	return "unassigned-" + m.Author.Username + "#" + m.Author.Discriminator
}
