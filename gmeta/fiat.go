package gmeta

type (
	FiatBankCode string
)

type (
	//FiatCurrency for type matching
	FiatCurrency struct {
		Coin
		Currency Currency
	}
)

func (f *FiatCurrency) GetNetwork() BlockchainNetwork {
	return "FIAT"
}

func (f *FiatCurrency) GetCurrency() Currency {
	return f.Currency
}

func (f *FiatBankCode) String() string {
	return string(*f)
}
