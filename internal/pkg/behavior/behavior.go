package behavior

import (
	"fmt"
	"io"
	"strings"
)

//AnimalBehavior respensent animal behavior config in csv
//and isShout is a env for control speak string to upper
type AnimalBehavior struct {
	Eat     string
	Move    string
	Speak   string
	IsShout bool
}

//Run exuecute behavior for animal type
func (animalBehavior *AnimalBehavior) Run(w io.Writer, behaviorName string) error {
	behaviorStr := ""
	isShout := false
	switch behaviorName {
	case "eat":
		behaviorStr = animalBehavior.Eat
	case "move":
		behaviorStr = animalBehavior.Move
	case "speak":
		behaviorStr = animalBehavior.Speak
		isShout = animalBehavior.IsShout
	default:
		return fmt.Errorf("Animal not have behavior %s!\n", behaviorName)
	}
	if isShout {
		behaviorStr = strings.ToUpper(behaviorStr)
	}
	fmt.Fprintln(w, behaviorStr)
	return nil
}
