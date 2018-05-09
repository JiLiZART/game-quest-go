package gameQuest

type World struct {
	Player *Player

	PlaceHall *Place
	PlaceKitchen *Place
	PlaceRoom *Place
	PlaceStreet *Place
	PlaceHome *Place

	actions []Action
}

func CreateWorld() *World {
	w := &World{}

	w.Player = createPlayer(w)
	w.PlaceHome = createHome(w)
	w.PlaceHall = createHall(w)
	w.PlaceKitchen = createKitchen(w)
	w.PlaceRoom = createRoom(w)
	w.PlaceStreet = createStreet(w)

	w.Player.Place = w.PlaceKitchen

	w.PlaceKitchen.Exits = []*Door{
		{destination: w.PlaceHall},
	}
	w.PlaceHall.Exits = []*Door{
		{destination: w.PlaceKitchen},
		{destination: w.PlaceRoom},
		{destination: w.PlaceStreet, isLocked: true},
	}
	w.PlaceRoom.Exits = []*Door{
		{destination: w.PlaceHall},
	}
	w.PlaceStreet.Exits = []*Door{
		{destination: w.PlaceHome},
	}

	w.actions = []Action{
		createActionLookAround(w),
		createActionWalk(w),
		createActionTake(w),
		createActionApply(w),
		createActionWear(w),
	}

	return w
}

func (w *World) HandleAction(name string, args []string) string {

	for _, action := range w.actions {
		if action.IsMatch(name) {
			return action.Execute(args)
		}
	}

	return "неизвестная команда"
}