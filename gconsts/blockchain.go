package gconsts

import (
	"gitea.alchemymagic.app/snap/go-common/types"
	"gitlab.com/snap-clickstaff/go-app/lib/gmeta"
)

const (
	BlockchainTypeUtxoBitcoin  gmeta.BlockchainType = "utxo-bitcoin"
	BlockchainTypeUtxoLitecoin gmeta.BlockchainType = "utxo-ltc"
	BlockchainTypeUtxoDogecoin gmeta.BlockchainType = "utxo-doge"
	BlockchainTypeEVM          gmeta.BlockchainType = "evm"
	BlockchainTypeSolana       gmeta.BlockchainType = "solana"
	BlockchainTypeTron         gmeta.BlockchainType = "tron"
	BlockchainTypeRipple       gmeta.BlockchainType = "ripple"
	BlockchainTypeTON          gmeta.BlockchainType = "ton"
)

const (
	BlockchainNetworkBitcoin                  gmeta.BlockchainNetwork = "BTC"
	BlockchainNetworkBitcoinTestnet           gmeta.BlockchainNetwork = "BTC:TESTNET"
	BlockchainNetworkBitcoinCash              gmeta.BlockchainNetwork = "BCH"
	BlockchainNetworkBitcoinCashTestnet       gmeta.BlockchainNetwork = "BCH:TESTNET"
	BlockchainNetworkEthereum                 gmeta.BlockchainNetwork = "ETH"
	BlockchainNetworkEthereumTestRopsten      gmeta.BlockchainNetwork = "ETH:TEST_ROPSTEN"
	BlockchainNetworkEthereumTestRinkeby      gmeta.BlockchainNetwork = "ETH:TEST_RINKEBY"
	BlockchainNetworkEthereumTestnet          gmeta.BlockchainNetwork = "ETH:TESTNET"
	BlockchainNetworkLitecoin                 gmeta.BlockchainNetwork = "LTC"
	BlockchainNetworkLitecoinTestnet          gmeta.BlockchainNetwork = "LTC:TESTNET"
	BlockchainNetworkTron                     gmeta.BlockchainNetwork = "TRX"
	BlockchainNetworkTronTestNile             gmeta.BlockchainNetwork = "TRX:TEST_NILE"
	BlockchainNetworkRipple                   gmeta.BlockchainNetwork = "XRP"
	BlockchainNetworkRippleTestnet            gmeta.BlockchainNetwork = "XRP:TESTNET"
	BlockchainNetworkAurora                   gmeta.BlockchainNetwork = "AOA"
	BlockchainNetworkAuroraTestnet            gmeta.BlockchainNetwork = "AOA:TESTNET"
	BlockchainNetworkAvalanche                gmeta.BlockchainNetwork = "AVAX"
	BlockchainNetworkAvalancheTestnet         gmeta.BlockchainNetwork = "AVAX:TESTNET"
	BlockchainNetworkBinanceSmartChain        gmeta.BlockchainNetwork = "BNB"
	BlockchainNetworkBinanceSmartChainTestnet gmeta.BlockchainNetwork = "BNB:TESTNET"
	BlockchainNetworkCelo                     gmeta.BlockchainNetwork = "CELO"
	BlockchainNetworkCeloTestnet              gmeta.BlockchainNetwork = "CELO:TESTNET"
	BlockchainNetworkCronos                   gmeta.BlockchainNetwork = "CRO"
	BlockchainNetworkCronosTestnet            gmeta.BlockchainNetwork = "CRO:TESTNET"
	BlockchainNetworkFantom                   gmeta.BlockchainNetwork = "FTM"
	BlockchainNetworkFantomTestnet            gmeta.BlockchainNetwork = "FTM:TESTNET"
	BlockchainNetworkGnosis                   gmeta.BlockchainNetwork = "XDAI"
	BlockchainNetworkGnosisTestnet            gmeta.BlockchainNetwork = "XDAI:TESTNET"
	BlockchainNetworkHarmony                  gmeta.BlockchainNetwork = "ONE"
	BlockchainNetworkHarmonyTestnet           gmeta.BlockchainNetwork = "ONE:TESTNET"
	BlockchainNetworkHecoChain                gmeta.BlockchainNetwork = "HECO"
	BlockchainNetworkHecoChainTestnet         gmeta.BlockchainNetwork = "HECO:TESTNET"
	BlockchainNetworkMoonriver                gmeta.BlockchainNetwork = "MOVR"
	BlockchainNetworkMoonriverTestnet         gmeta.BlockchainNetwork = "MOVR:TESTNET"
	BlockchainNetworkOKExChain                gmeta.BlockchainNetwork = "OEC"
	BlockchainNetworkOKExChainTestnet         gmeta.BlockchainNetwork = "OEC:TESTNET"
	BlockchainNetworkPolygon                  gmeta.BlockchainNetwork = "MATIC"
	BlockchainNetworkPolygonTestnet           gmeta.BlockchainNetwork = "MATIC:TESTNET"
	BlockchainNetworkVelas                    gmeta.BlockchainNetwork = "VLX"
	BlockchainNetworkVelasTestnet             gmeta.BlockchainNetwork = "VLX:TESTNET"
	BlockchainNetworkFuse                     gmeta.BlockchainNetwork = "FUSE"
	BlockchainNetworkFuseTestnet              gmeta.BlockchainNetwork = "FUSE:TESTNET"
	BlockchainNetworkSolana                   gmeta.BlockchainNetwork = "SOL"
	BlockchainNetworkSolanaTestnet            gmeta.BlockchainNetwork = "SOL:TESTNET"
	BlockchainNetworkOptimism                 gmeta.BlockchainNetwork = "OP"
	BlockchainNetworkOptimismTestnet          gmeta.BlockchainNetwork = "OP:TESTNET"
	BlockchainNetworkArbitrumOne              gmeta.BlockchainNetwork = "ARB"
	BlockchainNetworkArbitrumTestnet          gmeta.BlockchainNetwork = "ARB:TESTNET"
	BlockchainNetworkTon                      gmeta.BlockchainNetwork = "TON"
	BlockchainNetworkTonTestnet               gmeta.BlockchainNetwork = "TON:TESTNET"
	BlockchainNetworkDogecoin                 gmeta.BlockchainNetwork = "DOGE"
	BlockchainNetworkDogecoinTestnet          gmeta.BlockchainNetwork = "DOGE:TESTNET"

	BlockchainNetworkHunnyPlayNetwork gmeta.BlockchainNetwork = "HPN"
	BlockchainNetworkFiat             gmeta.BlockchainNetwork = "FIAT"
)

const (
	BlockchainTxnStatusFailed    gmeta.BlockchainTxnStatus = -1
	BlockchainTxnStatusPending   gmeta.BlockchainTxnStatus = 1
	BlockchainTxnStatusSucceeded gmeta.BlockchainTxnStatus = 10
)

var (
	BlockchainChannelUtxoCurrencySet = types.NewHashSet(
		CurrencyBitcoin,
		CurrencyBitcoinCash,
		CurrencyLitecoin,
	)
	BlockchainUtxoMainNetworkSet = types.NewHashSet(
		BlockchainNetworkBitcoin,
		BlockchainNetworkBitcoinCash,
		BlockchainNetworkLitecoin,
	)
)
