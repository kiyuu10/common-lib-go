package gconsts

import (
	"gitlab.com/snap-clickstaff/go-app/lib/gmeta"
)

var (
	StakingSystemPending    gmeta.StakingSystemEventStatus = 1
	StakingSystemProcessing gmeta.StakingSystemEventStatus = 3
	StakingSystemuccess     gmeta.StakingSystemEventStatus = 11
	StakingSystemFailed     gmeta.StakingSystemEventStatus = -1
)
