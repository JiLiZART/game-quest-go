package main

import "gopkg.in/AlecAivazis/survey.v1"

func createActionWalk(w *World) *ActionWalk {
	action := &ActionWalk{name: "идти", world: w}

	return action
}

type ActionWalk struct {
	name string
	world *World
}

func (a *ActionWalk) GetName() string {
	return a.name
}

func (a *ActionWalk) IsMatch(name string) bool {
	return a.name == name
}

func (a *ActionWalk) ExecuteInteractive() string {
	player := a.world.Player

	name := ""
	prompt := &survey.Select{
		Message: "Идти:",
		Options: player.Place.ExitsNames(),
	}

	survey.AskOne(prompt, &name, nil)

	return a.Execute([]string{name})
}

func (a *ActionWalk) Execute(args []string) string {
	world := a.world

	if len(args) > 0 {
		placeName := args[0]

		for _, door := range world.Player.Place.Exits {
			if door.destination.Name() == placeName {
				world.Player.Place = door.destination

				return world.Player.Place.EnteringMessage()
			}
		}

		return "нет пути в " + placeName
	}

	return "некуда идти"
}

