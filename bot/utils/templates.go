package utils

import "fmt"

var (
	appBrickieGo     = "https://git.io/JUvRv"
	appBrickieGithub = "https://git.io/JUvRm"
)

func GetHelpTemplate(username string) string {
	var helpTemplate = fmt.Sprintf(
		"Welcome to AppBrickie %s ! \n\n"+
			"Commands that you can use here or in a group chat with the bot added: \n\n"+
			"1. /getid - This returns your unique id to put in the YAML file to recieve your apk builds here \n\n"+
			"2. /greet - A simple normal function that greets you politely! \n\n"+
			"3. /help - To View a list of commands \n\n\n"+
			"Commads that you can use when you add the bot as an admin in a channel to send automated builds\n\n"+
			"1. !getid - Returns the unique id for the channel to put int he YAML file \n\n"+
			"2. !help - To Print a list of commands in a channel \n\n\n"+
			"If you like this bot , feel free to start this project on Github %s \n\n"+
			"You can find more details and instructions on %s \n", username, appBrickieGithub, appBrickieGo)
	return helpTemplate
}
