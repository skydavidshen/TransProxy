package main

import "fmt"

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

type Pig struct {
	Name string
}

func (p *Pig) run(s string) string {
	panic("implement me")
}

func (p *Pig) fly(s string) string {
	return "pig fly: " + p.Name + " + " + s
}

func main()  {
	var a1 Animal
	var a2 Animal

	d := new(Dog)
	c := new(Cat)
	d.Name = "dogwei"
	c.Name = "catzou"

	a1 = d
	a2 = c

	fmt.Printf("dog -> %s \n", a1.run("dog run ..."))
	fmt.Printf("cat -> %s", a2.run("dog run ..."))

}

