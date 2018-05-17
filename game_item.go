package main

type GameItem struct {
	name string
	actions []GameItemAction
}

func (g *GameItem) getName() string {
	return g.name
}