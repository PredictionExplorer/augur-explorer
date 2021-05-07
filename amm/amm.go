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
