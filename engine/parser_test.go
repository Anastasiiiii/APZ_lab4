package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	res := Parse("print hello")
	assert.Equal(t, &PrintCommand{arg: "hello"}, res)

	res = Parse("princ hello")
	assert.Equal(t, &ErrorMessage{message: "Unknown instruction"}, res)

	res = Parse("add 1 2")
	assert.Equal(t, &AddCommand{arg1: 1, arg2: 2}, res)

	res = Parse("add one 2")
	assert.Equal(t, &ErrorMessage{message: "strconv.Atoi: parsing \"one\": invalid syntax"}, res)

	res = Parse("add 1 two")
	assert.Equal(t, &ErrorMessage{message: "strconv.Atoi: parsing \"two\": invalid syntax"}, res)

	res = Parse("add 2")
	assert.Equal(t, &ErrorMessage{message: "`add` command requires exactly 2 arguments"}, res)

	res = Parse("add 1 2 3")
	assert.Equal(t, &ErrorMessage{message: "`add` command requires exactly 2 arguments"}, res)

	res = Parse("")
	assert.Equal(t, &ErrorMessage{message: "Incorrect number of args or no command. "}, res)
}
