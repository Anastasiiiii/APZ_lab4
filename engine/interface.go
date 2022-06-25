package engine

import (
	"fmt"
	"strconv"
)

type Command interface {
	Execute(h Handler)
}

type Handler interface {
	Post(cmd Command) error
}

type ErrorMessage struct {
	message string
}

func (e *ErrorMessage) Execute(loop Handler) {
	fmt.Printf("SYNTAX ERROR: %s\n", e.message)
}

type PrintCommand struct {
	arg string
}

func (p *PrintCommand) Execute(loop Handler) {
	fmt.Println(p.arg)
}

type AddCommand struct {
	arg1, arg2 int
}

func (a *AddCommand) Execute(h Handler) {
	res := a.arg1 + a.arg2
	h.Post(&PrintCommand{arg: strconv.Itoa(res)})
}

type StopCommand struct{}

func (sc StopCommand) Execute(h Handler) {
	h.(*EventLoop).stop = true
}
