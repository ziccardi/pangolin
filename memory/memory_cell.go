package memory

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
		cell.Yes = branch
	}
}
