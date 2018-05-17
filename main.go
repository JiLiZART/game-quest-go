package main

import (
	"strings"
	"gopkg.in/AlecAivazis/survey.v1"
	"fmt"
)

var world *World

func initGame() {
	world = CreateWorld()
}

func updateGame() {
	answer := ""
	prompt := &survey.Select{
		Message: "Выбери действие:",
		Options: world.GetActionNames(),
	}

	survey.AskOne(prompt, &answer, nil)

	fmt.Println(handleActionInteractive(answer))

	updateGame()
}

func handleActionInteractive(actionName string) string {
	return world.HandleActionInteractive(actionName)
}

func handleCommand(name string) string {
	if len(name) <= 0 {
		return "неизвестная команда"
	}

	args := strings.Fields(name)

	return world.HandleAction(args[0], args[1:])
}

func main(){
	initGame()

	updateGame()

	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	fmt.Println(handleCommand(scanner.Text()))
	//}
}
