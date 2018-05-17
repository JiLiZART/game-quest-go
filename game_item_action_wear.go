package main

type WearGameItemAction struct {
	callback func(item *GameItem, p *Player)
}

func (a WearGameItemAction) Execute(item *GameItem, p *Player) string {
	if (a.callback != nil) {
		a.callback(item, p)
	}

	return "вы одели: " + item.getName()
}
