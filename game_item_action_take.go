package main

type TakeGameItemAction struct {
	callback func(item *GameItem, p *Player)
}

func (a TakeGameItemAction) Execute(item *GameItem, p *Player) string {
	if a.callback != nil {
		a.callback(item, p)
	}

	return "предмет добавлен в инвентарь: " + item.getName()
}
