package animal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnimalTypeInfoNotFound(t *testing.T) {
	t.Parallel()
	animalTypes := AnimalTypeInfo{
		Animals: []*struct {
			Name string `yaml:"name"`
			Type string `yaml:"type"`
		}{
			{
				Name: "murcielago",
				Type: "cow",
			},
			{
				Name: "gallardo",
				Type: "cow",
			},
		},
	}

	result, err := animalTypes.GetType("mickey")
	assert.Empty(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), `Animal "mickey" doesn't exist!`)
}

func TestAnimalTypeInfoNormal(t *testing.T) {
	t.Parallel()
	animalTypes := AnimalTypeInfo{
		Animals: []*struct {
			Name string `yaml:"name"`
			Type string `yaml:"type"`
		}{
			{
				Name: "murcielago",
				Type: "cow",
			},
			{
				Name: "gallardo",
				Type: "cow",
			},
		},
	}

	result, err := animalTypes.GetType("murcielago")
	assert.Nil(t, err)
	assert.Equal(t, result, "cow")

}
