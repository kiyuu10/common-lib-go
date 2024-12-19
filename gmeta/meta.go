package gmeta

type (
	E struct{}
	O map[string]any
)

type (
	Direction    int8
	CommonStatus int8
)

type TimeRange struct {
	FromTime int64 `json:"from_time" validate:"required"`
	ToTime   int64 `json:"to_time" validate:"required"`
}
