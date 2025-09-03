package commands

import (
	"fmt"

	"github.com/CubicCommunity/BloqGo/assets"
	"github.com/CubicCommunity/BloqGo/include"
	"github.com/CubicCommunity/BloqGo/registry"

	"github.com/bwmarrin/discordgo"
)

var Say *include.Command = &include.Command{
	Data: &discordgo.ApplicationCommand{
		Name:        "say",
		Description: "Send a message in a channel",
		IntegrationTypes: &[]discordgo.ApplicationIntegrationType{
			discordgo.ApplicationIntegrationGuildInstall,
		},
		Contexts: &[]discordgo.InteractionContextType{
			discordgo.InteractionContextGuild,
		},
		NSFW: include.NSFW(false),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "message",
				Description: "The message to send in the channel",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
			{
				Name:        "channel",
				Description: "The channel to send a message to.",
				Type:        discordgo.ApplicationCommandOptionChannel,
				ChannelTypes: []discordgo.ChannelType{
					discordgo.ChannelTypeGuildText,
					discordgo.ChannelTypeGuildNews,
					discordgo.ChannelTypeGuildPublicThread,
					discordgo.ChannelTypeGuildPrivateThread,
					discordgo.ChannelTypeGuildNewsThread,
					discordgo.ChannelTypeGuildVoice,
					discordgo.ChannelTypeGuildStageVoice,
				},
				Required: false,
			},
		},
	},
	Handler: func(s *discordgo.Session, c *discordgo.ApplicationCommandInteractionData, i *discordgo.Interaction) error {
		message := c.GetOption("message")

		if message != nil {
			msg := message.StringValue()

			respondEmbed := &discordgo.MessageEmbed{
				Description: msg,
				Color:       assets.Colors.Primary,
			}

			channel := c.GetOption("channel")

			if channel != nil {
				chnl := channel.ChannelValue(s)

				msg, err := s.ChannelMessageSendEmbed(chnl.ID, respondEmbed)

				if err != nil {
					return err
				} else {
					if chnl.ID == i.ChannelID {
						return s.InteractionRespond(i, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Embeds: []*discordgo.MessageEmbed{
									{
										Description: fmt.Sprintf("%s Message sent.", assets.Icons.Check),
										Color:       assets.Colors.Primary,
									},
								},
								Flags: discordgo.MessageFlagsEphemeral,
							},
						})
					} else {
						url := fmt.Sprintf("https://discord.com/channels/%s/%s/%s", chnl.GuildID, chnl.ID, msg.ID)

						return s.InteractionRespond(i, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Embeds: []*discordgo.MessageEmbed{
									{
										Description: fmt.Sprintf("%s Message sent in %s.", assets.Icons.Check, url),
										Color:       assets.Colors.Primary,
									},
								},
							},
						})
					}
				}
			} else {
				return s.InteractionRespond(i, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							respondEmbed,
						},
					},
				})
			}
		} else {
			return fmt.Errorf("couldn't get message for say command")
		}
	},
}

func init() {
	registry.Register(Say)
}
