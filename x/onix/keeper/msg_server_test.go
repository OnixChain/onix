package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/OnixChain/onix/testutil/keeper"
	"github.com/OnixChain/onix/x/onix/keeper"
	"github.com/OnixChain/onix/x/onix/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.OnixKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
