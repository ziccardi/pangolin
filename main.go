package main

import (
	"bufio"
	"fmt"
	"os"
	"pangolini/memory"
	"strings"
)

const memoryFile = "pangolins.dat"

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

func addNewQuestion(cell *memory.Cell) {
	fmt.Println("What is it then?")
	animal := strings.ToLower(readLine())
	fmt.Println("Tell me a question that distinguishes between")
	fmt.Printf(" %s and\n", *cell.Animal)
	fmt.Printf(" %s\n", animal)

	question := readLine()

	yes := askQuestion("What is the answer for", fmt.Sprintf(" %s?", animal))

	cell.AddNewAnimal(animal, question, yes)
}

func main() {
	thinkAnAnimal := func() {
		fmt.Println("Think of an animal.")
		fmt.Println("Press enter to continue.")
		_ = readLine()
	}

	mem := memory.InitAnimals(memoryFile)
	cell := mem
	thinkAnAnimal()
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
					thinkAnAnimal()
					continue
				}
				mem.Save(memoryFile)
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
					thinkAnAnimal()
					continue
				}
				mem.Save(memoryFile)
				break
			}
			cell = cell.No
		}
	}
}
