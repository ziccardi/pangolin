package memory

import (
	"encoding/json"
	"os"
)

// InitAnimals - inits the bot memory with a few initial animals
func InitAnimals() *MemoryCell {
	d, err := os.ReadFile("./pangolins.dat")

	if err == nil {
		cell := &MemoryCell{}
		err := json.Unmarshal(d, cell)
		if err == nil {
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
