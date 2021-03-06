package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
)

func GetStoredRideExample() *StoredRide {
	return &StoredRide{
		DriverAddress:    alice,
		PassengerAddress: bob,
	}
}

func TestCanGetDriverAndPassengerAccount(t *testing.T) {
	aliceAddress, err1 := sdk.AccAddressFromBech32(alice)
	bobAddress, err2 := sdk.AccAddressFromBech32(bob)
	driver, err3 := GetStoredRideExample().GetDriverSdkAddress()
	passenger, err4 := GetStoredRideExample().GetPassengerSdkAddress()
	require.Equal(t, aliceAddress, driver)
	require.Equal(t, bobAddress, passenger)
	require.Nil(t, err1)
	require.Nil(t, err2)
	require.Nil(t, err3)
	require.Nil(t, err4)
}
