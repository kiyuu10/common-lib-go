package gmeta

import (
	"strconv"
	"strings"

	comutils "gitea.alchemymagic.app/snap/go-common/utils"
)

type UID uint64

func (u UID) U64() uint64 {
	return uint64(u)
}

func (u UID) String() string {
	return strconv.FormatUint(uint64(u), 10)
}

type Subject string

func (s Subject) String() string {
	return string(s)
}

func (s Subject) ToLower() Subject {
	return Subject(strings.ToLower(string(s)))
}

func (s Subject) UID() (_ UID, err error) {
	uidI64, err := comutils.ParseUint64(s.String())
	if err != nil {
		return
	}
	return UID(uidI64), nil
}

func (s Subject) UidF() UID {
	uid, err := s.UID()
	comutils.PanicOnError(err)
	return uid
}

type (
	UserStatus               int8
	UserVerificationDelivery string
)

type UserVerificationType string

func (t UserVerificationType) FromBlockchain() bool {
	return strings.HasPrefix(string(t), "blockchain.")
}

type TierType uint8
