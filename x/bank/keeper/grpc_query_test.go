package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/stretchr/testify/require"
)

func TestQueryBalance(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := testdata.KeyTestPubAddr()

	_, err := input.BankKeeper.Balance(sdk.WrapSDKContext(ctx), &types.QueryBalanceRequest{})
	require.Error(t, err)

	_, err = input.BankKeeper.Balance(sdk.WrapSDKContext(ctx), &types.QueryBalanceRequest{Address: addr.String()})
	require.Error(t, err)

	testDenom := testDenoms[0]
	req := types.NewQueryBalanceRequest(addr, testDenom)
	res, err := input.BankKeeper.Balance(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.True(t, res.Balance.IsZero())

	origCoins := sdk.NewCoins(sdk.NewCoin(testDenom, sdk.NewInt(50)), sdk.NewCoin(bondDenom, sdk.NewInt(30)))
	acc := input.AccountKeeper.NewAccountWithAddress(ctx, addr)

	input.AccountKeeper.SetAccount(ctx, acc)
	input.Faucet.Fund(ctx, acc.GetAddress(), origCoins...)

	res, err = input.BankKeeper.Balance(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.True(t, res.Balance.IsEqual(sdk.NewCoin(testDenom, sdk.NewInt(50))))
}

func TestQueryAllBalances(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := testdata.KeyTestPubAddr()
	_, err := input.BankKeeper.AllBalances(sdk.WrapSDKContext(ctx), &types.QueryAllBalancesRequest{})
	require.Error(t, err)

	pageReq := &query.PageRequest{
		Key:        nil,
		Limit:      1,
		CountTotal: false,
	}
	req := types.NewQueryAllBalancesRequest(addr, pageReq)
	res, err := input.BankKeeper.AllBalances(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.True(t, res.Balances.IsZero())

	testDenom := testDenoms[0]
	rewardCoin := sdk.NewCoin(testDenom, sdk.NewInt(50))
	bondCoin := sdk.NewCoin(bondDenom, sdk.NewInt(30))

	origCoins := sdk.NewCoins(rewardCoin, bondCoin)
	acc := input.AccountKeeper.NewAccountWithAddress(ctx, addr)

	input.AccountKeeper.SetAccount(ctx, acc)
	input.Faucet.Fund(ctx, acc.GetAddress(), origCoins...)

	res, err = input.BankKeeper.AllBalances(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, res.Balances.Len(), 1)
	require.NotNil(t, res.Pagination.NextKey)

	t.Log("query second page with nextkey")
	pageReq = &query.PageRequest{
		Key:        res.Pagination.NextKey,
		Limit:      1,
		CountTotal: true,
	}
	req = types.NewQueryAllBalancesRequest(addr, pageReq)
	res, err = input.BankKeeper.AllBalances(sdk.WrapSDKContext(ctx), req)
	require.NoError(t, err)
	require.Equal(t, res.Balances.Len(), 1)
	require.Nil(t, res.Pagination.NextKey)
}

func TestQueryTotalSupply(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	mintCoins := sdk.NewCoins(sdk.NewInt64Coin("test", 400000000))
	expectedTotalSupply := initialTotalSupply().Add(mintCoins...)
	require.NoError(t, input.BankKeeper.MintCoins(ctx, authtypes.Minter, mintCoins))

	res, err := input.BankKeeper.TotalSupply(sdk.WrapSDKContext(ctx), &types.QueryTotalSupplyRequest{})
	require.NoError(t, err)
	require.NotNil(t, res)

	require.Equal(t, expectedTotalSupply, res.Supply)
}

func TestQueryTotalSupplyOf(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	test1Supply := sdk.NewInt64Coin("foo", 4000000)
	test2Supply := sdk.NewInt64Coin("bar", 700000000)
	mintCoins := sdk.NewCoins(test1Supply, test2Supply)
	require.
		NoError(t, input.BankKeeper.MintCoins(ctx, authtypes.Minter, mintCoins))

	_, err := input.BankKeeper.SupplyOf(sdk.WrapSDKContext(ctx), &types.QuerySupplyOfRequest{})
	require.Error(t, err)

	res, err := input.BankKeeper.SupplyOf(sdk.WrapSDKContext(ctx), &types.QuerySupplyOfRequest{Denom: test1Supply.Denom})
	require.NoError(t, err)
	require.NotNil(t, res)

	require.Equal(t, test1Supply, res.Amount)
}

func TestQueryParams(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	res, err := input.BankKeeper.Params(sdk.WrapSDKContext(ctx), &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, input.BankKeeper.GetParams(ctx), res.GetParams())
}
