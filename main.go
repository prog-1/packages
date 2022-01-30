package main

import (
	"fmt"
	"math"

	"github.com/prog-1/packages/areas"
	"github.com/yarcat/fsm-go"
)

func main() {
	fmt.Println("figure   : got, want")
	fmt.Println("--------------------")
	fmt.Println("circle   :", areas.Circle(2), 2*2*math.Pi)
	fmt.Println("square   :", areas.Square(2), 2*2)
	fmt.Println("rectangle:", areas.Rectangle(2, 3), 2*3)

	const (
		stInit        fsm.StateType = "stInit"
		stFinal       fsm.StateType = "stFinal"
		evInitialized fsm.EventType = "evInitialized"
	)

	transitions := fsm.Transitions{
		fsm.When(stInit, evInitialized): stFinal,
	}

	fsm := fsm.New(stInit, transitions, nil, nil)
	fsm.Send(evInitialized)
}
