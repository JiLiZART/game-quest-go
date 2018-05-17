package main

type Action interface {
	GetName() string
	IsMatch(name string) bool
	Execute(args []string) string
	ExecuteInteractive() string
}