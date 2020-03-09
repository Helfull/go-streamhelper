package bot

// CommandHandler represents the structure
// for a Command
type CommandHandler struct {
	// Use is the trigger for
	// the command to be executed
	Use string

	// Run is the function to execute
	// if the trigger is used
	Run func(cmd *Command, args []string)
}
