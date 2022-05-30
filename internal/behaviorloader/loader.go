package behaviorloader

import (
	"animal/internal/pkg/behavior"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	nameLoc = iota
	eatLoc
	moveLoc
	speakLoc
)

//RunBehavior get animal type and animal behavior to execute behavior base on behavior config csv file
func RunBehavior(animalType, behaviorName string) error {
	behaviorInfo, err := loadAnimalBehavior(animalType)
	if err != nil {
		return err
	}
	return behaviorInfo.Run(os.Stdout, behaviorName)
}

func loadAnimalBehavior(animalType string) (*behavior.AnimalBehavior, error) {
	file, err := os.Open(viper.GetString("animal_behavior_config_file_path"))
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		behaviorInfo := strings.Split(scanner.Text(), ",")
		if len(behaviorInfo) != 4 || animalType != behaviorInfo[nameLoc] {
			continue
		}
		isShout := viper.GetBool("shout")
		return &behavior.AnimalBehavior{
			Eat:     behaviorInfo[eatLoc],
			Move:    behaviorInfo[moveLoc],
			Speak:   behaviorInfo[speakLoc],
			IsShout: isShout,
		}, nil
	}
	return nil, fmt.Errorf("Animal type %s not has behavior config!", animalType)
}
