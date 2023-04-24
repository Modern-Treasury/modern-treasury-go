package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type PaymentReference struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The id of the referenceable to search for. Must be accompanied by the
	// referenceable_type or will return an error.
	ReferenceableID string `json:"referenceable_id,required"`
	// One of the referenceable types. This must be accompanied by the id of the
	// referenceable or will return an error.
	ReferenceableType PaymentReferenceReferenceableType `json:"referenceable_type,required"`
	// The actual reference number assigned by the bank.
	ReferenceNumber string `json:"reference_number,required"`
	// The type of reference number.
	ReferenceNumberType PaymentReferenceReferenceNumberType `json:"reference_number_type,required"`
	JSON                PaymentReferenceJSON
}

type PaymentReferenceJSON struct {
	ID                  pjson.Metadata
	Object              pjson.Metadata
	LiveMode            pjson.Metadata
	CreatedAt           pjson.Metadata
	UpdatedAt           pjson.Metadata
	ReferenceableID     pjson.Metadata
	ReferenceableType   pjson.Metadata
	ReferenceNumber     pjson.Metadata
	ReferenceNumberType pjson.Metadata
	Raw                 []byte
	Extras              map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PaymentReference using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *PaymentReference) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type PaymentReferenceReferenceableType string

const (
	PaymentReferenceReferenceableTypePaymentOrder PaymentReferenceReferenceableType = "payment_order"
	PaymentReferenceReferenceableTypeReversal     PaymentReferenceReferenceableType = "reversal"
	PaymentReferenceReferenceableTypeReturn       PaymentReferenceReferenceableType = "return"
)

type PaymentReferenceReferenceNumberType string

