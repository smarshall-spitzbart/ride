// Helper methods involving a stored ride, that're not auto generated.
package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/types"
)

func (storedRide StoredRide) GetDriverAddress() (driver sdk.AccAddress, err error) {
	if !storedRide.HasAssignedDriver() {
		return nil, errors.Wrapf(err, ErrNoAssignedDriver.Error(), storedRide.Driver)
	}
	driver, err = sdk.AccAddressFromBech32(storedRide.Driver)
	return driver, errors.Wrapf(err, ErrInvalidDriver.Error(), storedRide.Driver)
}

func (storedRide StoredRide) HasAssignedDriver() bool {
	return storedRide.Driver != ""
}

func (storedRide StoredRide) IsFinished() bool {
	return storedRide.FinishTime != ""
}

func (storedRide StoredRide) GetPassengerAddress() (passenger sdk.AccAddress, err error) {
	passenger, err = sdk.AccAddressFromBech32(storedRide.Passenger)
	return passenger, errors.Wrapf(err, ErrInvalidPassenger.Error(), storedRide.Passenger)
}

func (storedGame *StoredRide) GetAcceptanceTimeFormatted() (acceptance time.Time, err error) {
	time, err := time.Parse(types.TimeFormat, storedGame.AcceptanceTime)
	return time, sdkerrors.Wrapf(err, ErrCantParseTime.Error(), storedGame.AcceptanceTime)
}

func (storedGame *StoredRide) GetFinishTimeFormatted() (finished time.Time, err error) {
	time, err := time.Parse(types.TimeFormat, storedGame.FinishTime)
	return time, sdkerrors.Wrapf(err, ErrCantParseTime.Error(), storedGame.FinishTime)
}

func TimeToString(time time.Time) string {
	return time.UTC().Format(types.TimeFormat)
}
