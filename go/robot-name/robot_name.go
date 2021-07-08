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

func (robot *Robot) Name() (string, error) {
	robot.name = GenerateName()

	return robot.name, nil
}

func (robot *Robot) Reset() (bool, error) {
	robot.name = GenerateName()

	return true, nil
}

func GenerateName() string {
	name := GeneratePrefix() + GenerateSuffix()
	for NameAlreadyUsed(name) {
		name = GeneratePrefix() + GenerateSuffix()
	}

	return name
}

// Check if a robot name has already been used.
func NameAlreadyUsed(name string) bool {
	_, seen := robotNames[name]

	return seen
}

func GeneratePrefix() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	char1 := letters[r.Intn(26)]
	char2 := letters[r.Intn(26)]

	return char1 + char2
}

func GenerateSuffix() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num1 := r.Intn(10)
	num2 := r.Intn(10)
	num3 := r.Intn(10)

	return fmt.Sprintf("%v%v%v", num1, num2, num3)
}
