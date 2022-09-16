package main

// Don't run this into the playground: it won't work. Playground doesn't support reading from stdin.

import (
	"bufio"
	"fmt"
	"os"
	. "pangolini/memory"
	"strings"
)

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSuffix(line, "\n")
}

func askQuestion(header string, question string) bool {
	for true {
		if header != "" {
			fmt.Println(header)
		}
		fmt.Println(question)
		answer := readLine()
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
	animal := strings.ToLower(readLine())
	fmt.Println("Tell me a question that distinguishes between")
	fmt.Printf(" %s and\n", *cell.Animal)
	fmt.Printf(" %s\n", animal)

	question := readLine()

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
	mem := InitAnimals()
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
				mem.Save("./pangolins.dat")
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
				mem.Save("./pangolins.dat")
				break
			}
			cell = cell.No
		}
	}
}
