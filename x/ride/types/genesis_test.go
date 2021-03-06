package types_test

import (
	"testing"

	"github.com/smarshall-spitzbart/ride/x/ride/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				NextRide: &types.NextRide{
					IdValue: 78,
				},
				StoredRideList: []types.StoredRide{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				RatingStructList: []types.RatingStruct{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated storedRide",
			genState: &types.GenesisState{
				StoredRideList: []types.StoredRide{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated ratingStruct",
			genState: &types.GenesisState{
				RatingStructList: []types.RatingStruct{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestDefaultGenesisIsCorrect(t *testing.T) {
	require.EqualValues(t,
		&types.GenesisState{
			StoredRideList:   []types.StoredRide{},
			NextRide:         &types.NextRide{uint64(1), "-1", "-1"},
			RatingStructList: []types.RatingStruct{},
		},
		types.DefaultGenesis())
}
