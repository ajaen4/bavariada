package db

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Match struct {
	Status          string           `json:"status"`
	HomeCode        string           `json:"home_code"`
	HomeLogo        string           `json:"home_logo_url"`
	HomeGoals       pgtype.Int4      `json:"home_goals"`
	AwayCode        string           `json:"away_code"`
	AwayLogo        string           `json:"away_logo_url"`
	AwayGoals       pgtype.Int4      `json:"away_goals"`
	KickoffDatetime pgtype.Timestamp `json:"kickoff_datetime"`
	HomePercent     pgtype.Int4      `json:"home_percent"`
	DrawPercent     pgtype.Int4      `json:"draw_percent"`
	AwayPercent     pgtype.Int4      `json:"away_percent"`
}

func GetMatches(matchday int32) ([]Match, error) {
	db := New()

	rows, err := db.Query(
		fmt.Sprintf(
			`SELECT
				status,
				home.code as home_code,
				home.logo_url as home_logo_url,
				home_goals,
				away.code as away_code,
				away.logo_url as away_logo_url,
				away_goals,
				kickoff_datetime,
				stats.home_percent,
				stats.draw_percent,
				stats.away_percent
			FROM bavariada.matches as matches
			LEFT JOIN bavariada.teams home ON matches.home_team_id = home.id
			LEFT JOIN bavariada.teams away ON matches.away_team_id = away.id
			LEFT JOIN bavariada.predictions_stats stats ON
				matches.season = stats.season
				AND matches.matchday = stats.matchday
				AND matches.match_num = stats.match_num
			WHERE matches.matchday = %d
			ORDER BY matches.match_num;`,
			matchday,
		),
	)
	if err != nil {
		return []Match{}, err
	}
	defer rows.Close()

	matches := []Match{}
	for rows.Next() {
		match := Match{}
		err = rows.Scan(
			&match.Status,
			&match.HomeCode,
			&match.HomeLogo,
			&match.HomeGoals,
			&match.AwayCode,
			&match.AwayLogo,
			&match.AwayGoals,
			&match.KickoffDatetime,
			&match.HomePercent,
			&match.DrawPercent,
			&match.AwayPercent,
		)
		if err != nil {
			return []Match{}, err
		}
		matches = append(matches, match)
	}
	return matches, nil
}
