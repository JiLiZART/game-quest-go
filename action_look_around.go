package main


type ActionLookAround struct {
	name string
	world *World
}

func createActionLookAround(world *World) *ActionLookAround {
	action := &ActionLookAround{name: "осмотреться", world: world}

	return action
}

func (a *ActionLookAround) GetName() string {
	return a.name
}

func (a *ActionLookAround) IsMatch(name string) bool {
	return a.name == name
}

func (a *ActionLookAround) ExecuteInteractive() string {
	return a.world.Player.Place.LookAround()
}

func (a *ActionLookAround) Execute(args []string) string {
	return a.world.Player.Place.LookAround()
}

