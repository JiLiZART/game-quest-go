package main

type GameItemAction interface {
	Execute(item *GameItem, p *Player) string
}

