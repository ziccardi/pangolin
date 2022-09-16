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

type MemoryCellOption func(*MemoryCell)

func NewMemoryCell(options ...MemoryCellOption) *MemoryCell {
	cell := &MemoryCell{}
	for _, o := range options {
		o(cell)
	}
	return cell
}

func UpdateMemoryCell(cell *MemoryCell, options ...MemoryCellOption) {
	for _, o := range options {
		o(cell)
	}
}

func WithAnimal(animal string) MemoryCellOption {
	return func(cell *MemoryCell) {
		if animal != "" {
			cell.Animal = &animal
		} else {
			cell.Animal = nil
		}
	}
}

func WithQuestion(question string) MemoryCellOption {
	return func(cell *MemoryCell) {
		if question != "" {
			cell.Question = &question
		} else {
			cell.Question = nil
		}
	}
}

func WithYes(branch *MemoryCell) MemoryCellOption {
	return func(cell *MemoryCell) {
		cell.Yes = branch
	}
}

func WithNo(branch *MemoryCell) MemoryCellOption {
	return func(cell *MemoryCell) {
		cell.No = branch
	}
}

func initMemory() *MemoryCell {

	whaleAnswer := NewMemoryCell(WithAnimal("a whale"))
	antAnswer := NewMemoryCell(WithAnimal("an ant"))
	pangolinAnswer := NewMemoryCell(WithAnimal("a pangolin"))
	blancMange := NewMemoryCell(WithAnimal("a blancmange"))

	eatAntsQuestion := NewMemoryCell(
		WithQuestion("Does it eat ants?"),
		WithYes(pangolinAnswer),
		WithNo(antAnswer),
	)

	haveScaleArmourQuestion := NewMemoryCell(
		WithQuestion("Is it scaly?"),
		WithYes(eatAntsQuestion),
		WithNo(blancMange),
	)

	liveInTheWaterQuestion := NewMemoryCell(
		WithQuestion("Does it live in the sea?"),
		WithYes(whaleAnswer),
		WithNo(haveScaleArmourQuestion),
	)

	return liveInTheWaterQuestion
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
	if yes {
		UpdateMemoryCell(
			cell,
			WithAnimal(""),
			WithQuestion(question),
			WithYes(&newAnimal),
			WithNo(&oldAnimal),
		)
	} else {
		UpdateMemoryCell(
			cell,
			WithAnimal(""),
			WithQuestion(question),
			WithYes(&oldAnimal),
			WithNo(&newAnimal),
		)
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
