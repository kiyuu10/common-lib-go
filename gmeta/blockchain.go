package gmeta

import (
	"strings"

	"github.com/shopspring/decimal"
)

type BlockchainType string

func (t BlockchainType) IsUtxoType() bool {
	return strings.HasPrefix(string(t), "utxo-")
}

type BlockchainNetwork string

func NewBlockchainNetwork(code string) BlockchainNetwork {
	return BlockchainNetwork(strings.ToUpper(code))
}

func (bn BlockchainNetwork) String() string {
	return string(bn)
}

func (bn BlockchainNetwork) AsLower() BlockchainNetwork {
	return BlockchainNetwork(strings.ToLower(string(bn)))
}

type BlockchainCoinIndex struct {
	Currency Currency          `json:"currency"`
	Network  BlockchainNetwork `json:"network"`
}

func (bi *BlockchainCoinIndex) GetCurrency() Currency {
	return bi.Currency
}

func (bi *BlockchainCoinIndex) GetNetwork() BlockchainNetwork {
	return bi.Network
}

type NetworkCurrency struct {
	Network  BlockchainNetwork `json:"network" validate:"required"`
	Currency Currency          `json:"currency" validate:"required"`
}

func (nc NetworkCurrency) GetCurrency() Currency {
	return nc.Currency
}

func (nc NetworkCurrency) GetNetwork() BlockchainNetwork {
	return nc.Network
}

func (nc NetworkCurrency) GetIndexCode() string {
	return nc.GetCurrency().String() + "@" + nc.GetNetwork().String()
}

type NetworkCurrencyAmount struct {
	Network  BlockchainNetwork `json:"network" validate:"required"`
	Currency Currency          `json:"currency" validate:"required"`
	Value    decimal.Decimal   `json:"value" validate:"required"`
}

type NetworkCurrencyAmountUSD struct {
	Network       BlockchainNetwork `json:"network" validate:"required"`
	Currency      Currency          `json:"currency" validate:"required"`
	CurrencyValue decimal.Decimal   `json:"currency_value" validate:"required"`
	USDValue      decimal.Decimal   `json:"usd_value" validate:"required"`
}

type BlockchainTypedDataDomain struct {
	Name      string         `json:"name" validate:"required"`
	Version   string         `json:"version" validate:"required"`
	ChainType BlockchainType `json:"chain_type" validate:"required"`
	ChainCode string         `json:"chain_code" validate:"required"`
	Contract  string         `json:"verifying_contract" validate:"required"`
}

type (
	BlockchainTxnStatus          int8
	BlockchainMerchantID         uint8
	BlockchainProductStatus      int8
	BlockchainProductTxnAction   int8
	BlockchainProductEvent       string
	BlockchainProductOfferStatus int8
)
