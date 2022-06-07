package keeper

import (
	"github.com/elision-io/elision/x/dex/types"
)

var _ types.QueryServer = Keeper{}
