package gameQuest

type ActionApply struct {
	name string
	world *World
}

func createActionApply(world *World) *ActionApply {
	action := &ActionApply{name: "применить", world: world}

	return action
}

func (a *ActionApply) IsMatch(name string) bool {
	return a.name == name
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



