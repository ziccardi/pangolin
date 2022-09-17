package main

import (
	"fmt"
	"pangolini/memory"
	"pangolini/utils"
)

const memoryFile = "pangolins.dat"

func addNewQuestion(cell memory.Cell) {
	fmt.Println("What is it then?")
	animal := string(utils.ReadAnswer())
	fmt.Println("Tell me a question that distinguishes between")
	fmt.Printf(" %s and\n", cell.GetData())
	fmt.Printf(" %s\n", animal)

	question := string(utils.ReadAnswer())

	answer := utils.AskQuestion("What is the answer for", fmt.Sprintf(" %s?", animal))

	cell.AddNewAnimal(animal, question, answer)
}

func main() {
	thinkAnAnimal := func() {
		fmt.Println("Think of an animal.")
		fmt.Println("Press enter to continue.")
		_ = utils.ReadAnswer()
	}

	mem := memory.InitAnimals(memoryFile)
	cell := mem
	thinkAnAnimal()
	for true {

		var answer utils.Answer

		if cell.IsLeaf() { // If it is a leaf, means we reached an answer.
			answer = utils.AskQuestion("Are you thinking of", fmt.Sprintf(" %s?", cell.GetData()))
		} else {
			answer = utils.AskQuestion("", cell.GetData())
		}

		if cell.IsLeaf() {
			if answer.IsYes() {
				fmt.Println("I thought as much")
			} else {
				addNewQuestion(cell)
				fmt.Println("That fooled me.")
			}
			if answer := utils.AskQuestion("", "Do you want another go?"); answer.IsYes() {
				cell = mem
				thinkAnAnimal()
				continue
			}
			mem.Save(memoryFile)
			break
		}
		cell = cell.Next(answer)
	}
}
