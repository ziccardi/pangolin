package memory_test

import (
	"github.com/stretchr/testify/assert"
	"pangolini/memory"
	"testing"
)

const (
	YES = "YES"
	NO  = "NO"
)

func initMemory() *memory.Cell {
	whaleAnswer := memory.NewMemoryCell(memory.WithAnimal("a whale"))
	antAnswer := memory.NewMemoryCell(memory.WithAnimal("an ant"))
	pangolinAnswer := memory.NewMemoryCell(memory.WithAnimal("a pangolin"))
	blancMange := memory.NewMemoryCell(memory.WithAnimal("a blancmange"))

	eatAntsQuestion := memory.NewMemoryCell(
		memory.WithQuestion("Does it eat ants?"),
		memory.WithYes(pangolinAnswer),
		memory.WithNo(antAnswer),
	)

	haveScaleArmourQuestion := memory.NewMemoryCell(
		memory.WithQuestion("Is it scaly?"),
		memory.WithYes(eatAntsQuestion),
		memory.WithNo(blancMange),
	)

	liveInTheWaterQuestion := memory.NewMemoryCell(
		memory.WithQuestion("Does it live in the sea?"),
		memory.WithYes(whaleAnswer),
		memory.WithNo(haveScaleArmourQuestion),
	)

	return liveInTheWaterQuestion
}

func TestMemoryTree(t *testing.T) {
	type args struct {
		answers []string
	}

	tests := []struct {
		name          string
		args          args
		wantQuestions []string
		wantAnimal    string
	}{
		{
			name: "Find whale",
			args: args{
				answers: []string{YES},
			},
			wantQuestions: []string{"Does it live in the sea?"},
			wantAnimal:    "a whale",
		},
		{
			name: "Find pangolin",
			args: args{
				answers: []string{NO, YES, YES},
			},
			wantQuestions: []string{"Does it live in the sea?", "Is it scaly?", "Does it eat ants?"},
			wantAnimal:    "a pangolin",
		},
		{
			name: "Find blancmange",
			args: args{
				answers: []string{NO, NO},
			},
			wantQuestions: []string{"Does it live in the sea?", "Is it scaly?"},
			wantAnimal:    "a blancmange",
		},
		{
			name: "Find ant",
			args: args{
				answers: []string{NO, YES, NO},
			},
			wantQuestions: []string{"Does it live in the sea?", "Is it scaly?", "Does it eat ants?"},
			wantAnimal:    "an ant",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mem := initMemory()
			var questions []string

			for _, answer := range tt.args.answers {
				if mem.Question != nil {
					questions = append(questions, *mem.Question)
				}
				if answer == "YES" {
					mem = mem.Yes
				} else {
					mem = mem.No
				}
			}

			assert.Equal(t, tt.wantQuestions, questions)
			assert.Equal(t, tt.wantAnimal, *mem.Animal)
		})
	}
}
