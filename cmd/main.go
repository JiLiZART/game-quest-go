package main

import (
	"fmt"
	"strings"
	gQ "gameQuest"
	"os"
	"bufio"
)

var world *gQ.World

func initGame() {
	world = gQ.CreateWorld()
}

func handleCommand(name string) string {
	if len(name) <= 0 {
		return "неизвестная команда"
	}

	args := strings.Fields(name)

	return world.HandleAction(args[0], args[1:])
}

func main() {
	fmt.Println(">> ")

	initGame()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(handleCommand(scanner.Text()))
	}
}
