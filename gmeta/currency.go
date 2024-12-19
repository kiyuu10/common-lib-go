package gmeta

import (
	"strings"

	"github.com/shopspring/decimal"

	comutils "gitea.alchemymagic.app/snap/go-common/utils"
)

type Currency string

func (c Currency) String() string {
	return string(c)
}

func (c Currency) StringU() string {
	return strings.ToUpper(c.String())
}

func (c Currency) StringL() string {
	return strings.ToLower(c.String())
}

func (c Currency) ToUpper() Currency {
	return Currency(c.StringU())
}

func (c Currency) ToLower() Currency {
	return Currency(c.StringL())
}

func NewCurrency(code string) Currency {
	return Currency(code)
}

func NewCurrencyU(code string) Currency {
	return NewCurrency(strings.ToUpper(code))
}

type (
	CurrencyMeta struct {
		Code          Currency
		DecimalPlaces uint8
	}
	CurrencyConversionRate struct {
		FromCurrency Currency `json:"from_currency"`
		ToCurrency   Currency `json:"to_currency"`
		Exponent     int32    `json:"exponent"`
	}
	CurrencyAmount struct {
		Currency Currency        `gorm:"column:currency" json:"currency" validate:"required"`
		Value    decimal.Decimal `gorm:"column:value" json:"value" validate:"required"`
	}
)

type CurrencyAmountMap map[Currency]decimal.Decimal

type CurrencyErrorMap map[Currency]error

type AmountMarkup struct {
	isPercentage bool
	value        decimal.Decimal
}

func NewAmountModifier(strValue string) (*AmountMarkup, error) {
	handler := AmountMarkup{}
	err := handler.UnmarshalText([]byte(strValue))
	if err != nil {
		return nil, err
	}

	return &handler, nil
}

func (am *AmountMarkup) String() string {
	strValue := am.value.String()
	if am.isPercentage {
		strValue += "%"
	}

	return strValue
}

func (am *AmountMarkup) For(value decimal.Decimal) decimal.Decimal {
	if am.value.IsZero() {
		return value
	}

	if am.isPercentage {
		rate := comutils.DecimalDivide(am.value, decimal.NewFromInt(100))
		return value.Add(value.Mul(rate))
	}

	return value.Add(am.value)
}

func (am AmountMarkup) MarshalText() ([]byte, error) {
	return []byte(am.String()), nil
}

func (am *AmountMarkup) UnmarshalText(text []byte) (err error) {
	strVal := strings.TrimSpace(string(text))

	if strings.HasSuffix(strVal, "%") {
		am.isPercentage = true
		strVal = strVal[:len(strVal)-1]
	}

	am.value, err = decimal.NewFromString(strVal)
	return
}

func (am AmountMarkup) MarshalBinary() ([]byte, error) {
	return am.MarshalText()
}

func (am *AmountMarkup) UnmarshalBinary(data []byte) (err error) {
	return am.UnmarshalText(data)
}
