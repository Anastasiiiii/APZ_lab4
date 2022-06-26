package engine

import (
  "strconv"
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestLoop(t *testing.T) {
  printCmd1 := &PrintCommand{
    arg: "first test",
  }
  printCmd2 := &PrintCommand{
    arg: "second test",
  }
  addCmd := &AddCommand{
    arg1: 1,
    arg2: 2,
  }

  eventLoop := new(EventLoop)
  eventLoop.Start()
  assert.Equal(t, false, eventLoop.stopped)
  assert.Equal(t, 0, len(eventLoop.q.commands))

  eventLoop.Post(addCmd)
  eventLoop.Post(printCmd1)
  eventLoop.Post(printCmd2)

  assert.Equal(t, 3, len(eventLoop.q.commands))
  eventLoop.AwaitFinish()
  assert.Equal(t, true, eventLoop.stopped)
  assert.Equal(t, 0, len(eventLoop.q.commands))

  assert.Equal(t, printCmd1, &PrintCommand{arg: "first test"})
  assert.Equal(t, printCmd2, &PrintCommand{arg: "second test"})
  assert.Equal(t, strconv.Itoa(addCmd.arg1+addCmd.arg2), "3")
}