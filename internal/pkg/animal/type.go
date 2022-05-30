package animal

import "fmt"

//AnimalTypeInfo respensent animal name and type config in yaml file
type AnimalTypeInfo struct {
	Animals []*struct {
		Name string `yaml:"name"`
		Type string `yaml:"type"`
	} `yaml:"animals"`
}

//GetType get animal type from animal name
func (typeConfig *AnimalTypeInfo) GetType(animalName string) (string, error) {
	for _, animal := range typeConfig.Animals {
		//use first match animal type if matched animal name
		if animal.Name == animalName {
			return animal.Type, nil
		}
	}
	return "", fmt.Errorf("Animal \"%s\" doesn't exist!", animalName)
}
