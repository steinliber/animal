package behavior

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBehaviorNotFoundBehavior(t *testing.T) {
	t.Parallel()
	animalBehavior := AnimalBehavior{
		Eat:     "eat",
		Speak:   "speak",
		Move:    "move",
		IsShout: false,
	}
	var output bytes.Buffer

	err := animalBehavior.Run(&output, "run")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Animal not have behavior run!\n")
	printResult, _ := ioutil.ReadAll(&output)
	assert.Empty(t, printResult)
}

func TestBehaviorNormal(t *testing.T) {
	t.Parallel()
	animalBehavior := AnimalBehavior{
		Eat:     "eat",
		Speak:   "speak",
		Move:    "move",
		IsShout: false,
	}
	var output bytes.Buffer

	err := animalBehavior.Run(&output, "speak")

	assert.Nil(t, err)
	printResult, _ := ioutil.ReadAll(&output)
	assert.Equal(t, string(printResult), "speak\n")

	animalBehavior.IsShout = true
	err = animalBehavior.Run(&output, "speak")

	assert.Nil(t, err)
	printResult, _ = ioutil.ReadAll(&output)
	assert.Equal(t, string(printResult), "SPEAK\n")

}
