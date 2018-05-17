package main

type ApplyGameItemAction struct {
	callback func(item *GameItem, p *Player) string
}

func (a ApplyGameItemAction) Execute(item *GameItem, p *Player) string {
	if (a.callback != nil) {
		return a.callback(item, p)
	}

	return ""
}
