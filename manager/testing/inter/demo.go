package _interface

type Animal interface {
	run(s string) string
}

type Dog struct {
	Name string
}

func (d *Dog) run(s string) string {
	return "dog content: " + d.Name + " - " + s
}

type Cat struct {
	Name string
}

func (c *Cat) run(s string) string {
	return "cat content: " + c.Name + " + " + s
}

