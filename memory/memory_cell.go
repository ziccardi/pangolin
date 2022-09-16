package memory

import (
	"encoding/json"
	"os"
)

type Cell struct {
	No       *Cell
	Yes      *Cell
	Animal   *string
	Question *string
}

func (cell *Cell) Save(file string) {
	b, _ := json.Marshal(cell)
	_ = os.WriteFile(file, b, 0644)
}

func (cell *Cell) AddNewAnimal(animal string, question string, answerYes bool) {
	newAnimal := Cell{Animal: &animal}
	oldAnimal := *cell

	if answerYes {
		updateMemoryCell(
			cell,
			WithAnimal(""),
			WithQuestion(question),
			WithYes(&newAnimal),
			WithNo(&oldAnimal),
		)
	} else {
		updateMemoryCell(
			cell,
			WithAnimal(""),
			WithQuestion(question),
			WithYes(&oldAnimal),
			WithNo(&newAnimal),
		)
	}
}

type CellOption func(*Cell)

func NewMemoryCell(options ...CellOption) *Cell {
	cell := &Cell{}
	for _, o := range options {
		o(cell)
	}
	return cell
}

func updateMemoryCell(cell *Cell, options ...CellOption) {
	for _, o := range options {
		o(cell)
	}
}

func WithAnimal(animal string) CellOption {
	return func(cell *Cell) {
		if animal != "" {
			cell.Animal = &animal
		} else {
			cell.Animal = nil
		}
	}
}

func WithQuestion(question string) CellOption {
	return func(cell *Cell) {
		if question != "" {
			cell.Question = &question
		} else {
			cell.Question = nil
		}
	}
}

func WithYes(branch *Cell) CellOption {
	return func(cell *Cell) {
		cell.Yes = branch
	}
}

func WithNo(branch *Cell) CellOption {
	return func(cell *Cell) {
		cell.Yes = branch
	}
}
