package commands

import (
	"BloqGo/assets"
	"BloqGo/include"
	"BloqGo/registry"

	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var Ping *include.Command = &include.Command{
	Data: &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Ping the bot, test its latency.",
		IntegrationTypes: &[]discordgo.ApplicationIntegrationType{
			discordgo.ApplicationIntegrationUserInstall,
			discordgo.ApplicationIntegrationGuildInstall,
		},
		Contexts: &[]discordgo.InteractionContextType{
			discordgo.InteractionContextGuild,
			discordgo.InteractionContextPrivateChannel,
			discordgo.InteractionContextBotDM,
		},
		NSFW: include.NSFW(false),
	},
	Handler: func(s *discordgo.Session, c *discordgo.ApplicationCommandInteractionData, i *discordgo.Interaction) error {
		created, err := discordgo.SnowflakeTimestamp(i.ID)

		if err == nil {
			latency := time.Since(created).Milliseconds()

			respondEmbed := &discordgo.MessageEmbed{
				Title: fmt.Sprintf("%s Ping", assets.Icons.Info),
				Color: assets.Colors.Primary,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Latency",
						Value:  fmt.Sprintf("%vms", latency),
						Inline: false,
					},
					{
						Name:   "API Latency",
						Value:  fmt.Sprintf("%vms", s.HeartbeatLatency().Milliseconds()),
						Inline: false,
					},
				},
			}

			return s.InteractionRespond(i, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						respondEmbed,
					},
					Flags: discordgo.MessageFlagsEphemeral,
				},
			})
		} else {
			return err
		}
	},
}

func init() {
	registry.Register(&include.Command{
		Data:    Ping.Data,
		Handler: Ping.Handler,
	})
}
