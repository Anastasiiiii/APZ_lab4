package engine

import (
	"fmt"
	"strconv"
)

type Command interface {
	Execute(h Hadler)
}

type Handler interface {
	Post(cmd Command) error
}

type PrintCommand string

func (p *PrintCommand) Execute(loop Handler) {
	fmt.Println(p.arg)
}

type AddCommand struct {
	a, b int
}

func (add AddCommand) Execute(h Handler) {
	res := add.a + add.b
	h.Post(PrintCommand(strconv.Itoa(res)))
}

type StopCommand struct{}

func (sc StopCommand) Execute(h Handler) {
	h.(*EventLoop).stop = true
}
