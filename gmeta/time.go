package gmeta

import (
	"time"

	comutils "gitea.alchemymagic.app/snap/go-common/utils"
)

type UnixTime int64

func (t UnixTime) Time() time.Time {
	return comutils.TimeUnix(int64(t))
}

func (t UnixTime) I64() int64 {
	return int64(t)
}
