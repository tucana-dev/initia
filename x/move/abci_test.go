package move_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	abci "github.com/cometbft/cometbft/abci/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	"github.com/initia-labs/initia/x/move/types"
	vmtypes "github.com/initia-labs/initiavm/types"
)

func Test_BeginBlocker(t *testing.T) {
	app := createApp(t)

	// initialize staking for secondBondDenom
	header := tmproto.Header{Height: app.LastBlockHeight() + 1}
	app.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := app.BaseApp.NewContext(false, header)
	err := app.MoveKeeper.InitializeStaking(ctx, secondBondDenom)
	require.NoError(t, err)

	// fund addr2
	app.BankKeeper.SendCoins(ctx, types.StdAddr, addr2, sdk.NewCoins(secondBondCoin))

	app.EndBlock(abci.RequestEndBlock{})
	app.Commit()

	// delegate coins via move staking module
	denomLP := "ulp"
	metadataLP := types.NamedObjectAddress(vmtypes.StdAddress, denomLP)

	valAddrArg, err := vmtypes.SerializeString(sdk.ValAddress(addr1).String())
	require.NoError(t, err)

	amountArg, err := vmtypes.SerializeUint64(secondBondCoin.Amount.Uint64())
	require.NoError(t, err)

	delegateMsg := types.MsgExecute{
		Sender:        addr2.String(),
		ModuleAddress: types.StdAddr.String(),
		ModuleName:    types.MoveModuleNameStaking,
		FunctionName:  types.FunctionNameStakingDelegate,
		TypeArgs:      []string{},
		Args:          [][]byte{metadataLP[:], valAddrArg, amountArg},
	}

	err = executeMsgs(t, app, []sdk.Msg{&delegateMsg}, []uint64{1}, []uint64{0}, priv2)
	require.NoError(t, err)

	// check balance
	checkBalance(t, app, types.MoveStakingModuleAddress, sdk.Coins{})

	// new block
	header = tmproto.Header{Height: app.LastBlockHeight() + 1}
	app.BeginBlock(abci.RequestBeginBlock{Header: header})

	// generate rewards
	ctx = app.BaseApp.NewContext(false, header)
	validator := app.StakingKeeper.Validator(ctx, sdk.ValAddress(addr1))

	rewardCoins := sdk.NewCoins(sdk.NewInt64Coin(bondDenom, 1_000_000))
	delegatorRewardCoins := sdk.NewCoins(sdk.NewInt64Coin(bondDenom, 500_000))

	err = app.BankKeeper.MintCoins(ctx, authtypes.Minter, rewardCoins)
	require.NoError(t, err)
	err = app.BankKeeper.SendCoinsFromModuleToModule(ctx, authtypes.Minter, distrtypes.ModuleName, rewardCoins)
	require.NoError(t, err)
	app.DistrKeeper.AllocateTokensToValidatorPool(
		ctx,
		validator,
		secondBondDenom,
		sdk.NewDecCoinsFromCoins(rewardCoins...))
	app.EndBlock(abci.RequestEndBlock{})
	app.Commit()

	// withdraw rewards to move module
	header = tmproto.Header{Height: app.LastBlockHeight() + 1}
	app.BeginBlock(abci.RequestBeginBlock{Header: header})
	app.EndBlock(abci.RequestEndBlock{})
	app.Commit()

	// undelegate coins
	undelegateMsg := types.MsgExecute{
		Sender:        addr2.String(),
		ModuleAddress: types.StdAddr.String(),
		ModuleName:    types.MoveModuleNameStaking,
		FunctionName:  types.FunctionNameStakingUndelegate,
		TypeArgs:      []string{},
		Args:          [][]byte{metadataLP[:], valAddrArg, amountArg},
	}

	err = executeMsgs(t, app, []sdk.Msg{&undelegateMsg}, []uint64{1}, []uint64{1}, priv2)
	require.NoError(t, err)

	// half rewards and undelegated coins
	checkBalance(t, app, addr2, genCoins.Add(delegatorRewardCoins...))
}
