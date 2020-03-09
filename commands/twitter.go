package cmds

import (
	"github.com/helfull/go-streamhelper/bot"
)

var TwitterCmd = &bot.CommandHandler{
	Use: "twitter",
	Run: func(cmd *bot.Command, args []string) {
		cmd.ReplyChannel("https://twitter.com/helfuli")
	},
}
