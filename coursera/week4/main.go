package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	locomotion, food, noise string
}
type Bird struct {
	locomotion, food, noise string
}
type Snake struct {
	locomotion, food, noise string
}

func (cow Cow) Move() {
	fmt.Println(cow.locomotion)
}

func (cow Cow) Eat() {
	fmt.Println(cow.food)
}

func (cow Cow) Speak() {
	fmt.Println(cow.noise)
}

func (bird Bird) Move() {
	fmt.Println(bird.locomotion)
}

func (bird Bird) Eat() {
	fmt.Println(bird.food)
}

func (bird Bird) Speak() {
	fmt.Println(bird.noise)
}

func (snake Snake) Move() {
	fmt.Println(snake.locomotion)
}

func (snake Snake) Eat() {
	fmt.Println(snake.food)
}

func (snake Snake) Speak() {
	fmt.Println(snake.noise)
}

type AnimalsMap map[string]Animal

var animalsMap AnimalsMap = AnimalsMap{"cow": Cow{locomotion: "walk", food: "grass", noise: "moo"},
	"bird":  Bird{locomotion: "fly", food: "warms", noise: "peep"},
	"snake": Snake{locomotion: "slither", food: "mice", noise: "hsss"}}

type Animals map[string]Animal

var animals Animals = make(map[string]Animal)

// type AnimalsInfo map[string]func()
// var animalsInfo AnimalsInfo = AnimalsInfo{"move": }
func CreateAnimal(animalName string, animalType string) {
	_, ok := animals[animalName]
	if ok {
		fmt.Printf("Animal with name %s already exists", animalName)
		return
	}
	animals[animalName] = animalsMap[animalType]
	fmt.Println("Created it!")
}

func QueryAnimal(animalName string, animalInfo string) {
	animal := animals[animalName]
	switch animalInfo {
	case "move":
		animal.Move()
	case "eat":
		animal.Eat()
	case "speak":
		animal.Speak()
	}
}

func main() {
	var command, animalName, animalType, animalInfo string
	for true {
		fmt.Print("> ")
		fmt.Scan(&command, &animalName)

		switch command {
		case "newanimal":
			fmt.Scan(&animalType)
			CreateAnimal(animalName, animalType)

		case "query":
			fmt.Scan(&animalInfo)
			QueryAnimal(animalName, animalInfo)
		}
	}
}
