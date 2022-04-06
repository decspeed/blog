package keeper

import (
	"github.com/decspeed/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
