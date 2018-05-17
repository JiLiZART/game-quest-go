package gameQuest

import (
	"strings"
)

type Place struct {
	id string
	description string
	enteringMessage string
	itemPlaces []*PlaceItem
	Exits []*Door
}

func createHall(w *World) *Place {
	place := &Place{
		id: "коридор",
		enteringMessage: "ничего интересного.",
	}

	return place
}

func createHome(w *World) *Place {
	place := &Place{
		id: "домой",
		enteringMessage: "вот и дома.",
	}

	return place
}

func createKitchen(w *World) *Place {
	place := &Place{
		id: "кухня",
		enteringMessage: "кухня, ничего интересного.",
		description: "ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ. ",
	}

	return place
}

func createRoom(w *World) *Place {
	place := &Place{
		id: "комната",
		enteringMessage: "ты в своей комнате.",
		itemPlaces: []*PlaceItem{
			{
				name: "стол",
				description: "на столе: ",
				gameItems: []*GameItem{
					{
						name: "ключи",
						actions: []GameItemAction {
							TakeGameItemAction {},
							ApplyGameItemAction {
								callback: func(item *GameItem, p *Player) string {
									for _, door := range p.Place.Exits {
										if door.isLocked  {
											door.isLocked = false
											return "дверь открыта"
										}
									}

									return "не к чему применить"
								},
							},
						},
					},
					{
						name: "конспекты",
						actions: []GameItemAction {
							TakeGameItemAction {},
						},
					},
				},
			},
			{
				name: "стул",
				description:"на стуле - ",
				gameItems: []*GameItem{
					{
						name: "рюкзак",
						actions: []GameItemAction {
							WearGameItemAction {
								callback: func(item *GameItem, p *Player) {
									p.hasInventory = true
								},
							},
						},
					},
				},
			},
		},
	}

	return place
}

func createStreet(w *World) *Place {
	place := &Place{
		id: "улица",
		enteringMessage: "на улице весна.",
	}

	return place
}

func (p *Place) ExitsMessage() string {
	result := make([]string, len(p.Exits))

	if len(p.Exits) <= 0 {
		return ""
	}

	for idx, door := range p.Exits {
		result[idx] = door.destination.Name()
	}

	return "можно пройти - " + strings.Join(result[:],", ")
}

func (p *Place) ItemsMessage() string {
	var result []string

	if len(p.getGameItems()) <= 0 {
		return ""
	}

	for _, itemPLace := range p.itemPlaces {
		if len(itemPLace.gameItems) > 0 {
			result = append(result, itemPLace.getGameItemsMessage())
		}
	}

	return strings.Join(result[:],", ") + ". "
}

func (p *Place) getNameItemsNames() []string {
	names := make([]string, len(p.getGameItems()))

	for i, item := range p.getGameItems() {
		names[i] = item.getName()
	}

	return names
}

func (p *Place) getGameItems() []*GameItem {
	var gameItems []*GameItem

	for _, itemPLace := range p.itemPlaces {
		gameItems = append(gameItems, itemPLace.gameItems...)
	}

	return gameItems
}

func (p *Place) getItemByName(name string) (*GameItem, bool) {
	gameItems := p.getGameItems()

	for _, gameItem := range gameItems {
		if gameItem.getName() == name {
			return gameItem, true
		}
	}

	return nil, false
}

func (p *Place) removeGameItem(item *GameItem) bool {

	for _, itemPLace := range p.itemPlaces {
		if itemPLace.removeGameItem(item) {
			return true
		}
	}

	return false
}

func (p *Place) Name() string {
	return p.id
}

func (p *Place) LookAround() string {
	string := p.description + p.ItemsMessage()

	if len(string) == 0 {
		string = "пустая комната. "
	}
	return string + p.ExitsMessage()
}

func (p *Place) EnteringMessage() string {
	return p.enteringMessage + " " + p.ExitsMessage()
}

func (p *Place) ExitsNames() []string {
	names := make([]string, len(p.Exits))

	for i, door := range p.Exits {
		names[i] = door.destination.Name()
	}

	return names
}