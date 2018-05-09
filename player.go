package gameQuest

type Player struct {
	name string
	Place *Place
	hasInventory bool
	inventory []*GameItem
}

func createPlayer(w *World) *Player {
	p := &Player{}

	return p
}

func (p *Player) addToInventory(item *GameItem) string {
	if p.hasInventory {
		p.inventory = append(p.inventory, item)

		p.Place.removeGameItem(item)

		return "предмет добавлен в инвентарь:  " + item.getName()
	}

	return "некуда класть"
}

func (p *Player) wearGameItem(item *GameItem) string {
	p.inventory = append(p.inventory, item)

	for _, action := range item.actions {
		if i, ok := action.(WearGameItemAction); ok {
			return i.Execute(item, p)
		}
	}

	return "нечего надеть"
}

func (p *Player) takeGameItem(item *GameItem) string {
	p.inventory = append(p.inventory, item)

	for _, action := range item.actions {
		if i, ok := action.(TakeGameItemAction); ok {
			return i.Execute(item, p)
		}
	}

	return "нельзя взять"
}

func (p *Player) getItemByName(name string) (*GameItem, bool) {
	if (!p.hasInventory) {
		return nil, false
	}

	for _, gameItem := range p.inventory {
		if gameItem.getName() == name {
			return gameItem, true
		}
	}

	return nil, false
}


func (p *Player) applyGameItem(item *GameItem) string {
	for _, action := range item.actions {
		if i, ok := action.(ApplyGameItemAction); ok {
			return i.Execute(item, p)
		}
	}

	return "нечего применить"
}
