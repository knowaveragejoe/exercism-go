package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
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

type By func(team1, team2 *Team) bool

// teamSorter joins a By function and a slice of teams to be sorted.
type teamSorter struct {
	teams []*Team
	by    func(team1, team2 *Team) bool // Closure used in the Less method.
}

var (
	alaska     = Team{name: "Allegoric Alaskians"}
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

		tournament.teams[team1].points += outcomes[outcome]
	case "loss":
		tournament.teams[team1].losses += 1
		tournament.teams[team2].wins += 1

		tournament.teams[team2].points += outcomes[outcome]
	case "draw":
		tournament.teams[team1].draws += 1
		tournament.teams[team2].draws += 1

		tournament.teams[team1].points += outcomes[outcome]
		tournament.teams[team2].points += outcomes[outcome]
	}

	tournament.teams[team1].matches += 1
	tournament.teams[team2].matches += 1

	return nil
}

func WriteScores(writer io.Writer) error {
	// Get sorted teams
	sortedTeams, err := SortTeamsByScore()
	if err != nil {
		return err
	}

	// Output header
	fmt.Fprint(writer, outputHeader)

	// Loop over sorted teams and output
	for _, team := range sortedTeams {
		// Calculate offset for padding on team names
		offset := 31 - len(team.name)
		line := fmt.Sprintf("%v", team.name)
		line += strings.Repeat(" ", offset)
		line += fmt.Sprintf("|  %v |  %v |  %v |  %v |  %v\n",
			team.matches,
			team.wins,
			team.draws,
			team.losses,
			team.points,
		)

		// Write line to output buffer
		_, err := fmt.Fprint(writer, line)
		if err != nil {
			return err
		}
	}

	return nil
}

func SortTeamsByScore() ([]*Team, error) {
	// Make a slice of Teams
	teams := make([]*Team, 0, len(tournament.teams))
	for _, team := range tournament.teams {
		teams = append(teams, team)
	}

	// Closure defining how to sort by points
	points := func(team1, team2 *Team) bool {
		return team1.points > team2.points
	}

	// Perform the sort
	By(points).Sort(teams)

	return teams, nil
}

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(teams []*Team) {
	ts := &teamSorter{
		teams: teams,
		by:    by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ts)
}

func (s *teamSorter) Swap(i, j int) {
	s.teams[i], s.teams[j] = s.teams[j], s.teams[i]
}

func (s *teamSorter) Len() int {
	return len(s.teams)
}

func (s *teamSorter) Less(i, j int) bool {
	return s.by(s.teams[i], s.teams[j])
}
