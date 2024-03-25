package tasks

import "fmt"

type Command interface {
	Execute() int
}

type AddCommand struct {
	a, b int
}

func (c *AddCommand) Execute() int {
	return c.a + c.b
}

type SubtractCommand struct {
	a, b int
}

func (c *SubtractCommand) Execute() int {
	return c.a - c.b
}

type calculator struct {
}

func (c *calculator) Compute(command Command) int {
	return command.Execute()
}

func Task2() {
	calc := calculator{}

	add := &AddCommand{3, 2}
	subtract := &SubtractCommand{10, 5}

	fmt.Println(calc.Compute(add))
	fmt.Println(calc.Compute(subtract))
}
