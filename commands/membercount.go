package commands

import (
	"fmt"

	"github.com/CubicCommunity/BloqGo/assets"
	"github.com/CubicCommunity/BloqGo/include"
	"github.com/CubicCommunity/BloqGo/log"
	"github.com/CubicCommunity/BloqGo/registry"

	"github.com/bwmarrin/discordgo"
)

var MemberCount *include.Command = &include.Command{
	Data: &discordgo.ApplicationCommand{
		Name:        "member-count",
		Description: "View the member count of this server.",
		IntegrationTypes: &[]discordgo.ApplicationIntegrationType{
			discordgo.ApplicationIntegrationGuildInstall,
		},
		Contexts: &[]discordgo.InteractionContextType{
			discordgo.InteractionContextGuild,
		},
		NSFW: include.NSFW(false),
	},
	Handler: func(s *discordgo.Session, c *discordgo.ApplicationCommandInteractionData, i *discordgo.Interaction) error {
		guild, err := s.GuildWithCounts(i.GuildID)

		if err != nil {
			return err
		} else {
			total := guild.MemberCount
			bots := 0

			log.Debug("Getting member and bot counts for guild of ID %s", guild.ID)

			for _, Member := range guild.Members {
				if Member.User.Bot {
					bots++
				}
			}

			respondEmbed := &discordgo.MessageEmbed{
				Author: &discordgo.MessageEmbedAuthor{
					Name:    guild.Name,
					IconURL: guild.IconURL("128"),
				},
				Color: assets.Colors.Primary,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Member Count",
						Value:  fmt.Sprintf("**```%v```**", total),
						Inline: false,
					},
				},
				Footer: &discordgo.MessageEmbedFooter{
					Text: fmt.Sprintf("%v Bots", bots),
				},
			}

			log.Info("Guild of ID %s has %v members (%v bots)", guild.ID, total, bots)

			return s.InteractionRespond(i, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						respondEmbed,
					},
				},
			})
		}
	},
}

func init() {
	registry.Register(&include.Command{
		Data:    MemberCount.Data,
		Handler: MemberCount.Handler,
	})
}
