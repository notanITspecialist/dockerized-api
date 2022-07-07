package app

import "fmt"

type (
	Application interface {
		Run()
	}

	application struct{}
)

func NewApplication() Application {
	return &application{}
}

func (app *application) Run() {
	fmt.Println("Start!!")
}
