package bot

type Command struct {
    Raw []string
    Name string
    Args []string
    User *User
    Bot *Bot
    Msg *Message
}


func (cmd *Command) ReplyPrivate(message string) {
    cmd.Bot.Reply(cmd.User.id, message)
}

func (cmd *Command) ReplyChannel(message string) {
    cmd.Bot.ReplyChannel(message)
}