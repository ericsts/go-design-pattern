package flyweight

import "time"

// Team ...
type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

const (
	// TeamA ...
	TeamA = iota
	// TeamB ...
	TeamB
)

// Player ...
type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

// HistoricalData ...
type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

// Match ...
type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

// TeamFlyweightFactory ...
type TeamFlyweightFactory struct {
	createdTeams map[int]*Team
}

// NewTeamFactory ...
func NewTeamFactory() TeamFlyweightFactory {
	return TeamFlyweightFactory{
		createdTeams: make(map[int]*Team),
	}
}

// GetTeam ...
func (t *TeamFlyweightFactory) GetTeam(teamID int) *Team {
	// return nil
	if t.createdTeams[teamID] != nil {
		return t.createdTeams[teamID]
	}
	team := getTeamFactory(teamID)
	t.createdTeams[teamID] = &team
	return t.createdTeams[teamID]
}

// GetNumberOfObjects ...
func (t *TeamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}

func getTeamFactory(team int) Team {
	switch team {
	case TeamB:
		return Team{
			ID:   2,
			Name: "TEAM_B",
		}
	default:
		return Team{
			ID:   1,
			Name: "TEAM_A",
		}
	}
}
