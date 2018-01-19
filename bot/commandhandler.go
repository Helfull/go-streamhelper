package bot

type CommandHandler struct {
    Use string
    Run func(cmd *Command, args []string)
}