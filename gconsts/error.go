package gconsts

import "gitlab.com/snap-clickstaff/go-app/lib/gmeta"

const (
	ErrorCodeSuccess         gmeta.ErrorCode = "success"
	ErrorCodeUnknown         gmeta.ErrorCode = "error_unknown"
	ErrorCodeSystem          gmeta.ErrorCode = "error_system"
	ErrorCodeAuth            gmeta.ErrorCode = "error_auth"
	ErrorCodeAccess          gmeta.ErrorCode = "error_access"
	ErrorCodeAuthBlocked     gmeta.ErrorCode = "error_auth_blocked"
	ErrorCodeIgnored         gmeta.ErrorCode = "error_ignored"
	ErrorCodeInvalidParams   gmeta.ErrorCode = "error_invalid_params"
	ErrorCodeInvalidData     gmeta.ErrorCode = "error_invalid_data"
	ErrorCodeTimeout         gmeta.ErrorCode = "error_timeout"
	ErrorCodeTooManyRequests gmeta.ErrorCode = "error_too_many_requests"
	ErrorCodeDataNotFound    gmeta.ErrorCode = "error_data_not_found"
	ErrorCodeDataExists      gmeta.ErrorCode = "error_data_exists"
	ErrorCodeDataExpired     gmeta.ErrorCode = "error_data_expired"
	ErrorCodeDataClosed      gmeta.ErrorCode = "error_data_closed"
	ErrorCodeDataLocked      gmeta.ErrorCode = "error_data_locked"
	ErrorCodeDataDuplicate   gmeta.ErrorCode = "error_data_duplicate"
	ErrorCodeDependency      gmeta.ErrorCode = "error_dependency"
	ErrorCodeMaintenance     gmeta.ErrorCode = "error_maintenance"

	ErrorCodeAddress                          gmeta.ErrorCode = "error_address"
	ErrorCodeAmount                           gmeta.ErrorCode = "error_amount"
	ErrorCodeAmountTooHigh                    gmeta.ErrorCode = "error_amount_too_high"
	ErrorCodeAmountTooHighWithValue           gmeta.ErrorCode = "error_amount_too_high_with_value"
	ErrorCodeAmountTooLow                     gmeta.ErrorCode = "error_amount_too_low"
	ErrorCodeAmountTooLowWithValue            gmeta.ErrorCode = "error_amount_too_low_with_value"
	ErrorCodeAuthTOTP                         gmeta.ErrorCode = "error_auth_totp"
	ErrorCodeAuthInput                        gmeta.ErrorCode = "error_auth_input"
	ErrorCodeAuthPassword                     gmeta.ErrorCode = "error_auth_password"
	ErrorCodeAuthTelegram                     gmeta.ErrorCode = "error_auth_telegram"
	ErrorCodeBalanceNotEnough                 gmeta.ErrorCode = "error_balance_not_enough"
	ErrorCodeBlockchainBalanceNotEnoughForFee gmeta.ErrorCode = "error_blockchain_balance_not_enough_for_fee"
	ErrorCodeBlockchainNetwork                gmeta.ErrorCode = "error_blockchain_network"
	ErrorCodeChannelNotAvailable              gmeta.ErrorCode = "error_channel_not_available"
	ErrorCodeCountryBanned                    gmeta.ErrorCode = "error_country_banned"
	ErrorCodeCurrency                         gmeta.ErrorCode = "error_currency"
	ErrorCodeDobInvalid                       gmeta.ErrorCode = "error_dob_invalid"
	ErrorCodeEmailInvalid                     gmeta.ErrorCode = "error_email_invalid"
	ErrorCodeEmailNotSent                     gmeta.ErrorCode = "error_email_not_sent"
	ErrorCodeFeatureNotSupport                gmeta.ErrorCode = "error_feature_not_support"
	ErrorCodeKycRequestInvalidStatus          gmeta.ErrorCode = "error_kyc_request_invalid_status"
	ErrorCodeKycRequired                      gmeta.ErrorCode = "error_kyc_required"
	ErrorCodeKycUserBlacklist                 gmeta.ErrorCode = "error_kyc_user_blacklist"
	ErrorCodeOrderConcurrent                  gmeta.ErrorCode = "error_order_concurrent"
	ErrorCodeOrderDuplicated                  gmeta.ErrorCode = "error_order_duplicated"
	ErrorCodeOrderInvalid                     gmeta.ErrorCode = "error_order_invalid"
	ErrorCodeOrderNotFound                    gmeta.ErrorCode = "error_order_not_found"
	ErrorCodeOrderStatus                      gmeta.ErrorCode = "error_order_status"
	ErrorCodeStatus                           gmeta.ErrorCode = "error_status"
	ErrorCodeUserActionLocked                 gmeta.ErrorCode = "error_user_action_locked"
	ErrorCodeUserLogin                        gmeta.ErrorCode = "error_user_login"
	ErrorCodeUserNotFound                     gmeta.ErrorCode = "error_user_not_found"
	ErrorCodeUserReferralNotFound             gmeta.ErrorCode = "error_user_referral_not_found"
	ErrorCodeUserSubmittedOverLimit           gmeta.ErrorCode = "error_user_submitted_over_limit"
	ErrorCodeUserTierTooHigh                  gmeta.ErrorCode = "error_user_tier_too_high"
	ErrorCodeUserTierNotEnough                gmeta.ErrorCode = "error_user_tier_not_enough"
	ErrorCodeUserLinkedAnotherAccount         gmeta.ErrorCode = "error_user_linked_another_account"
	ErrorCodeGameRestrictedByVoucher          gmeta.ErrorCode = "error_game_restricted_by_voucher"
)

