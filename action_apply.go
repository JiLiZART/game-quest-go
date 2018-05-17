package main

import "gopkg.in/AlecAivazis/survey.v1"

type ActionApply struct {
	name string
	world *World
}

func createActionApply(world *World) *ActionApply {
	action := &ActionApply{name: "применить", world: world}

	return action
}

func (a *ActionApply) GetName() string {
	return a.name
}

func (a *ActionApply) IsMatch(name string) bool {
	return a.name == name
}

func (a *ActionApply) ExecuteInteractive() string {
	player := a.world.Player

	gameItemName := ""
	prompt := &survey.Select{
		Message: "Применить:",
		Options: player.GetItemNames(),
	}

	survey.AskOne(prompt, &gameItemName, nil)

	return a.Execute([]string{gameItemName})
}

func (a *ActionApply) Execute(args []string) string {
	player := a.world.Player

	if len(args) > 0 {
		gameItemName := args[0]

		if gameItem, exist := player.getItemByName(gameItemName); exist {
			return player.applyGameItem(gameItem)
		}

		return "нет предмета в инвентаре - " + gameItemName
	}

	return ""
}



