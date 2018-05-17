package main

import "strings"

type PlaceItem struct {
	name string
	description string
	gameItems []*GameItem
}

func (i *PlaceItem) getGameItemsMessage() string {
	result := make([]string, len(i.gameItems))

	if len(i.gameItems) == 0 {
		return ""
	}

	for idx, item := range i.gameItems {
		result[idx] = item.getName()
	}

	return i.description + strings.Join(result[:],", ")
}

func (p *PlaceItem) removeGameItem(item *GameItem) bool {
	for i, gameItem := range p.gameItems {
		if gameItem == item {
			p.gameItems = append(p.gameItems[:i], p.gameItems[i+1:]...)
			return true
		}
	}

	return false
}