package keeper_test

import (
	"testing"

	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	"github.com/stretchr/testify/require"

	"github.com/initia-labs/initia/x/move/keeper"
	"github.com/initia-labs/initia/x/move/types"
	vmtypes "github.com/initia-labs/initiavm/types"
)

func TestInitializeIBCCoin(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	moveBankKeeper := keeper.NewMoveBankKeeper(&input.MoveKeeper)

	denomTrace := transfertypes.DenomTrace{
		Path:      "",
		BaseDenom: "ufoo",
	}
	denom := denomTrace.IBCDenom()

	err := moveBankKeeper.InitializeCoin(ctx, denom)
	require.NoError(t, err)

	metadata, err := types.MetadataAddressFromDenom(denom)
	require.NoError(t, err)

	bz, err := input.MoveKeeper.GetResourceBytes(ctx, metadata, vmtypes.StructTag{
		Address:  vmtypes.StdAddress,
		Module:   types.MoveModuleNameFungibleAsset,
		Name:     types.ResourceNameMetadata,
		TypeArgs: []vmtypes.TypeTag{},
	})
	require.NoError(t, err)

	symbol := types.ReadSymbolFromMetadata(bz)
	require.Equal(t, denom, symbol)
}
