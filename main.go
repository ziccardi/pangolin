package main

// Don't run this into the playground: it won't work. Playground doesn't support reading from stdin.

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MemoryCell struct {
	No       *MemoryCell
	Yes      *MemoryCell
	Animal   *string
	Question *string
}

func initMemory() *MemoryCell {
	stringPointer := func(s string) *string {
		return &s
	}

	liveInTheWaterQuestion := MemoryCell{
		Question: stringPointer("Does it live in the sea?"),
	}
	eatAntsQuestion := MemoryCell{
		Question: stringPointer("Does it eat ants?"),
	}
	haveScaleArmourQuestion := MemoryCell{
		Question: stringPointer("Is it scaly?"),
	}

	whaleAnswer := MemoryCell{
		Animal: stringPointer("a whale"),
	}
	antAnswer := MemoryCell{
		Animal: stringPointer("an ant"),
	}
	pangolinAnswer := MemoryCell{
		Animal: stringPointer("a pangolin"),
	}
	blancMange := MemoryCell{
		Animal: stringPointer("a blancmange"),
	}

	liveInTheWaterQuestion.Yes = &whaleAnswer
	liveInTheWaterQuestion.No = &haveScaleArmourQuestion
	haveScaleArmourQuestion.Yes = &eatAntsQuestion
	haveScaleArmourQuestion.No = &blancMange

	eatAntsQuestion.Yes = &pangolinAnswer
	eatAntsQuestion.No = &antAnswer

	return &liveInTheWaterQuestion
}

func askQuestion(header string, question string) bool {
	for true {
		if header != "" {
			fmt.Println(header)
		}
		fmt.Println(question)
		reader := bufio.NewReader(os.Stdin)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSuffix(strings.ToLower(answer), "\n")
		if answer == "y" || answer == "yes" {
			return true
		}
		if answer == "n" || answer == "no" {
			return false
		}

		fmt.Println("Answer me property when I'm talking to you")
	}
	return false
}

func addNewQuestion(cell *MemoryCell) {
	fmt.Println("What is it then?")
	reader := bufio.NewReader(os.Stdin)
	animal, _ := reader.ReadString('\n')
	animal = strings.TrimSuffix(strings.ToLower(animal), "\n")
	fmt.Println("Tell me a question that distinguishes between")
	fmt.Printf(" %s and\n", *cell.Animal)
	fmt.Printf(" %s\n", animal)

	question, _ := reader.ReadString('\n')
	question = strings.TrimSuffix(question, "\n")

	yes := askQuestion("What is the answer for", fmt.Sprintf(" %s?", animal))

	newAnimal := MemoryCell{Animal: &animal}
	oldAnimal := *cell

	// replace current animal with the new question
	newQuestion := cell
	newQuestion.Animal = nil
	newQuestion.Question = &question
	if yes {
		newQuestion.Yes = &newAnimal
		newQuestion.No = &oldAnimal
	} else {
		newQuestion.Yes = &oldAnimal
		newQuestion.No = &newAnimal
	}

}

func main() {
	mem := initMemory()
	cell := mem
	for true {

		var yes bool

		if cell.Question != nil {
			yes = askQuestion("", *cell.Question)
		} else {
			yes = askQuestion("Are you thinking of", fmt.Sprintf(" %s?", *cell.Animal))
		}

		if yes {
			if cell.Yes == nil {
				// I won
				fmt.Println("I thought as much")
				if y := askQuestion("", "Do you want another go?"); y {
					cell = mem
					continue
				}
				break
			}
			cell = cell.Yes
		} else {
			if cell.No == nil {
				// Need a new question
				addNewQuestion(cell)
				fmt.Println("That fooled me.")
				if y := askQuestion("", "Do you want another go?"); y {
					cell = mem
					continue
				}
				break
			}
			cell = cell.No
		}
	}
}
