package gconsts

import "gitlab.com/snap-clickstaff/go-app/lib/gmeta"

const (
	DateFormatISO          = "2006-01-02"
	DateFormatSlug         = "20060102"
	DateTimeFormat         = "2006-01-02 15:04:05"
	DateTimeFormatTz       = "2006-01-02 15:04:05 MST"
	DateTimeFormatTzOffset = "2006-01-02 15:04:05 -07:00"
	DateTimeFormatSlug     = "20060102150405"

	DateMonthFormatISO  = "2006-01"
	DateMonthFormatSlug = "200601"

	DateTimeStamp = "15:04:05"
)

const (
	DirectionTypeUnknown gmeta.Direction = 0
	DirectionTypeSend    gmeta.Direction = -1
	DirectionTypeReceive gmeta.Direction = 1
)

const (
	CommonStatusActive   gmeta.CommonStatus = 1
	CommonStatusUnknown  gmeta.CommonStatus = 0
	CommonStatusInactive gmeta.CommonStatus = -1
	CommonStatusCanceled gmeta.CommonStatus = -2
)
