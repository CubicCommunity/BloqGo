package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/CubicCommunity/BloqGo/assets"
	_ "github.com/CubicCommunity/BloqGo/commands"
	"github.com/CubicCommunity/BloqGo/log"
	"github.com/CubicCommunity/BloqGo/registry"

	"github.com/bwmarrin/discordgo"
)

func main() {
	log.Print("Starting BloqGo...")

	log.Debug("Getting token...")
	token := os.Getenv("MAIN_TOKEN")
	if token == "" {
		log.Error("MAIN_TOKEN not set")
		return
	}

	log.Debug("Creating client...")
	dg, err := discordgo.New("Bot " + strings.TrimSpace(token))
	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Debug("Adding intents...")
	dg.Identify.Intents = discordgo.IntentsAll

	log.Debug("Connecting client...")
	dg.AddHandlerOnce(Ready)

	err = dg.Open()
	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info("Starting handlers...")

	dg.AddHandler(MessageCreate)
	dg.AddHandler(InteractionCreate)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	dg.Close()
}

func Ready(s *discordgo.Session, m *discordgo.Ready) {
	log.Print("Bot is ready. Registering commands...")

	for _, cmd := range registry.Commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd.Data)

		if err == nil {
			log.Debug("Registered command: %s", cmd.Data.Name)
		} else {
			log.Error("Failed to register command %s: %v", cmd.Data.Name, err)
		}
	}

	whId := os.Getenv("MAIN_LOG_WH_ID")
	whToken := os.Getenv("MAIN_LOG_WH_TOKEN")
	if whId == "" || whToken == "" {
		log.Error("Developer logging webhook credentials not found")
	} else {
		_, err := s.WebhookExecute(whId, whToken, true, &discordgo.WebhookParams{
			AvatarURL: s.State.User.AvatarURL("512"),
			Embeds: []*discordgo.MessageEmbed{
				{
					Author: &discordgo.MessageEmbedAuthor{
						Name: "Service Status",
					},
					Description: fmt.Sprintf("%s **%s** is now __online__", assets.Icons.Check, s.State.User.Username),
					Color:       assets.Colors.Primary,
					Footer: &discordgo.MessageEmbedFooter{
						Text:    s.State.User.Username,
						IconURL: s.State.User.AvatarURL("512"),
					},
				},
			},
		})

		if err != nil {
			log.Error(err.Error())
		} else {
			log.Info("Sent developer log webhook for startup")
		}
	}

	log.Done("BloqGo is online!")
	log.Print("change this part") // ptero egg lol
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		log.Debug("Detected bot message in channel in guild of ID %s", m.GuildID)
	} else {
		log.Debug("Received message in channel in guild of ID %s", m.GuildID)
	}
}

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		data := i.ApplicationCommandData()
		log.Debug("Command /%s used", data.Name)

		for _, cmd := range registry.Commands {
			if cmd.Data.Name == data.Name {
				err := cmd.Handler(s, &data, i.Interaction)

				if err != nil {
					log.Error(err.Error())

					respondEmbed := &discordgo.MessageEmbed{
						Description: fmt.Sprintf("%s There was an error while executing this command.", assets.Icons.XMark),
						Color:       assets.Colors.Secondary,
					}

					e := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Embeds: []*discordgo.MessageEmbed{
								respondEmbed,
							},
							Flags: discordgo.MessageFlagsEphemeral,
						},
					})

					if e != nil {
						log.Error(e.Error())
					}
				}

				break
			}
		}
	} else {
		log.Error("Command '%s' did not match any accepted integration types")
	}
}
