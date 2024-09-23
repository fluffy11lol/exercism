package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(s string) string
}
type Italian struct {
}

func (it Italian) Greet(s string) string {
	return fmt.Sprintf("Ciao %s!", s)
}
func (it Italian) LanguageName() string {
	return "Italian"
}

type Portuguese struct {
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(s string) string {
	return fmt.Sprintf("Olá %s!", s)
}

func SayHello(s string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(s))
}
