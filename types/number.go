package types

import (
	"database/sql"
	"strconv"

	"github.com/shopspring/decimal"
)

type NullInt64 struct {
	sql.NullInt64
}

func (i NullInt64) MarshalJSON() ([]byte, error) {
	var jsonText string
	if !i.Valid {
		jsonText = "null"
	} else {
		jsonText = strconv.FormatInt(i.Int64, 10)
	}
	return []byte(jsonText), nil
}

func (i *NullInt64) UnmarshalJSON(data []byte) error {
	text := string(data)
	switch text {
	case "null":
	default:
		value, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			return err
		}
		i.Valid = true
		i.Int64 = value
	}
	return nil
}

type JsonNumberDecimal struct {
	decimal.Decimal
}

func NewJsonNumberDecimal(d decimal.Decimal) JsonNumberDecimal {
	return JsonNumberDecimal{d}
}

func (d JsonNumberDecimal) MarshalJSON() ([]byte, error) {
	return d.Decimal.MarshalText()
}
