package commands

import (
	"fmt"
	"runtime"

	"github.com/CubicCommunity/BloqGo/assets"
	"github.com/CubicCommunity/BloqGo/include"
	"github.com/CubicCommunity/BloqGo/log"
	"github.com/CubicCommunity/BloqGo/registry"

	"github.com/bwmarrin/discordgo"
)

var About *include.Command = &include.Command{
	Data: &discordgo.ApplicationCommand{
		Name:        "about",
		Description: "View detailed information about the current installation of BloqGo.",
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
		ver, err := include.Version()

		if err != nil {
			return err
		} else {
			log.Info("Sending about message in guild of ID %s", i.GuildID)

			respondEmbed := &discordgo.MessageEmbed{
				Author: &discordgo.MessageEmbedAuthor{
					Name:    s.State.User.Username,
					IconURL: s.State.User.AvatarURL("512"),
				},
				Title:       fmt.Sprintf("BloqGo `v%s`", ver),
				Description: fmt.Sprintf("Running under Discord bot client **`%s`**`#%s` (`%s`) on shard **#%v**", s.State.User.Username, s.State.User.Discriminator, s.State.User.ID, s.ShardID),
				Color:       assets.Colors.Primary,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Go Version",
						Value:  fmt.Sprintf("`%s`", runtime.Version()),
						Inline: false,
					},
				},
				Footer: &discordgo.MessageEmbedFooter{
					Text: "This command is in the works - expect more information added soon.",
				},
			}

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
	registry.Register(About)
}
