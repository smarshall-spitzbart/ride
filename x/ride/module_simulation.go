package ride

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/smarshall-spitzbart/ride/testutil/sample"
	ridesimulation "github.com/smarshall-spitzbart/ride/x/ride/simulation"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = ridesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRequestRide = "op_weight_msg_request_ride"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRequestRide int = 100

	opWeightMsgAccept = "op_weight_msg_accept"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAccept int = 100

	opWeightMsgFinish = "op_weight_msg_finish"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFinish int = 100

	opWeightMsgRate = "op_weight_msg_rate"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRate int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	rideGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&rideGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRequestRide int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRequestRide, &weightMsgRequestRide, nil,
		func(_ *rand.Rand) {
			weightMsgRequestRide = defaultWeightMsgRequestRide
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestRide,
		ridesimulation.SimulateMsgRequestRide(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAccept int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAccept, &weightMsgAccept, nil,
		func(_ *rand.Rand) {
			weightMsgAccept = defaultWeightMsgAccept
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAccept,
		ridesimulation.SimulateMsgAccept(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgFinish int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgFinish, &weightMsgFinish, nil,
		func(_ *rand.Rand) {
			weightMsgFinish = defaultWeightMsgFinish
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgFinish,
		ridesimulation.SimulateMsgFinish(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRate, &weightMsgRate, nil,
		func(_ *rand.Rand) {
			weightMsgRate = defaultWeightMsgRate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRate,
		ridesimulation.SimulateMsgRate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
