package gameQuest

import "github.com/kataras/iris/core/errors"

type World struct {
	Player *Player

	PlaceHall *Place
	PlaceKitchen *Place
	PlaceRoom *Place
	PlaceStreet *Place
	PlaceHome *Place

	Actions []Action
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

	w.Actions = []Action{
		createActionLookAround(w),
		createActionWalk(w),
		createActionTake(w),
		createActionApply(w),
		createActionWear(w),
	}

	return w
}

func (w *World) GetActionByName(name string) (Action, error) {
	for _, action := range w.Actions {
		if action.IsMatch(name) {
			return action, nil
		}
	}

	return nil, errors.New("действие не найдено")
}

func (w *World) GetActionNames() []string {
	actionNames := make([]string, len(w.Actions))

	for i, action := range w.Actions {
		actionNames[i] = action.GetName()
	}

	return actionNames
}

func (w *World) HandleActionInteractive(name string) string {
	for _, action := range w.Actions {
		if action.IsMatch(name) {
			return action.ExecuteInteractive()
		}
	}

	return "неизвестная команда"
}

func (w *World) HandleAction(name string, args []string) string {
	for _, action := range w.Actions {
		if action.IsMatch(name) {
			return action.Execute(args)
		}
	}

	return "неизвестная команда"
}