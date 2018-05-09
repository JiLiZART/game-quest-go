package gameQuest

type Action interface {
	IsMatch(name string) bool
	Execute(args []string) string
}