package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var (
	robotNames map[string]int
	letters    = [27]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
)

// Retrieve the Robot's name if already set, otherwise generate, set and return.
func (robot *Robot) Name() (string, error) {
	// If name is already set, return it
	if len(robot.name) == 5 {
		return robot.name, nil
	}
	// Otherwise, generate and set it on the Robot.
	robot.name = GenerateName()

	return robot.name, nil
}

// Reset a robot's name to something unique.
func (robot *Robot) Reset() (bool, error) {
	robot.name = GenerateName()

	return true, nil
}

// Generate a robot name. If already used, continue trying until a unique one is found.
func GenerateName() string {
	// robotNames might need to be initialized
	if len(robotNames) <= 0 {
		robotNames = make(map[string]int)
	}

	name := GeneratePrefix() + GenerateSuffix()
	for NameAlreadyUsed(name) {
		name = GeneratePrefix() + GenerateSuffix()
	}
	
	robotNames[name] = 1

	return name
}

// Check if a robot name has already been used.
func NameAlreadyUsed(name string) bool {
	_, seen := robotNames[name]

	return seen
}

// Generate a random 2-letter prefix.
func GeneratePrefix() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	char1 := letters[r.Intn(26)]
	char2 := letters[r.Intn(26)]

	return char1 + char2
}

// Generate a random 3-digit suffix.
func GenerateSuffix() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num1 := r.Intn(10)
	num2 := r.Intn(10)
	num3 := r.Intn(10)

	return fmt.Sprintf("%v%v%v", num1, num2, num3)
}
