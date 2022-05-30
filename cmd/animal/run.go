package main

import (
	"fmt"

	"animal/internal/animalloader"
	"animal/internal/behaviorloader"

	"github.com/spf13/cobra"
)

func runAnimalBeHaviorFunc(cmd *cobra.Command, args []string) {
	animalName := args[0]
	animalType, err := animalloader.GetAnimalType(animalName)
	if err != nil {
		fmt.Println(err)
		return
	}
	animalBehaviorName := args[1]
	err = behaviorloader.RunBehavior(
		animalType, animalBehaviorName,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
