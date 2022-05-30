package animalloader

import (
	"animal/internal/pkg/animal"
	"io/ioutil"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

//LoadAnimalTypeMap receive animal name retrurn correspond animal type
func GetAnimalType(animalName string) (string, error) {
	animalTypeInfo, err := parseYamlFile(viper.GetString("animal_type_config_file_path"))
	if err != nil {
		return "", err
	}
	return animalTypeInfo.GetType(animalName)
}

func parseYamlFile(filePath string) (*animal.AnimalTypeInfo, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	animalTypeInfo := new(animal.AnimalTypeInfo)
	err = yaml.Unmarshal(fileBytes, animalTypeInfo)
	return animalTypeInfo, err
}
