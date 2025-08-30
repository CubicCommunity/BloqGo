package commands

import (
	"fmt"

	"github.com/CubicCommunity/BloqGo/assets"
	"github.com/CubicCommunity/BloqGo/include"
	"github.com/CubicCommunity/BloqGo/registry"

	"github.com/bwmarrin/discordgo"
)

var Avatar *include.Command = &include.Command{
	Data: &discordgo.ApplicationCommand{
		Name:        "avatar",
		Description: "View a user's profile picture.",
		IntegrationTypes: &[]discordgo.ApplicationIntegrationType{
			discordgo.ApplicationIntegrationUserInstall,
			discordgo.ApplicationIntegrationGuildInstall,
		},
		Contexts: &[]discordgo.InteractionContextType{
			discordgo.InteractionContextGuild,
		},
		NSFW: include.NSFW(false),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "user",
				Description: "The user whose profile picture to view.",
				Type:        discordgo.ApplicationCommandOptionUser,
				Required:    false,
			},
		},
	},
	Handler: func(s *discordgo.Session, c *discordgo.ApplicationCommandInteractionData, i *discordgo.Interaction) error {
		user := c.GetOption("user")

		var username string

		var globalAv string
		var serverAv string

		if user != nil {
			u := user.UserValue(s)

			if u != nil {
				m, err := s.GuildMember(i.GuildID, u.ID)

				globalAv = u.AvatarURL("1024")
				if err != nil {
					serverAv = u.AvatarURL("1024")
				} else {
					serverAv = m.AvatarURL("1024")
				}

				username = u.Username
			} else {
				return fmt.Errorf("couldn't get user from option")
			}
		} else {
			if i.Member != nil {
				globalAv = i.Member.User.AvatarURL("1024")
				serverAv = i.Member.AvatarURL("1024")

				username = i.Member.User.Username
			} else if i.User != nil {
				globalAv = i.User.AvatarURL("1024")
				serverAv = globalAv

				username = i.User.Username
			} else {
				return fmt.Errorf("couldn't get member who executed avatar command")
			}
		}

		respondEmbed := &discordgo.MessageEmbed{
			Title: fmt.Sprintf("%s %s's Avatar", assets.Icons.Info, username),
			Color: assets.Colors.Primary,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL:    globalAv,
				Width:  1024,
				Height: 1024,
			},
			Image: &discordgo.MessageEmbedImage{
				URL:    serverAv,
				Width:  1024,
				Height: 1024,
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
	},
}

func init() {
	registry.Register(&include.Command{
		Data:    Avatar.Data,
		Handler: Avatar.Handler,
	})
}
