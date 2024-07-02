package main

import "fmt"

type Animal struct {
	locomotion, food, noise string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

type Animals map[string]Animal
type AnimalsInfo map[string]func()

var animal Animal
var animalsInfo AnimalsInfo = AnimalsInfo{"move": animal.Move, "speak": animal.Speak, "eat": animal.Eat}

func (animalInfo AnimalsInfo) Update(animal Animal) {
	animalInfo["move"] = animal.Move
	animalInfo["speak"] = animal.Speak
	animalInfo["eat"] = animal.Eat

}

func main() {
	var cow, bird, snake Animal
	cow = Animal{locomotion: "walk", food: "grass", noise: "moo"}
	bird = Animal{locomotion: "fly", food: "warms", noise: "peep"}
	snake = Animal{locomotion: "slither", food: "mice", noise: "hsss"}
	animals := Animals{"cow": cow, "bird": bird, "snake": snake}

	var animalName, animalInfo string
	for true {
		fmt.Print("> ")
		fmt.Scan(&animalName, &animalInfo)
		animal = animals[animalName]
		animalsInfo.Update(animal)
		animalsInfo[animalInfo]()
	}

}
