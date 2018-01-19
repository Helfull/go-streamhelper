package main

import (
    // "fmt"

    "github.com/helfull/go-streamhelper/bot"
    "github.com/helfull/go-streamhelper/commands"
)


func registerCommandHandlers(cmdHandlers *[]*bot.CommandHandler) {
    addCommand(cmdHandlers, cmds.TwitterCmd)
}

func main() {
    cmdHandlers := make([]*bot.CommandHandler, 0)
    registerCommandHandlers(&cmdHandlers)

    ircBot := bot.Create("./settings.json", cmdHandlers)
    ircBot.Loop()
}

func addCommand(cmdHandlers *[]*bot.CommandHandler, cmd *bot.CommandHandler) {
    *cmdHandlers = append(*cmdHandlers, cmd)
}
