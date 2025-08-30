package include

import "github.com/bwmarrin/discordgo"

type Command struct {
	Data    *discordgo.ApplicationCommand
	Handler func(s *discordgo.Session, c *discordgo.ApplicationCommandInteractionData, i *discordgo.Interaction) error
}

func NSFW(b bool) *bool { return &b }
