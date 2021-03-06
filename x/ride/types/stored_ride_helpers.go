// Helper methods involving a stored ride, that're not auto generated.
package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/types"
)

func (storedRide StoredRide) GetDriverSdkAddress() (driver sdk.AccAddress, err error) {
	if !storedRide.HasAssignedDriver() {
		return nil, errors.Wrapf(err, ErrNoAssignedDriver.Error(), storedRide.DriverAddress)
	}
	driver, err = sdk.AccAddressFromBech32(storedRide.DriverAddress)
	return driver, errors.Wrapf(err, ErrInvalidDriver.Error(), storedRide.DriverAddress)
}

func (storedRide StoredRide) GetPassengerSdkAddress() (passenger sdk.AccAddress, err error) {
	passenger, err = sdk.AccAddressFromBech32(storedRide.PassengerAddress)
	return passenger, errors.Wrapf(err, ErrInvalidPassenger.Error(), storedRide.PassengerAddress)
}

func (storedRide StoredRide) HasAssignedDriver() bool {
	return storedRide.DriverAddress != ""
}

func (storedRide StoredRide) IsFinished() bool {
	return storedRide.FinishTime != ""
}

func (storedRide StoredRide) HasExpired(ctx sdk.Context) (bool, error) {
	deadlineAsTime, err := storedRide.GetDeadlineFormatted()
	if err != nil {
		return false, err
	}
	return deadlineAsTime.Before(ctx.BlockTime()), nil
}

func (storedRide *StoredRide) GetAcceptanceTimeFormatted() (accepted time.Time, err error) {
	accepted, err = time.Parse(types.TimeFormat, storedRide.AcceptanceTime)
	return accepted, sdkerrors.Wrapf(err, ErrCantParseTime.Error(), storedRide.AcceptanceTime)
}

func (storedRide *StoredRide) GetFinishTimeFormatted() (finished time.Time, err error) {
	finished, err = time.Parse(types.TimeFormat, storedRide.FinishTime)
	return finished, sdkerrors.Wrapf(err, ErrCantParseTime.Error(), storedRide.FinishTime)
}

func (storedRide *StoredRide) GetDeadlineFormatted() (deadline time.Time, err error) {
	deadline, err = time.Parse(types.TimeFormat, storedRide.Deadline)
	return deadline, sdkerrors.Wrapf(err, ErrCantParseTime.Error(), storedRide.FinishTime)
}

func (storedRide StoredRide) GetMutualStakeInCoin() (stake sdk.Coin) {
	return sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(storedRide.MutualStake)))
}

func TimeToString(time time.Time) string {
	return time.UTC().Format(types.TimeFormat)
}
