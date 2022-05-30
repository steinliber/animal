package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFileDir = "hack"
	long          = `animal is a commandline tool to get animal behavior from yaml config
- Allow the user to query information about a list of animals defined in a YAML file.
- Each animal has a name(unique) and can only be either a cow, bird, or snake.
- With each command, the user can request information about an animal.`
	short   = "animal is a commandline tool to get animal behavior from yaml config"
	example = `$ animal viper eat
mice
$ animal tweety speak
peep
$ animal gallardo move
walk
$ animal mickey move
Animal "mickey" doesn't exist!`
)

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AutomaticEnv()
	viper.BindEnv("shout", "SHOUT")
	viper.SetDefault("animal_type_config_file_path", fmt.Sprintf(
		"%s/%s", configFileDir, "config.yaml",
	))
	viper.SetDefault("animal_behavior_config_file_path", fmt.Sprintf(
		"%s/%s", configFileDir, "animal_behavior.csv",
	))
}

func main() {
	rootCMD := NewRootCmd()
	err := rootCMD.Execute()
	if err != nil {
		log.Fatalf("Run animal Command Failed: %+v", err)
	}
}

func NewRootCmd() *cobra.Command {
	rootCMD := &cobra.Command{
		Use:     "animal",
		Short:   short,
		Long:    long,
		Example: example,
		Args:    cobra.ExactArgs(2),
		Run:     runAnimalBeHaviorFunc,
	}
	return rootCMD
}
