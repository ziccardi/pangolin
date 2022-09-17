package memory

import (
	"encoding/json"
	"os"
	"pangolini/utils"
)

type Cell interface {
	IsLeaf() bool
	GetData() string
	AddNewAnimal(animal string, question string, answer utils.Answer)
	Save(file string)
	Next(answer utils.Answer) Cell
}

var _ Cell = &cellImpl{}

type cellImpl struct {
	NoBranch  *cellImpl
	YesBranch *cellImpl
	Data      string
}

func (cell *cellImpl) Next(answer utils.Answer) Cell {
	if answer.IsYes() {
		return cell.YesBranch
	} else {
		return cell.NoBranch
	}
}

func (cell *cellImpl) IsLeaf() bool {
	return cell.YesBranch == nil && cell.NoBranch == nil
}

func (cell *cellImpl) GetData() string {
	return cell.Data
}

func (cell *cellImpl) Save(file string) {
	b, _ := json.Marshal(cell)
	_ = os.WriteFile(file, b, 0644)
}

func (cell *cellImpl) AddNewAnimal(animal string, question string, answer utils.Answer) {
	newAnimal := cellImpl{Data: animal}
	oldAnimal := *cell

	if answer.IsYes() {
		updateMemoryCell(
			cell,
			WithQuestion(question),
			WithYes(&newAnimal),
			WithNo(&oldAnimal),
		)
	} else {
		updateMemoryCell(
			cell,
			WithQuestion(question),
			WithYes(&oldAnimal),
			WithNo(&newAnimal),
		)
	}
}

type CellOption func(*cellImpl)

func NewMemoryCell(options ...CellOption) Cell {
	cell := &cellImpl{}
	for _, o := range options {
		o(cell)
	}
	return cell
}

func updateMemoryCell(cell *cellImpl, options ...CellOption) {
	for _, o := range options {
		o(cell)
	}
}

func WithAnimal(animal string) CellOption {
	return func(cell *cellImpl) {
		cell.Data = animal
	}
}

func WithQuestion(question string) CellOption {
	return func(cell *cellImpl) {
		cell.Data = question
	}
}

func WithYes(branch Cell) CellOption {
	return func(cell *cellImpl) {
		cell.YesBranch = branch.(*cellImpl)
	}
}

func WithNo(branch Cell) CellOption {
	return func(cell *cellImpl) {
		cell.NoBranch = branch.(*cellImpl)
	}
}
