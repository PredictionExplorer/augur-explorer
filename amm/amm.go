package amm

import (
	"os"
	"fmt"
	"io/ioutil"

	"encoding/json"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)

func Load_amm_constants(path string) p.AMM_Constants {

	var constants p.AMM_Constants

	var fname string
	fname = path + "/sports.json"
	f_sports_categories, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Error opening %v : %v\n",fname,err)
		os.Exit(1)
	}
	defer f_sports_categories.Close()

	data,_ := ioutil.ReadAll(f_sports_categories)
	var cat_entries p.AMM_CatEntries
	err = json.Unmarshal([]byte(data),&cat_entries)
	if err != nil {
		fmt.Printf("Error parsing json (%v) : %v\n",fname,err)
		os.Exit(1)
	}
	constants.Categories = cat_entries

	fname = path + "/teams.json"
	f_sports_teams, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Error opening %v : %v\n",fname,err)
		os.Exit(1)
	}
	defer f_sports_categories.Close()

	data,_ = ioutil.ReadAll(f_sports_teams)
	var team_entries p.AMM_TeamEntries
	err = json.Unmarshal([]byte(data),&team_entries)
	if err != nil {
		fmt.Printf("Error parsing json (%v) : %v\n",fname,err)
		os.Exit(1)
	}
	constants.Teams = team_entries

	return constants
}
/*
func get_outcome_name( outcome_id int64,sport_id: int64,home_team: string,
  awayTeam: string,
  sportsMarketType: number,
  line: string
) => {
  const marketOutcome = getMarketOutcome(sportId, sportsMarketType, outcomeId);
  // create outcome name using market type and line
  if (outcomeId === NO_CONTEST_OUTCOME_ID) return marketOutcome;

  if (sportsMarketType === SPORTS_MARKET_TYPE.MONEY_LINE) {
    return populateHomeAway(marketOutcome, homeTeam, awayTeam);
  }

  if (sportsMarketType === SPORTS_MARKET_TYPE.SPREAD) {
    // spread
    let pLine = Number(line) > 0 ? `+${line}` : line;
    if (outcomeId === AWAY_TEAM_OUTCOME) {
      const newLine = Number(line) * -1; // invert for away team
      pLine = newLine > 0 ? `+${newLine}` : `${newLine}`;
    }
    const outcomes = populateHomeAway(marketOutcome, homeTeam, awayTeam).replace(NAMING_LINE.SPREAD_LINE, pLine);
    return outcomes;
  }

  if (sportsMarketType === SPORTS_MARKET_TYPE.OVER_UNDER) {
    // over/under
    return marketOutcome.replace(NAMING_LINE.OVER_UNDER_LINE, line);
  }

  return `Outcome ${outcomeId}`;
};
*/
type SportsDataEntry struct {
	Title			string
	Description		string
	Outcomes		[]string
}
type SportsTypes struct {
	Name			string
	Types			map[int64]SportsDataEntry
}
type SportsData struct {
	Entries			map[int64]SportsTypes
}
const (
	SPORTS_MARKET_TYPE_MONEY_LINE = 0
	SPORTS_MARKET_TYPE_SPREAD = 1
	SPORTS_MARKET_TYPE_OVER_UNDER = 2
)
const (
	NAMING_TEAM_HOME_TEAM = "HOME_TEAM"
	NAMING_TEAM_AWAY_TEAM = "AWAY_TEAM"
	NAMING_TEAM_FAV_TEAM = "FAV_TEAM"
	NAMING_TEAM_UNDERDOG_TEAM = "UNDERDOG_TEAM"
)
const (
	NAMING_LINE_SPREAD_LINE = "SPREAD_LINE"
	NAMING_LINE_OVER_UNDER_LINE = "OVER_UNDER_LINE"
)
const NO_CONTEST = "No Contest";
const NO_CONTEST_TIE = "Tie/No Contest";
const HOME_TEAM_OUTCOME = 1;
const AWAY_TEAM_OUTCOME = 2;
var	sportsData = map[int64]SportsTypes{
		2: SportsTypes {
			Name: "NFL",
			Types: map[int64]SportsDataEntry {
					SPORTS_MARKET_TYPE_MONEY_LINE : SportsDataEntry {
						Title: `Which team will win?`,
						Description: NAMING_TEAM_AWAY_TEAM+` vs `+NAMING_TEAM_HOME_TEAM+`?`,
						Outcomes: []string{NO_CONTEST_TIE,NAMING_TEAM_AWAY_TEAM,NAMING_TEAM_HOME_TEAM},
					},
					SPORTS_MARKET_TYPE_SPREAD: SportsDataEntry {
						Title: `Will the `+NAMING_TEAM_FAV_TEAM+` defeat the `+NAMING_TEAM_UNDERDOG_TEAM+` by more than `+NAMING_LINE_SPREAD_LINE+`.5 points?`,
						Description: ``,
						Outcomes: []string{NO_CONTEST,NAMING_TEAM_AWAY_TEAM+" "+NAMING_LINE_SPREAD_LINE+".5",NAMING_TEAM_HOME_TEAM+" "+NAMING_LINE_SPREAD_LINE+".5"},
					},
					SPORTS_MARKET_TYPE.OVER_UNDER: SportDataEntry {
						Title: "Will there be over "+NAMING_LINE_OVER_UNDER_LINE+".5 total points scored?",
						Description: NAMING_TEAM_AWAY_TEAM+" vs "+ NAMING_TEAM_HOME_TEAM,
						Outcomes: []string{NO_CONTEST, "Over "+NAMING_LINE_OVER_UNDER_LINE+".5", "Under "+NAMING_LINE_OVER_UNDER_LINE+".5"},
					},
				},
			},
		},
		3: {
			Name: "MLB",
			Types: map[int64]SportsDataEntry {
				SPORTS_MARKET_TYPE_MONEY_LINE: SportsDataEntry {
					Title: `Which team will win?`,
					Description: NAMING_TEAM_AWAY_TEAM+` vs `+NAMING_TEAM_HOME_TEAM+`?`,
					Outcomes: []string{NO_CONTEST, NAMING_TEAM_AWAY_TEAM+`, `+NAMING_TEAM_HOME_TEAM},
				},
				SPORTS_MARKET_TYPE+SPREAD: SportsDataEntry {
					Title: `Will the `+NAMING_TEAM_FAV_TEAM+` defeat the `+NAMING_TEAM_UNDERDOG_TEAM+` by more than `+NAMING_LINE_SPREAD_LINE+`.5 runs?`,
					Description: ``,
					Outcomes: []string{NO_CONTEST,NAMING_TEAM_AWAY_TEAM+` `+NAMING_LINE_SPREAD_LINE+`.5`,NAMING_TEAM_HOME_TEAM+`+ `+NAMING_LINE_SPREAD_LINE+`.5`},
				},
				SPORTS_MARKET_TYPE.OVER_UNDER: SportsDataEntry {
					Title: `Will there be over `+NAMING_LINE_OVER_UNDER_LINE+`.5 total runs scored?`,
					Description: NAMING_TEAM_AWAY_TEAM+` vs `+NAMING_TEAM.HOME_TEAM,
					Outcomes: []string{NO_CONTEST, `Over `+NAMING_LINE_OVER_UNDER_LINE+`.5`,`Under `+NAMING_LINE_OVER_UNDER_LINE+`.5`},
				},
			},
		},
		4: {
			Name: "NBA",
			Types: map[int64]SportsDataEntry {
				SPORTS_MARKET_TYPE_MONEY_LINE: SportsDataEntry {
					Ttitle: `Which team will win?`,
					Description: NAMING_TEAM.AWAY_TEAM+` vs `+NAMING_TEAM.HOME_TEAM+`?`,
					Outcomes: []string{NO_CONTEST, NAMING_TEAM.AWAY_TEAM+` `+NAMING_TEAM.HOME_TEAM},
				},
				SPORTS_MARKET_TYPE_SPREAD: SportsDataEntry {
					Title: `Will the `+NAMING_TEAM_FAV_TEAM+` defeat the `+NAMING_TEAM_UNDERDOG_TEAM+` by more than `+NAMING_LINE_SPREAD_LINE+`.5 points?`,
					Description: ``,
					Outcomes: []string{NO_CONTEST,NAMING_TEAM_AWAY_TEAM+` `+NAMING_LINE_SPREAD_LINE+`.5`,NAMING_TEAM_HOME_TEAM+` `+NAMING_LINE_SPREAD_LINE+`.5`},
				},
				SPORTS_MARKET_TYPE_OVER_UNDER: {
					Title: `Will there be over `+NAMING_LINE_OVER_UNDER_LINE+`.5 total points scored?`,
					Description: NAMING_TEAM_AWAY_TEAM+` vs `+NAMING_TEAM.HOME_TEAM,
					Outcomes: []string{NO_CONTEST, `Over `+NAMING_LINE_OVER_UNDER_LINE+`.5`, `Under `+NAMING_LINE.OVER_UNDER_LINE+`.5`},
				},
			},
		},
		6: {
			Name: "NHL",
			Types: map[int64]SportsDataEntry {
				SPORTS_MARKET_TYPE_MONEY_LINE: SportsDataEntry {
					Title: `Which team will win?`,
					Description: NAMING_TEAM_AWAY_TEAM+` vs `+NAMING_TEAM_HOME_TEAM+`?`,
					Outcomes: []string{NO_CONTEST, NAMING_TEAM_AWAY_TEAM+`, `+NAMING_TEAM_HOME_TEAM+`]`},
				},
				SPORTS_MARKET_TYPE_SPREAD: SportsDataEntry {
					Title: `Will the `+NAMING_TEAM_FAV_TEAM+` defeat the `+NAMING_TEAM_UNDERDOG_TEAM+` by more than `+NAMING_LINE_SPREAD_LINE+`.5 goals?`,
					Description: ``,
					Outcomes: []string{NO_CONTEST,NAMING_TEAM_AWAY_TEAM+` `+NAMING_LINE_SPREAD_LINE+`.5`,NAMING_TEAM.HOME_TEAM+` `+NAMING_LINE_SPREAD_LINE+`.5`},
				},
				SPORTS_MARKET_TYPE_OVER_UNDER: SportsDataEntry {
					Title: `Will there be over `+NAMING_LINE_OVER_UNDER_LINE+`.5 total goals scored?`,
					Description: NAMING_TEAM.AWAY_TEAM +" vs "+NAMING_TEAM.HOME_TEAM,
					Outcomes: []string{NO_CONTEST, `Over `+NAMING_LINE_OVER_UNDER_LINE+`.5`, `Under `+NAMING_LINE_OVER_UNDER_LINE+`.5`},
				},
			},
		},
		7: {
			Name: "MMA",
			Types: map[int64]SportsDataEntry {
				SPORTS_MARKET_TYPE_MONEY_LINE: SportsDataEntry {
					Title: `Who will win?`,
					Description: NAMING_TEAM_HOME_TEAM+` vs `+NAMING_TEAM_AWAY_TEAM+`?`,
					Outcomes: []string{NO_CONTEST, NAMING_TEAM_HOME_TEAM+`, `+NAMING_TEAM_AWAY_TEAM},
				},
				SPORTS_MARKET_TYPE_SPREAD: SportsDataEntry {
					Title: ``,
					Description: ``,
				},
				SPORTS_MARKET_TYPE.OVER_UNDER: SportsDataEntry {
					Title: `Will fight go the distance?`,
					Description: ``,
					Outcomes: []string{NO_CONTEST, `Yes`, `No`},
				},
			},
		},
}
type ResolutionRules struct {
	Types		map[int64][]string
}
const sportsResolutionRules = map[int64]ResolutionRules{
	2: ResolutionRules {
			Types: map[int64][]string{
				SPORTS_MARKET_TYPE_MONEY_LINE: []string{
					`At least 55 minutes of play must have elapsed for the game to be deemed official. If the game is not played or if less than 55 minutes of play have been completed, the game is not considered
			an official game and the market should resolve as 'No Contest'.`,
					`Overtime counts towards settlement purposes.`,
					`If the game ends in a tie, the market should resolve as 'No Contest'`,
					`If the game is not played, the market should resolve as 'No Contest'.`,
					`Results are determined by their natural conclusion and do not recognize postponed games,
			protests, or overturned decisions.`,
				},
				SPORTS_MARKET_TYPE_SPREAD: []string{
					`At least 55 minutes of play must have elapsed for the game to be deemed official. If the game is
			not played or if less than 55 minutes of play have been completed, the game is not considered
			an official game and the market should resolve as 'No Contest'.`,
					`Overtime counts towards settlement purposes.`,
					`If the game is not played, the market should resolve as 'No Contest'.`,
					`Results are determined by their natural conclusion and do not recognize postponed games,
			protests, or overturned decisions.`,
				},
				SPORTS_MARKET_TYPE_OVER_UNDER: []string{
					`At least 55 minutes of play must have elapsed for the game to be deemed official. If the game is
			not played or if less than 55 minutes of play have been completed, the game is not considered
			an official game and the market should resolve as 'No Contest'.`,
					`Overtime count towards settlement purposes.`,
					`If the game is not played, the market should resolve as 'No Contest'.`,
					`Results are determined by their natural conclusion and do not recognize postponed games,
			protests, or overturned decisions.`,
				},
			},
	},
	3: ResolutionRules {
			Types: map[int64][]string{
				SPORTS_MARKET_TYPE_MONEY_LINE: []string{
					`The results of a game are official after (and, unless otherwise stated, bets shall be settled subject to the completion of) 5 innings of play, or 4.5 innings should the home team be leading at the commencement of the bottom of the 5th innings. Should a game be called, if the result is official in accordance with this rule, the winner will be determined by the score/stats after the last completed inning.`,
					`If the game does not reach the "official” time limit, or ends in a tie, the market should resolve as 'No Contest'.`,
					`If the game is not played, the market should resolve as 'No Contest'.`,
					`Extra innings count towards settlement purposes.`,
					`Results are determined by the natural conclusion and do not recognize postponed games, protests, or overturned decisions.`,
				},
				SPORTS_MARKET_TYPE_SPREAD: []string{
					`The results of a game are official after (and, unless otherwise stated, bets shall be settled subject to the completion of) 5 innings of play, or 4.5 innings should the home team be leading at the commencement of the bottom of the 5th innings. Should a game be called, if the result is official in accordance with this rule, the winner will be determined by the score/stats after the last completed inning.`,
					`If the game does not reach the "official” time limit, or ends in a tie, the market should resolve as 'No Contest'.`,
					`If the game is not played, the market should resolve as 'No Contest'.`,
					`Extra innings count towards settlement purposes.`,
					`Results are determined by their natural conclusion and do not recognize postponed games, protests, or overturned decisions.`,
				},
				SPORTS_MARKET_TYPE.OVER_UNDER: []string{
					`The results of a game are official after (and, unless otherwise stated, bets shall be settled subject to the completion of) 5 innings of play, or 4.5 innings should the home team be leading at the commencement of the bottom of the 5th innings. Should a game be called, if the result is official in accordance with this rule, the winner will be determined by the score/stats after the last completed inning.`,
					`If the game does not reach the "official” time limit, the market should resolve as 'No Contest'.`,
					`If the game is not played, the market should resolve as 'No Contest'.`,
					`Extra innings count towards settlement purposes.`,
					`Results are determined by their natural conclusion and do not recognize postponed games, protests, or overturned decisions.`,
				},
			},
	},
	4: ResolutionRules {
			Types: map[int64][]string{
				SPORTS_MARKET_TYPE_MONEY_LINE: []string{
					`At least 43 minutes of play must have elapsed for the game to be deemed official. If the game is not played or if less than 43 minutes of play have been completed, the game is not considered an official game and the market should resolve as 'No Contest'.`,
					`Overtime count towards settlement purposes.`,
					`If the game is not played, the market should resolve as 'No Contest'.`,
					`Results are determined by their natural conclusion and do not recognize postponed games, protests, or overturned decisions.`,
				},
				SPORTS_MARKET_TYPE_SPREAD: []string{
					`At least 43 minutes of play must have elapsed for the game to be deemed official. If the game is not played or if less than 43 minutes of play have been completed, the game is not considered an official game and the market should resolve as 'No Contest'.`,
					`Overtime count towards settlement purposes.`,
					`If the game is not played, the market should resolve as 'No Contest'.`,
					`Results are determined by their natural conclusion and do not recognize postponed games, protests, or overturned decisions.`,
				},
				SPORTS_MARKET_TYPE_OVER_UNDER: []string{
					`At least 43 minutes of play must have elapsed for the game to be deemed official. If the game is not played or if less than 43 minutes of play have been completed, the game is not considered an official game and the market should resolve as 'No Contest'.`,
					`Overtime count towards settlement purposes.`,
					`If the game is not played, the market should resolve as 'No Contest'.`,
					`Results are determined by their natural conclusion and do not recognize postponed games, protests, or overturned decisions.`,
				},
		},
	},
	6: ResolutionRules {
			Types: {
				SPORTS_MARKET_TYPE_MONEY_LINE: []string{},
				SPORTS_MARKET_TYPE_SPREAD: []string{},
				SPORTS_MARKET_TYPE_OVER_UNDER: []string{},
			},
	},
	7: ResolutionRules {
			Types: {
				SPORTS_MARKET_TYPE_MONEY_LINE: []string{
					`A fight is considered official once the first round begins, regardless of the scheduled or actual duration.`,
					`Market resolves based on the official result immediately following the fight. Later announcements, enquirers, or changes to the official result will not affect market settlement.`,
					`If a fighter is substituted before the fight begins the market should resolve as "Draw/No Contest".`,
					`If a fighter is disqualified during the fight, the opposing fighter should be declared the winner. If both fighters are disqualified the market should resolve as "Draw/No Contest".`,
					`If the fight is cancelled before it starts for any reason, the market should resolve as 'No Contest'.`,
					`A draw can occur when the fight is either stopped before completion or after all rounds are completed and goes to the judges' scorecards for decision. If the match ends in a draw, only the “Draw/No Contest” result should be the winning outcome`,
				},
				SPORTS_MARKET_TYPE_SPREAD: []string{},
				SPORTS_MARKET_TYPE_OVER_UNDER: []string{
					`A fight is considered official once the first round begins, regardless of the scheduled or actual duration.`,
					`Market resolves based on the official result immediately following the fight. Later announcements, enquirers, or changes to the official result will not affect market settlement.`,
					`If a fighter is substituted before the fight begins the market should resolve as "Draw/No Contest".`,
					`If the fight is cancelled before it starts for any reason, the market should resolve as 'No Contest'.`,
					`If the official time is exactly on (equal to) the over/under number the market should resolve as “Over”.`,
					`Markets referring to round/fight duration represents the actual time passed in the round/fight, as applicable, depending on the scheduled round/fight duration. For example Over 2.5 Total Rounds will be settled as “Over” once two and a half minutes or more in the 3rd Round has passed.`,
				},
			},
	},
};
