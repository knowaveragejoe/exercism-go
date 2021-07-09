package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

type Team struct {
	name    string
	matches int
	wins    int
	losses  int
	draws   int
	points  int
}

type Tournament struct {
	teams map[string]*Team
}

var (
	alaska     = Team{name: "Allegoric Alaskans"}
	badgers    = Team{name: "Blithering Badgers"}
	california = Team{name: "Courageous Californians"}
	donkeys    = Team{name: "Devastating Donkeys"}

	tournament = Tournament{
		teams: map[string]*Team{
			"Allegoric Alaskians":     &alaska,
			"Blithering Badgers":      &badgers,
			"Courageous Californians": &california,
			"Devastating Donkeys":     &donkeys,
		},
	}

	outcomes = map[string]int{
		"win":  3,
		"loss": 0,
		"draw": 1,
	}
)

const (
	outputHeader = "Team                           | MP |  W |  D |  L |  P\n"
)

func Tally(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		// Get line
		line := scanner.Text()

		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		// Attempt to split on semicolon
		values := strings.Split(line, ";")

		// Every line should contain 3 values after splitting on semicolons.
		if len(values) != 3 {
			fmt.Println(line)
			return errors.New("Invalid match format: " + fmt.Sprintf("%v", values))
		}

		team1 := values[0]
		team2 := values[1]
		outcome := values[2]

		// Update our team scores
		err := UpdateScores(team1, team2, outcome)
		if err != nil {
			return err
		}
	}

	// Write output
	err := WriteScores(w)
	if err != nil {
		return err
	}

	return nil
}

func UpdateScores(team1 string, team2 string, outcome string) error {
	if _, ok := tournament.teams[team1]; !ok {
		return errors.New("Team not found: " + team1)
	}

	if _, ok := tournament.teams[team2]; !ok {
		return errors.New("Team not found: " + team2)
	}

	if _, ok := outcomes[outcome]; !ok {
		return errors.New("Invalid match outcome: " + outcome)
	}

	switch outcome {
	case "win":
		tournament.teams[team1].wins += 1
		tournament.teams[team2].losses += 1

		tournament.teams[team1].points = 3
	case "loss":
		tournament.teams[team1].losses += 1
		tournament.teams[team2].wins += 1

		tournament.teams[team2].points = 3
	case "draw":
		tournament.teams[team1].draws += 1
		tournament.teams[team2].draws += 1

		tournament.teams[team1].points = 1
		tournament.teams[team2].points = 1
	}

	tournament.teams[team1].matches += 1
	tournament.teams[team2].matches += 1

	return nil
}

func WriteScores(writer io.Writer) error {
	fmt.Fprint(writer, outputHeader)

	for _, team := range tournament.teams {
		// offset := 32 - len(team.name)
		line := fmt.Sprintf("%v | %v |  %v |  %v |  %v |  %v\n",
			// offset,
			team.name,
			team.matches,
			team.wins,
			team.draws,
			team.losses,
			team.points,
		)

		fmt.Fprint(writer, line)
	}

	return nil
}
