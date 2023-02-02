package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %v: %v", g.LanguageName(), g.Greet(name))
}

type Italian struct {
}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %v!", name)
}

type Portuguese struct {
}

func (i Portuguese) LanguageName() string {
	return "Portuguese"
}

func (i Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %v!", name)
}
