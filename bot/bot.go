package bot

import (
	"fmt"
	"strings"
)

type Channels struct {
	Errors   chan error
	Commands chan *Command
}

type Server struct {
	Domain string
}

type Settings struct {
	Debug bool

	Nickname string
	Oauth    string
	Channel  string

	Server          Server
	CommandHandlers []*CommandHandler
}

type Bot struct {
	Settings *Settings
	irccon   *Connection
	Channels *Channels
}

func (bot *Bot) handleMessage(msg *Message) {
	if strings.HasPrefix(msg.Text, "!") {
		bot.onCommand(bot.makeCommand(msg))
	}
}

func Create(configFile string, cmdHandlers []*CommandHandler) *Bot {
	config := GetConfig("./settings.json")

	settings := &Settings{}
	settings.Debug = config.Debug
	settings.Nickname = config.Nickname
	settings.Oauth = config.Oauth
	settings.Server = Server{}
	settings.Server.Domain = config.Server
	settings.Channel = config.Channel
	settings.CommandHandlers = cmdHandlers

	bot := Bot{}
	bot.Channels = &Channels{make(chan error), make(chan *Command)}
	bot.Settings = settings
	bot.irccon = NewConnection(&bot)

	return &bot
}

func (bot *Bot) Loop() {
	bot.irccon.Connect()
	bot.irccon.On(MSG_PRIVMSG, func(msg *Message) {
		switch msg.Type {
		case MSG_PRIVMSG:
			bot.handleMessage(msg)
		case MSG_UNKNOWN:
			fmt.Printf("Unknown Message Type", msg.raw)
		}
	})
	bot.irccon.Loop()
}

func (bot *Bot) onCommand(command *Command) {
	for _, cmdHandler := range bot.Settings.CommandHandlers {
		if cmdHandler.Use == command.Name {
			fmt.Printf("Command -> %s\n", command.Name)
			cmdHandler.Run(command, command.Args)
			break
		}
	}
}

func (bot *Bot) Reply(channel string, message string) *Bot {
	bot.irccon.Send(channel, message)
	return bot
}
func (bot *Bot) ReplyChannel(message string) *Bot {
	return bot.Reply(bot.Settings.Channel, message)
}

func (bot *Bot) makeCommand(msg *Message) *Command {
	command := &Command{}

	commandSplit := strings.Split(strings.TrimLeft(msg.Text, "!"), " ")
	command.Raw = commandSplit
	command.Name = strings.Join(commandSplit[0:1], "")
	command.Args = commandSplit[1:]
	command.Msg = msg
	command.User = msg.User
	command.Bot = bot

	return command
}
