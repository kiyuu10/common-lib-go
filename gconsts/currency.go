package gconsts

import (
	"fmt"

	"gitea.alchemymagic.app/snap/go-common/types"
	comutils "gitea.alchemymagic.app/snap/go-common/utils"
	"gitlab.com/snap-clickstaff/go-app/lib/gmeta"
)

const (
	CoinCodeSep                = "@"
	CurrencyUSD gmeta.Currency = "USD"
	CurrencyCNY gmeta.Currency = "CNY"
	CurrencyMYR gmeta.Currency = "MYR"
	CurrencyTHB gmeta.Currency = "THB"
	CurrencyIDR gmeta.Currency = "IDR"

	CurrencyUnknown           gmeta.Currency = "UNKNOWN"
	CurrencyBitcoin           gmeta.Currency = "BTC"
	CurrencyBitcoinCash       gmeta.Currency = "BCH"
	CurrencyEthereum          gmeta.Currency = "ETH"
	CurrencyLitecoin          gmeta.Currency = "LTC"
	CurrencyTron              gmeta.Currency = "TRX"
	CurrencyRipple            gmeta.Currency = "XRP"
	CurrencyBinanceCoin       gmeta.Currency = "BNB"
	CurrencyAvalanche         gmeta.Currency = "AVAX"
	CurrencyOKEx              gmeta.Currency = "OKT"
	CurrencyFantom            gmeta.Currency = "FTM"
	CurrencyPolygonMatic      gmeta.Currency = "MATIC"
	CurrencyCelo              gmeta.Currency = "CELO"
	CurrencyHarmony           gmeta.Currency = "ONE"
	CurrencyCronos            gmeta.Currency = "CRO"
	CurrencyHECO              gmeta.Currency = "HT"
	CurrencyXDAI              gmeta.Currency = "XDAI"
	CurrencyMoonriver         gmeta.Currency = "MOVR"
	CurrencyVelas             gmeta.Currency = "VLX"
	CurrencyFuse              gmeta.Currency = "FUSE"
	CurrencyTetherUSD         gmeta.Currency = "USDT"
	CurrencyBinanceUSD        gmeta.Currency = "BUSD"
	CurrencyHunny             gmeta.Currency = "HUNNY"
	CurrencyHUSD              gmeta.Currency = "HUSD"
	CurrencyCake              gmeta.Currency = "CAKE"
	CurrencySolana            gmeta.Currency = "SOL"
	CurrencyUSDCoin           gmeta.Currency = "USDC"
	CurrencyAxelarWrappedUSDC gmeta.Currency = "USDC_AXL"
	CurrencyLove              gmeta.Currency = "LOVE"
	CurrencyTON               gmeta.Currency = "TON"
	CurrencyDogecoin          gmeta.Currency = "DOGE"

	CurrencySubBitcoinSatoshi gmeta.Currency = "satoshi"
	CurrencySubEvmGwei        gmeta.Currency = "gwei"
	CurrencySubEvmWei         gmeta.Currency = "wei"
	CurrencySubTronSun        gmeta.Currency = "sun"
	CurrencySubTronBandwidth  gmeta.Currency = "bandwidth"
	CurrencySubTronEnergy     gmeta.Currency = "energy"
	CurrencySubRippleDrop     gmeta.Currency = "drop"
	CurrencySubSolanaLamport  gmeta.Currency = "lamport"
	CurrencySubTONGrams       gmeta.Currency = "grams"
)

