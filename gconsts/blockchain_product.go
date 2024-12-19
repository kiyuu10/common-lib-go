package gconsts

import "gitlab.com/snap-clickstaff/go-app/lib/gmeta"

const (
	BlockchainProductStatusInit          gmeta.BlockchainProductStatus = 1
	BlockchainProductStatusCreating      gmeta.BlockchainProductStatus = 2
	BlockchainProductStatusCreated       gmeta.BlockchainProductStatus = 3
	BlockchainProductStatusMarketListing gmeta.BlockchainProductStatus = 4
	BlockchainProductStatusMarketListed  gmeta.BlockchainProductStatus = 10
	BlockchainProductStatusSelling       gmeta.BlockchainProductStatus = 11
	// BlockchainProductStatusSold          gmeta.BlockchainProductStatus = 20
)

const (
	_ gmeta.BlockchainProductTxnAction = iota
	BlockchainProductTxnActionPublish
	BlockchainProductTxnActionMarketList
	BlockchainProductTxnActionMarketClose
)

const (
	BlockchainProductEventListingOpen  gmeta.BlockchainProductEvent = "ListingOpened"
	BlockchainProductEventListingClose gmeta.BlockchainProductEvent = "ListingClosed"
	BlockchainProductEventOfferOpen    gmeta.BlockchainProductEvent = "OfferOpened"
	BlockchainProductEventOfferCancel  gmeta.BlockchainProductEvent = "OfferCancel"
	BlockchainProductEventOneSaleBuy   gmeta.BlockchainProductEvent = "OnSaleBuy"
)

const (
	BlockchainProductOfferStatusOpen   gmeta.BlockchainProductOfferStatus = 1
	BlockchainProductOfferStatusWin    gmeta.BlockchainProductOfferStatus = 10
	BlockchainProductOfferStatusLose   gmeta.BlockchainProductOfferStatus = 9
	BlockchainProductOfferStatusCancel gmeta.BlockchainProductOfferStatus = -1
)

// var (
//	BlockchainChannelTypes = []gmeta.ChannelType{
//		ChannelTypeDstBlockchainTxn,
//		ChannelTypeDstTorqueConvert,
//	}
// )
