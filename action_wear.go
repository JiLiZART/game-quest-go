package gameQuest

type ActionWear struct {
	name string
	world *World
}

func createActionWear(w *World) *ActionWear {
	action := &ActionWear{name: "одеть", world: w}

	return action
}

func (a *ActionWear) IsMatch(name string) bool {
	return a.name == name
}

func (a *ActionWear) Execute(args []string) string {
	player := a.world.Player
	place := player.Place

	if len(args) > 0 {
		gameItemName := args[0]

		if gameItem, exist := place.getItemByName(gameItemName); exist {
			player.Place.removeGameItem(gameItem)
			return player.wearGameItem(gameItem)
		}
	}

	return "нечего брать"
}

