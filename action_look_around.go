package gameQuest


type ActionLookAround struct {
	name string
	world *World
}

func createActionLookAround(world *World) *ActionLookAround {
	action := &ActionLookAround{name: "осмотреться", world: world}

	return action
}

func (a *ActionLookAround) IsMatch(name string) bool {
	return a.name == name
}

func (a *ActionLookAround) Execute(args []string) string {
	p := a.world.Player.Place

	return p.LookAround()
}

