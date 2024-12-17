package types

import "time"

type TimeDuration struct {
	time.Duration
}

func (d TimeDuration) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *TimeDuration) UnmarshalText(text []byte) (err error) {
	duration, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}
	d.Duration = duration
	return
}

func (d TimeDuration) MarshalBinary() ([]byte, error) {
	return d.MarshalText()
}

func (d *TimeDuration) UnmarshalBinary(text []byte) (err error) {
	return d.UnmarshalText(text)
}
