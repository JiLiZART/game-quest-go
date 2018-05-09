package gameQuest

type ActionTake struct {
	name string
	world *World
}

func createActionTake(w *World) *ActionTake {
	action := &ActionTake{name: "взять", world: w}

	return action
}

func (a *ActionTake) IsMatch(name string) bool {
	return a.name == name
}

func (a *ActionTake) Execute(args []string) string {
	player := a.world.Player
	place := player.Place

	if !player.hasInventory {
		return "некуда класть"
	}

	if len(args) > 0 {
		gameItemName := args[0]

		if gameItem, exist := place.getItemByName(gameItemName); exist {
			player.Place.removeGameItem(gameItem)
			return player.takeGameItem(gameItem)
		}
	}

	return "нет такого"
}

