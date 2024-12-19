package gmeta

type Coin interface {
	GetCurrency() Currency
	GetNetwork() BlockchainNetwork
	GetIndexCode() string
}