var (
	CurrencyIdenticalMap    = make(map[gmeta.Currency]gmeta.Currency)
	CurrencyConversionRates = []gmeta.CurrencyConversionRate{
		{
			FromCurrency: CurrencyBitcoin,
			ToCurrency:   CurrencySubBitcoinSatoshi,
			Exponent:     8,
		},
		{
			FromCurrency: CurrencyLitecoin,
			ToCurrency:   CurrencySubBitcoinSatoshi, // Litecoin is litoshi, but we use satoshi for adaptation
			Exponent:     8,
		},
		{
			FromCurrency: CurrencyEthereum,
			ToCurrency:   CurrencySubEvmGwei,
			Exponent:     9,
		},
		{
			FromCurrency: CurrencyEthereum,
			ToCurrency:   CurrencySubEvmWei,
			Exponent:     18,
		},
		{
			FromCurrency: CurrencySubEvmGwei,
			ToCurrency:   CurrencySubEvmWei,
			Exponent:     9,
		},
		{
			FromCurrency: CurrencyTron,
			ToCurrency:   CurrencySubTronSun,
			Exponent:     6,
		},
		{
			FromCurrency: CurrencyTron,
			ToCurrency:   CurrencySubTronBandwidth,
			Exponent:     5,
		},
		{
			FromCurrency: CurrencySubTronBandwidth,
			ToCurrency:   CurrencySubTronSun,
			Exponent:     1,
		},
		{
			FromCurrency: CurrencyTron,
			ToCurrency:   CurrencySubTronEnergy,
			Exponent:     5,
		},
		{
			FromCurrency: CurrencySubTronEnergy,
			ToCurrency:   CurrencySubTronSun,
			Exponent:     1,
		},
		{
			FromCurrency: CurrencyRipple,
			ToCurrency:   CurrencySubRippleDrop,
			Exponent:     6,
		},
		{
			FromCurrency: CurrencySolana,
			ToCurrency:   CurrencySubSolanaLamport,
			Exponent:     9,
		},
		{
			FromCurrency: CurrencyTON,
			ToCurrency:   CurrencySubTONGrams,
			Exponent:     9,
		},
	}
	CurrencyConversionRateMap = comutils.ListToMapFlat(CurrencyConversionRates, func(i int, rate gmeta.CurrencyConversionRate) []types.KeyValue[string, gmeta.CurrencyConversionRate] {
		var reveredItem = types.KeyValue[string, gmeta.CurrencyConversionRate]{

			Key: fmt.Sprintf("%v-%v", rate.ToCurrency, rate.FromCurrency),
			Value: gmeta.CurrencyConversionRate{
				FromCurrency: rate.ToCurrency,
				ToCurrency:   rate.FromCurrency,
				Exponent:     -rate.Exponent,
			},
		}
		return []types.KeyValue[string, gmeta.CurrencyConversionRate]{
			{
				Key:   fmt.Sprintf("%v-%v", rate.FromCurrency, rate.ToCurrency),
				Value: rate,
			},
			reveredItem,
		}
	})
)

func RegisterIdenticalCurrency(currency gmeta.Currency, identicalTo gmeta.Currency) {
	if currency != identicalTo {
		CurrencyIdenticalMap[currency] = identicalTo
	}
}

var CurrencyMetaMap = map[gmeta.Currency]gmeta.CurrencyMeta{
	CurrencyUSD: {
		Code:          CurrencyUSD,
		DecimalPlaces: 2,
	},
	CurrencyCNY: {
		Code:          CurrencyCNY,
		DecimalPlaces: 2,
	},
	CurrencyHunny: {
		Code:          CurrencyHunny,
		DecimalPlaces: 18,
	},

	CurrencyBitcoin: {
		Code:          CurrencyBitcoin,
		DecimalPlaces: 8,
	},
	CurrencyEthereum: {
		Code:          CurrencyEthereum,
		DecimalPlaces: 18,
	},
	CurrencyTron: {
		Code:          CurrencyTron,
		DecimalPlaces: 6,
	},
	CurrencyRipple: {
		Code:          CurrencyRipple,
		DecimalPlaces: 6,
	},
}

func GetCurrencyMeta(currency gmeta.Currency) (_ gmeta.CurrencyMeta, exists bool) {
	if currencyMeta, ok := CurrencyMetaMap[currency]; ok {
		return currencyMeta, true
	}
	if identicalCurrency, ok := CurrencyIdenticalMap[currency]; ok {
		if currencyMeta, ok := CurrencyMetaMap[identicalCurrency]; ok {
			currencyMeta.Code = currency
			return currencyMeta, true
		}
	}
	return
}