const (
	PaymentReferenceReferenceNumberTypeACHOriginalTraceNumber          PaymentReferenceReferenceNumberType = "ach_original_trace_number"
	PaymentReferenceReferenceNumberTypeACHTraceNumber                  PaymentReferenceReferenceNumberType = "ach_trace_number"
	PaymentReferenceReferenceNumberTypeBankprovPaymentActivityDate     PaymentReferenceReferenceNumberType = "bankprov_payment_activity_date"
	PaymentReferenceReferenceNumberTypeBankprovPaymentID               PaymentReferenceReferenceNumberType = "bankprov_payment_id"
	PaymentReferenceReferenceNumberTypeBnkDevPrenotificationID         PaymentReferenceReferenceNumberType = "bnk_dev_prenotification_id"
	PaymentReferenceReferenceNumberTypeBnkDevTransferID                PaymentReferenceReferenceNumberType = "bnk_dev_transfer_id"
	PaymentReferenceReferenceNumberTypeBofaEndToEndID                  PaymentReferenceReferenceNumberType = "bofa_end_to_end_id"
	PaymentReferenceReferenceNumberTypeBofaTransactionID               PaymentReferenceReferenceNumberType = "bofa_transaction_id"
	PaymentReferenceReferenceNumberTypeCheckNumber                     PaymentReferenceReferenceNumberType = "check_number"
	PaymentReferenceReferenceNumberTypeColumnFxQuoteID                 PaymentReferenceReferenceNumberType = "column_fx_quote_id"
	PaymentReferenceReferenceNumberTypeColumnTransferID                PaymentReferenceReferenceNumberType = "column_transfer_id"
	PaymentReferenceReferenceNumberTypeCrossRiverPaymentID             PaymentReferenceReferenceNumberType = "cross_river_payment_id"
	PaymentReferenceReferenceNumberTypeCrossRiverTransactionID         PaymentReferenceReferenceNumberType = "cross_river_transaction_id"
	PaymentReferenceReferenceNumberTypeCurrencycloudConversionID       PaymentReferenceReferenceNumberType = "currencycloud_conversion_id"
	PaymentReferenceReferenceNumberTypeCurrencycloudPaymentID          PaymentReferenceReferenceNumberType = "currencycloud_payment_id"
	PaymentReferenceReferenceNumberTypeDcBankTransactionID             PaymentReferenceReferenceNumberType = "dc_bank_transaction_id"
	PaymentReferenceReferenceNumberTypeDwollaTransactionID             PaymentReferenceReferenceNumberType = "dwolla_transaction_id"
	PaymentReferenceReferenceNumberTypeEftTraceNumber                  PaymentReferenceReferenceNumberType = "eft_trace_number"
	PaymentReferenceReferenceNumberTypeFedwireImad                     PaymentReferenceReferenceNumberType = "fedwire_imad"
	PaymentReferenceReferenceNumberTypeFedwireOmad                     PaymentReferenceReferenceNumberType = "fedwire_omad"
	PaymentReferenceReferenceNumberTypeFirstRepublicInternalID         PaymentReferenceReferenceNumberType = "first_republic_internal_id"
	PaymentReferenceReferenceNumberTypeGoldmanSachsCollectionRequestID PaymentReferenceReferenceNumberType = "goldman_sachs_collection_request_id"
	PaymentReferenceReferenceNumberTypeGoldmanSachsEndToEndID          PaymentReferenceReferenceNumberType = "goldman_sachs_end_to_end_id"
	PaymentReferenceReferenceNumberTypeGoldmanSachsPaymentRequestID    PaymentReferenceReferenceNumberType = "goldman_sachs_payment_request_id"
	PaymentReferenceReferenceNumberTypeGoldmanSachsRequestID           PaymentReferenceReferenceNumberType = "goldman_sachs_request_id"
	PaymentReferenceReferenceNumberTypeGoldmanSachsUniquePaymentID     PaymentReferenceReferenceNumberType = "goldman_sachs_unique_payment_id"
	PaymentReferenceReferenceNumberTypeInteracMessageID                PaymentReferenceReferenceNumberType = "interac_message_id"
	PaymentReferenceReferenceNumberTypeJpmcCcn                         PaymentReferenceReferenceNumberType = "jpmc_ccn"
	PaymentReferenceReferenceNumberTypeJpmcCustomerReferenceID         PaymentReferenceReferenceNumberType = "jpmc_customer_reference_id"
	PaymentReferenceReferenceNumberTypeJpmcEndToEndID                  PaymentReferenceReferenceNumberType = "jpmc_end_to_end_id"
	PaymentReferenceReferenceNumberTypeJpmcFirmRootID                  PaymentReferenceReferenceNumberType = "jpmc_firm_root_id"
	PaymentReferenceReferenceNumberTypeJpmcP3ID                        PaymentReferenceReferenceNumberType = "jpmc_p3_id"
	PaymentReferenceReferenceNumberTypeJpmcPaymentBatchID              PaymentReferenceReferenceNumberType = "jpmc_payment_batch_id"
	PaymentReferenceReferenceNumberTypeJpmcPaymentInformationID        PaymentReferenceReferenceNumberType = "jpmc_payment_information_id"
	PaymentReferenceReferenceNumberTypeLobCheckID                      PaymentReferenceReferenceNumberType = "lob_check_id"
	PaymentReferenceReferenceNumberTypeOther                           PaymentReferenceReferenceNumberType = "other"
	PaymentReferenceReferenceNumberTypePartialSwiftMir                 PaymentReferenceReferenceNumberType = "partial_swift_mir"
	PaymentReferenceReferenceNumberTypePncClearingReference            PaymentReferenceReferenceNumberType = "pnc_clearing_reference"
	PaymentReferenceReferenceNumberTypePncInstructionID                PaymentReferenceReferenceNumberType = "pnc_instruction_id"
	PaymentReferenceReferenceNumberTypePncMultipaymentID               PaymentReferenceReferenceNumberType = "pnc_multipayment_id"
	PaymentReferenceReferenceNumberTypePncPaymentTraceID               PaymentReferenceReferenceNumberType = "pnc_payment_trace_id"
	PaymentReferenceReferenceNumberTypeRtpInstructionID                PaymentReferenceReferenceNumberType = "rtp_instruction_id"
	PaymentReferenceReferenceNumberTypeSignetAPIReferenceID            PaymentReferenceReferenceNumberType = "signet_api_reference_id"
	PaymentReferenceReferenceNumberTypeSignetConfirmationID            PaymentReferenceReferenceNumberType = "signet_confirmation_id"
	PaymentReferenceReferenceNumberTypeSignetRequestID                 PaymentReferenceReferenceNumberType = "signet_request_id"
	PaymentReferenceReferenceNumberTypeSilvergatePaymentID             PaymentReferenceReferenceNumberType = "silvergate_payment_id"
	PaymentReferenceReferenceNumberTypeSwiftMir                        PaymentReferenceReferenceNumberType = "swift_mir"
	PaymentReferenceReferenceNumberTypeSwiftUetr                       PaymentReferenceReferenceNumberType = "swift_uetr"
	PaymentReferenceReferenceNumberTypeUsbankPaymentID                 PaymentReferenceReferenceNumberType = "usbank_payment_id"
	PaymentReferenceReferenceNumberTypeWellsFargoPaymentID             PaymentReferenceReferenceNumberType = "wells_fargo_payment_id"
	PaymentReferenceReferenceNumberTypeWellsFargoTraceNumber           PaymentReferenceReferenceNumberType = "wells_fargo_trace_number"
)