var (
	ErrorUnknown         = gmeta.NewOurError(ErrorCodeUnknown)
	ErrorSystem          = gmeta.NewOurError(ErrorCodeSystem)
	ErrorAuth            = gmeta.NewOurError(ErrorCodeAuth)
	ErrorAccess          = gmeta.NewOurError(ErrorCodeAccess)
	ErrorAuthBlocked     = gmeta.NewOurError(ErrorCodeAuthBlocked)
	ErrorIgnored         = gmeta.NewOurError(ErrorCodeIgnored)
	ErrorInvalidParams   = gmeta.NewOurError(ErrorCodeInvalidParams)
	ErrorInvalidData     = gmeta.NewOurError(ErrorCodeInvalidData)
	ErrorTimeout         = gmeta.NewOurError(ErrorCodeTimeout)
	ErrorTooManyRequests = gmeta.NewOurError(ErrorCodeTooManyRequests)
	ErrorDataNotFound    = gmeta.NewOurError(ErrorCodeDataNotFound)
	ErrorDataExists      = gmeta.NewOurError(ErrorCodeDataExists)
	ErrorDataExpired     = gmeta.NewOurError(ErrorCodeDataExpired)
	ErrorDataClosed      = gmeta.NewOurError(ErrorCodeDataClosed)
	ErrorDataLocked      = gmeta.NewOurError(ErrorCodeDataLocked)
	ErrorDataDuplicate   = gmeta.NewOurError(ErrorCodeDataDuplicate)
	ErrorDataMaintenance = gmeta.NewOurError(ErrorCodeMaintenance)

	ErrorAddress                          = gmeta.NewOurError(ErrorCodeAddress)
	ErrorAmount                           = gmeta.NewOurError(ErrorCodeAmount)
	ErrorAmountTooHigh                    = gmeta.NewOurError(ErrorCodeAmountTooHigh)
	ErrorAmountTooHighWithValue           = gmeta.NewOurError(ErrorCodeAmountTooHighWithValue)
	ErrorAmountTooLow                     = gmeta.NewOurError(ErrorCodeAmountTooLow)
	ErrorAmountTooLowWithValue            = gmeta.NewOurError(ErrorCodeAmountTooLowWithValue)
	ErrorAuthInput                        = gmeta.NewOurError(ErrorCodeAuthInput)
	ErrorAuthTOTP                         = gmeta.NewOurError(ErrorCodeAuthTOTP)
	ErrorAuthPassword                     = gmeta.NewOurError(ErrorCodeAuthPassword)
	ErrorAuthTelegram                     = gmeta.NewOurError(ErrorCodeAuthTelegram)
	ErrorBalanceNotEnough                 = gmeta.NewOurError(ErrorCodeBalanceNotEnough)
	ErrorBlockchainBalanceNotEnoughForFee = gmeta.NewOurError(ErrorCodeBlockchainBalanceNotEnoughForFee)
	ErrorBlockchainNetwork                = gmeta.NewOurError(ErrorCodeBlockchainNetwork)
	ErrorChannelNotAvailable              = gmeta.NewOurError(ErrorCodeChannelNotAvailable)
	ErrorCountryBanned                    = gmeta.NewOurError(ErrorCodeCountryBanned)
	ErrorDependency                       = gmeta.NewOurError(ErrorCodeDependency)
	ErrorCurrency                         = gmeta.NewOurError(ErrorCodeCurrency)
	ErrorDobInvalid                       = gmeta.NewOurError(ErrorCodeDobInvalid)
	ErrorEmailInvalid                     = gmeta.NewOurError(ErrorCodeEmailInvalid)
	ErrorEmailNotSent                     = gmeta.NewOurError(ErrorCodeEmailNotSent)
	ErrorFeatureNotSupport                = gmeta.NewOurError(ErrorCodeFeatureNotSupport)
	ErrorKycRequestInvalidStatus          = gmeta.NewOurError(ErrorCodeKycRequestInvalidStatus)
	ErrorKycRequired                      = gmeta.NewOurError(ErrorCodeKycRequired)
	ErrorKycUserBlacklist                 = gmeta.NewOurError(ErrorCodeKycUserBlacklist)
	ErrorOrderConcurrent                  = gmeta.NewOurError(ErrorCodeOrderConcurrent)
	ErrorOrderDuplicated                  = gmeta.NewOurError(ErrorCodeOrderDuplicated)
	ErrorOrderInvalid                     = gmeta.NewOurError(ErrorCodeOrderInvalid)
	ErrorOrderNotFound                    = gmeta.NewOurError(ErrorCodeOrderNotFound)
	ErrorOrderStatus                      = gmeta.NewOurError(ErrorCodeOrderStatus)
	ErrorStatus                           = gmeta.NewOurError(ErrorCodeStatus)
	ErrorUserActionLocked                 = gmeta.NewOurError(ErrorCodeUserActionLocked)
	ErrorUserLogin                        = gmeta.NewOurError(ErrorCodeUserLogin)
	ErrorUserNotFound                     = gmeta.NewOurError(ErrorCodeUserNotFound)
	ErrorUserReferralNotFound             = gmeta.NewOurError(ErrorCodeUserReferralNotFound)
	ErrorUserSubmittedOverLimit           = gmeta.NewOurError(ErrorCodeUserSubmittedOverLimit)
	ErrorUserTierTooHigh                  = gmeta.NewOurError(ErrorCodeUserTierTooHigh)
	ErrorUserTierNotEnough                = gmeta.NewOurError(ErrorCodeUserTierNotEnough)
	ErrorUserLinkedAnotherAccount         = gmeta.NewOurError(ErrorCodeUserLinkedAnotherAccount)
	ErrorGameRestrictedByVoucher          = gmeta.NewOurError(ErrorCodeGameRestrictedByVoucher)
)