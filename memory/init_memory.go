package memory

import (
	"encoding/json"
	"os"
)

// InitAnimals - inits the bot memory with a few initial animals
func InitAnimals(file string) Cell {
	d, err := os.ReadFile(file)

	if err == nil {
		cell := &cellImpl{}

		if err := json.Unmarshal(d, cell); err == nil {
			return cell
		}
	}

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
