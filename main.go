package main

import (
	"flag"
	"fmt"
	"log"

	ckgamerollbot "github.com/TauTheProgrammer/CKGameRollBot/CKGameRollBot"

	"github.com/bwmarrin/discordgo"
)

// TODO Figure out what to do with these variables
var (
	CKGuildID       = flag.String("guild", "532945585238441994", "Cool Kids Guild ID")
	BotChannelID    = flag.String("channel", "755428432921493594", "Channel to listen for slash commands")
	DiscordBotToken = flag.String("token", "MTA0NDAxMzIwMTc0ODE0NDE5OA.G6VpCn.iHyamaOt9hyH4kuRDPk5vQyyeOvf2iExJTYTEc", "Bot access token")
)

var session *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error
	session, err := discordgo.New("Bot " + *DiscordBotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

var commands = []*discordgo.ApplicationCommand{
	{
		GuildID:     "532945585238441994",
		Name:        "roll",
		Description: "Enter a list of games, separated by commas, and I will pick one at random",
		Type:        discordgo.ChatApplicationCommand,
	},
}

func main() {
	fmt.Println(ckgamerollbot.ABC())
	session, err := discordgo.New("")
}
