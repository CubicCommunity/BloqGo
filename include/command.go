package include

import "github.com/bwmarrin/discordgo"

type Command struct {
	Data    *discordgo.ApplicationCommand                                                                              // Discord slash command data
	Handler func(s *discordgo.Session, c *discordgo.ApplicationCommandInteractionData, i *discordgo.Interaction) error // Handler function for this command
}

func NSFW(b bool) *bool { return &b }
