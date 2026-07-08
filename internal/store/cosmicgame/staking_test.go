package cosmicgame

import (
	"strings"
	"testing"
)

// TestStakeActionQueryShapes pins the SQL the two stake-action info queries
// generate for their (only) production table pairs: the CST variant carries
// the four reward columns, the RWalk variant must not, and both join stake
// and unstake rows on action_id with a single bind parameter.
func TestStakeActionQueryShapes(t *testing.T) {
	cst := stakeActionQueryCST("cg_nft_staked_cst", "cg_nft_unstaked_cst")
	rwalk := stakeActionQueryRWalk("cg_nft_staked_rwalk", "cg_nft_unstaked_rwalk")

	for name, q := range map[string]string{"cst": cst, "rwalk": rwalk} {
		if !strings.HasPrefix(q, "SELECT ") {
			t.Errorf("%s query does not start with SELECT: %q", name, q[:20])
		}
		if !strings.HasSuffix(q, "WHERE st.action_id=$1") {
			t.Errorf("%s query does not filter on st.action_id=$1", name)
		}
		if strings.Count(q, "$") != 1 {
			t.Errorf("%s query must bind exactly one parameter, got %d", name, strings.Count(q, "$"))
		}
		for _, join := range []string{
			"LEFT JOIN transaction ts ON ts.id=st.tx_id",
			"LEFT JOIN address sa ON st.staker_aid=sa.address_id",
			"LEFT JOIN transaction tu ON tu.id=u.tx_id",
			"LEFT JOIN address ua ON u.staker_aid=ua.address_id",
		} {
			if !strings.Contains(q, join) {
				t.Errorf("%s query is missing join %q", name, join)
			}
		}
	}

	if !strings.Contains(cst, "FROM cg_nft_staked_cst st") ||
		!strings.Contains(cst, "LEFT JOIN cg_nft_unstaked_cst u ON st.action_id=u.action_id") {
		t.Error("cst query does not reference the CST stake/unstake tables")
	}
	if !strings.Contains(rwalk, "FROM cg_nft_staked_rwalk st") ||
		!strings.Contains(rwalk, "LEFT JOIN cg_nft_unstaked_rwalk u ON st.action_id=u.action_id") {
		t.Error("rwalk query does not reference the RWalk stake/unstake tables")
	}

	for _, rewardCol := range []string{"u.reward,", "u.reward/1e18,", "u.reward_per_tok,", "u.reward_per_tok/1e18,"} {
		if !strings.Contains(cst, rewardCol) {
			t.Errorf("cst query is missing reward column %q", rewardCol)
		}
		if strings.Contains(rwalk, rewardCol) {
			t.Errorf("rwalk query must not select reward column %q", rewardCol)
		}
	}
}
