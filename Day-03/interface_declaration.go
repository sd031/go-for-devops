package main

import "fmt"

// 1. Single-method Interface
type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says Woof!"
}

// 2. Embedding Interfaces
type Animal interface {
	Speaker
	Move() string
}

type Bird struct {
	Name string
}

func (b Bird) Speak() string {
	return b.Name + " says Tweet!"
}

func (b Bird) Move() string {
	return b.Name + " flies"
}

// 3. Empty Interface
func PrintAnything(v interface{}) {
	fmt.Println(v)
}

// 4. Interface with Multiple Methods
type Vehicle interface {
	Start() string
	Stop() string
}

type Car struct {
	Model string
}

func (c Car) Start() string {
	return c.Model + " car started"
}

func (c Car) Stop() string {
	return c.Model + " car stopped"
}

// 5. Functional Interface (Single Method Interface)
// Similar to Speaker interface

// 6. Interface as a Constraint (Generics)
func Describe[T Speaker](t T) {
	fmt.Println(t.Speak())
}

func main() {
	dog := Dog{Name: "Buddy"}
	bird := Bird{Name: "Tweety"}
	car := Car{Model: "Tesla"}

	fmt.Println(dog.Speak())
	fmt.Println(bird.Speak())
	fmt.Println(bird.Move())

	PrintAnything("An empty interface can hold anything")
	PrintAnything(42)

	fmt.Println(car.Start())
	fmt.Println(car.Stop())

	Describe(dog)
	Describe(bird)
}
