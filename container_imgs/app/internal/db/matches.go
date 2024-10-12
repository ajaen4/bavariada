package db

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Match struct {
	HomeName        string           `json:"home_name"`
	AwayName        string           `json:"away_name"`
	KickoffDatetime pgtype.Timestamp `json:"kickoff_datetime"`
}

func GetMatches(matchday int32) ([]Match, error) {
	db := New()

	rows, err := db.Query(
		fmt.Sprintf(
			`SELECT home.name as home_name, away.name as away_name, kickoff_datetime
			FROM bavariada.matches as matches
			LEFT JOIN bavariada.teams home ON matches.home_team_id = home.id
			LEFT JOIN bavariada.teams away ON matches.away_team_id = away.id
			WHERE matchday = %d
			ORDER BY match_num;`,
			matchday,
		),
	)
	if err != nil {
		return []Match{}, err
	}

	matches := []Match{}
	for rows.Next() {
		match := Match{}
		err = rows.Scan(&match.HomeName, &match.AwayName, &match.KickoffDatetime)
		if err != nil {
			return []Match{}, err
		}
		matches = append(matches, match)
	}
	return matches, nil
}
