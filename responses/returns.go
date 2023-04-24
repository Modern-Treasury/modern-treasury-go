package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type ReturnObject struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The ID of the object being returned or `null`.
	ReturnableID string `json:"returnable_id,required,nullable" format:"uuid"`
	// The type of object being returned or `null`.
	ReturnableType ReturnObjectReturnableType `json:"returnable_type,required,nullable"`
	// The return code. For ACH returns, this is the required ACH return code.
	Code ReturnObjectCode `json:"code,required,nullable"`
	// Often the bank will provide an explanation for the return, which is a short
	// human readable string.
	Reason string `json:"reason,required,nullable"`
	// If the return code is `R14` or `R15` this is the date the deceased counterparty
	// passed away.
	DateOfDeath time.Time `json:"date_of_death,required,nullable" format:"date"`
	// Some returns may include additional information from the bank. In these cases,
	// this string will be present.
	AdditionalInformation string `json:"additional_information,required,nullable"`
	// The current status of the return.
	Status ReturnObjectStatus `json:"status,required"`
	// The ID of the relevant Transaction Line Item or `null`.
	TransactionLineItemID string `json:"transaction_line_item_id,required,nullable" format:"uuid"`
	// The ID of the relevant Transaction or `null`.
	TransactionID string `json:"transaction_id,required,nullable" format:"uuid"`
	// The ID of the relevant Internal Account.
	InternalAccountID string `json:"internal_account_id,required,nullable" format:"uuid"`
	// The type of return. Can be one of: `ach`, `ach_noc`, `au_becs`, `bacs`, `eft`,
	// `interac`, `manual`, `paper_item`, `wire`.
	Type ReturnObjectType `json:"type,required"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount int64 `json:"amount,required"`
	// Currency that this transaction is denominated in.
	Currency Currency `json:"currency,required,nullable"`
	// If an originating return failed to be processed by the bank, a description of
	// the failure reason will be available.
	FailureReason string `json:"failure_reason,required,nullable"`
	// The role of the return, can be `originating` or `receiving`.
	Role ReturnObjectRole `json:"role,required"`
	// An array of Payment Reference objects.
	ReferenceNumbers []ReturnObjectReferenceNumbers `json:"reference_numbers,required"`
	// The ID of the ledger transaction linked to the return.
	LedgerTransactionID string `json:"ledger_transaction_id,required,nullable" format:"uuid"`
	JSON                ReturnObjectJSON
}

type ReturnObjectJSON struct {
	ID                    pjson.Metadata
	Object                pjson.Metadata
	LiveMode              pjson.Metadata
	CreatedAt             pjson.Metadata
	UpdatedAt             pjson.Metadata
	ReturnableID          pjson.Metadata
	ReturnableType        pjson.Metadata
	Code                  pjson.Metadata
	Reason                pjson.Metadata
	DateOfDeath           pjson.Metadata
	AdditionalInformation pjson.Metadata
	Status                pjson.Metadata
	TransactionLineItemID pjson.Metadata
	TransactionID         pjson.Metadata
	InternalAccountID     pjson.Metadata
	Type                  pjson.Metadata
	Amount                pjson.Metadata
	Currency              pjson.Metadata
	FailureReason         pjson.Metadata
	Role                  pjson.Metadata
	ReferenceNumbers      pjson.Metadata
	LedgerTransactionID   pjson.Metadata
	Raw                   []byte
	Extras                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ReturnObject using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ReturnObject) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ReturnObjectReturnableType string

const (
	ReturnObjectReturnableTypeIncomingPaymentDetail ReturnObjectReturnableType = "incoming_payment_detail"
	ReturnObjectReturnableTypePaperItem             ReturnObjectReturnableType = "paper_item"
	ReturnObjectReturnableTypePaymentOrder          ReturnObjectReturnableType = "payment_order"
	ReturnObjectReturnableTypeReturn                ReturnObjectReturnableType = "return"
	ReturnObjectReturnableTypeReversal              ReturnObjectReturnableType = "reversal"
)

type ReturnObjectCode string

const (
	ReturnObjectCode901           ReturnObjectCode = "901"
	ReturnObjectCode902           ReturnObjectCode = "902"
	ReturnObjectCode903           ReturnObjectCode = "903"
	ReturnObjectCode904           ReturnObjectCode = "904"
	ReturnObjectCode905           ReturnObjectCode = "905"
	ReturnObjectCode907           ReturnObjectCode = "907"
	ReturnObjectCode908           ReturnObjectCode = "908"
	ReturnObjectCode909           ReturnObjectCode = "909"
	ReturnObjectCode910           ReturnObjectCode = "910"
	ReturnObjectCode911           ReturnObjectCode = "911"
	ReturnObjectCode912           ReturnObjectCode = "912"
	ReturnObjectCode914           ReturnObjectCode = "914"
	ReturnObjectCodeR01           ReturnObjectCode = "R01"
	ReturnObjectCodeR02           ReturnObjectCode = "R02"
	ReturnObjectCodeR03           ReturnObjectCode = "R03"
	ReturnObjectCodeR04           ReturnObjectCode = "R04"
	ReturnObjectCodeR05           ReturnObjectCode = "R05"
	ReturnObjectCodeR06           ReturnObjectCode = "R06"
	ReturnObjectCodeR07           ReturnObjectCode = "R07"
	ReturnObjectCodeR08           ReturnObjectCode = "R08"
	ReturnObjectCodeR09           ReturnObjectCode = "R09"
	ReturnObjectCodeR10           ReturnObjectCode = "R10"
	ReturnObjectCodeR11           ReturnObjectCode = "R11"
	ReturnObjectCodeR12           ReturnObjectCode = "R12"
	ReturnObjectCodeR14           ReturnObjectCode = "R14"
	ReturnObjectCodeR15           ReturnObjectCode = "R15"
	ReturnObjectCodeR16           ReturnObjectCode = "R16"
	ReturnObjectCodeR17           ReturnObjectCode = "R17"
	ReturnObjectCodeR20           ReturnObjectCode = "R20"
	ReturnObjectCodeR21           ReturnObjectCode = "R21"
	ReturnObjectCodeR22           ReturnObjectCode = "R22"
	ReturnObjectCodeR23           ReturnObjectCode = "R23"
	ReturnObjectCodeR24           ReturnObjectCode = "R24"
	ReturnObjectCodeR29           ReturnObjectCode = "R29"
	ReturnObjectCodeR31           ReturnObjectCode = "R31"
	ReturnObjectCodeR33           ReturnObjectCode = "R33"
	ReturnObjectCodeR37           ReturnObjectCode = "R37"
	ReturnObjectCodeR38           ReturnObjectCode = "R38"
	ReturnObjectCodeR39           ReturnObjectCode = "R39"
	ReturnObjectCodeR51           ReturnObjectCode = "R51"
	ReturnObjectCodeR52           ReturnObjectCode = "R52"
	ReturnObjectCodeR53           ReturnObjectCode = "R53"
	ReturnObjectCodeCurrencycloud ReturnObjectCode = "currencycloud"
)

type ReturnObjectStatus string

const (
	ReturnObjectStatusCompleted  ReturnObjectStatus = "completed"
	ReturnObjectStatusFailed     ReturnObjectStatus = "failed"
	ReturnObjectStatusPending    ReturnObjectStatus = "pending"
	ReturnObjectStatusProcessing ReturnObjectStatus = "processing"
	ReturnObjectStatusReturned   ReturnObjectStatus = "returned"
	ReturnObjectStatusSent       ReturnObjectStatus = "sent"
)

type ReturnObjectType string

const (
	ReturnObjectTypeACH       ReturnObjectType = "ach"
	ReturnObjectTypeACHNoc    ReturnObjectType = "ach_noc"
	ReturnObjectTypeAuBecs    ReturnObjectType = "au_becs"
	ReturnObjectTypeBacs      ReturnObjectType = "bacs"
	ReturnObjectTypeBook      ReturnObjectType = "book"
	ReturnObjectTypeEft       ReturnObjectType = "eft"
	ReturnObjectTypeInterac   ReturnObjectType = "interac"
	ReturnObjectTypeManual    ReturnObjectType = "manual"
	ReturnObjectTypePaperItem ReturnObjectType = "paper_item"
	ReturnObjectTypeSepa      ReturnObjectType = "sepa"
	ReturnObjectTypeWire      ReturnObjectType = "wire"
)

type ReturnObjectRole string

const (
	ReturnObjectRoleOriginating ReturnObjectRole = "originating"
	ReturnObjectRoleReceiving   ReturnObjectRole = "receiving"
)

type ReturnObjectReferenceNumbers struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The vendor reference number.
	ReferenceNumber string `json:"reference_number,required"`
	// The type of the reference number. Referring to the vendor payment id.
	ReferenceNumberType ReturnObjectReferenceNumbersReferenceNumberType `json:"reference_number_type,required"`
	JSON                ReturnObjectReferenceNumbersJSON
}

type ReturnObjectReferenceNumbersJSON struct {
	ID                  pjson.Metadata
	Object              pjson.Metadata
	LiveMode            pjson.Metadata
	CreatedAt           pjson.Metadata
	UpdatedAt           pjson.Metadata
	ReferenceNumber     pjson.Metadata
	ReferenceNumberType pjson.Metadata
	Raw                 []byte
	Extras              map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ReturnObjectReferenceNumbers
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ReturnObjectReferenceNumbers) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ReturnObjectReferenceNumbersReferenceNumberType string

const (
	ReturnObjectReferenceNumbersReferenceNumberTypeACHOriginalTraceNumber          ReturnObjectReferenceNumbersReferenceNumberType = "ach_original_trace_number"
	ReturnObjectReferenceNumbersReferenceNumberTypeACHTraceNumber                  ReturnObjectReferenceNumbersReferenceNumberType = "ach_trace_number"
	ReturnObjectReferenceNumbersReferenceNumberTypeBankprovPaymentActivityDate     ReturnObjectReferenceNumbersReferenceNumberType = "bankprov_payment_activity_date"
	ReturnObjectReferenceNumbersReferenceNumberTypeBankprovPaymentID               ReturnObjectReferenceNumbersReferenceNumberType = "bankprov_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeBnkDevPrenotificationID         ReturnObjectReferenceNumbersReferenceNumberType = "bnk_dev_prenotification_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeBnkDevTransferID                ReturnObjectReferenceNumbersReferenceNumberType = "bnk_dev_transfer_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeBofaEndToEndID                  ReturnObjectReferenceNumbersReferenceNumberType = "bofa_end_to_end_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeBofaTransactionID               ReturnObjectReferenceNumbersReferenceNumberType = "bofa_transaction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeCheckNumber                     ReturnObjectReferenceNumbersReferenceNumberType = "check_number"
	ReturnObjectReferenceNumbersReferenceNumberTypeColumnFxQuoteID                 ReturnObjectReferenceNumbersReferenceNumberType = "column_fx_quote_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeColumnTransferID                ReturnObjectReferenceNumbersReferenceNumberType = "column_transfer_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeCrossRiverPaymentID             ReturnObjectReferenceNumbersReferenceNumberType = "cross_river_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeCrossRiverTransactionID         ReturnObjectReferenceNumbersReferenceNumberType = "cross_river_transaction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeCurrencycloudConversionID       ReturnObjectReferenceNumbersReferenceNumberType = "currencycloud_conversion_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeCurrencycloudPaymentID          ReturnObjectReferenceNumbersReferenceNumberType = "currencycloud_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeDcBankTransactionID             ReturnObjectReferenceNumbersReferenceNumberType = "dc_bank_transaction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeDwollaTransactionID             ReturnObjectReferenceNumbersReferenceNumberType = "dwolla_transaction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeEftTraceNumber                  ReturnObjectReferenceNumbersReferenceNumberType = "eft_trace_number"
	ReturnObjectReferenceNumbersReferenceNumberTypeFedwireImad                     ReturnObjectReferenceNumbersReferenceNumberType = "fedwire_imad"
	ReturnObjectReferenceNumbersReferenceNumberTypeFedwireOmad                     ReturnObjectReferenceNumbersReferenceNumberType = "fedwire_omad"
	ReturnObjectReferenceNumbersReferenceNumberTypeFirstRepublicInternalID         ReturnObjectReferenceNumbersReferenceNumberType = "first_republic_internal_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsCollectionRequestID ReturnObjectReferenceNumbersReferenceNumberType = "goldman_sachs_collection_request_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsEndToEndID          ReturnObjectReferenceNumbersReferenceNumberType = "goldman_sachs_end_to_end_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsPaymentRequestID    ReturnObjectReferenceNumbersReferenceNumberType = "goldman_sachs_payment_request_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsRequestID           ReturnObjectReferenceNumbersReferenceNumberType = "goldman_sachs_request_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsUniquePaymentID     ReturnObjectReferenceNumbersReferenceNumberType = "goldman_sachs_unique_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeInteracMessageID                ReturnObjectReferenceNumbersReferenceNumberType = "interac_message_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcCcn                         ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_ccn"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcCustomerReferenceID         ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_customer_reference_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcEndToEndID                  ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_end_to_end_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcFirmRootID                  ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_firm_root_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcP3ID                        ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_p3_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcPaymentBatchID              ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_payment_batch_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcPaymentInformationID        ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_payment_information_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeLobCheckID                      ReturnObjectReferenceNumbersReferenceNumberType = "lob_check_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeOther                           ReturnObjectReferenceNumbersReferenceNumberType = "other"
	ReturnObjectReferenceNumbersReferenceNumberTypePartialSwiftMir                 ReturnObjectReferenceNumbersReferenceNumberType = "partial_swift_mir"
	ReturnObjectReferenceNumbersReferenceNumberTypePncClearingReference            ReturnObjectReferenceNumbersReferenceNumberType = "pnc_clearing_reference"
	ReturnObjectReferenceNumbersReferenceNumberTypePncInstructionID                ReturnObjectReferenceNumbersReferenceNumberType = "pnc_instruction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypePncMultipaymentID               ReturnObjectReferenceNumbersReferenceNumberType = "pnc_multipayment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypePncPaymentTraceID               ReturnObjectReferenceNumbersReferenceNumberType = "pnc_payment_trace_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeRtpInstructionID                ReturnObjectReferenceNumbersReferenceNumberType = "rtp_instruction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeSignetAPIReferenceID            ReturnObjectReferenceNumbersReferenceNumberType = "signet_api_reference_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeSignetConfirmationID            ReturnObjectReferenceNumbersReferenceNumberType = "signet_confirmation_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeSignetRequestID                 ReturnObjectReferenceNumbersReferenceNumberType = "signet_request_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeSilvergatePaymentID             ReturnObjectReferenceNumbersReferenceNumberType = "silvergate_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeSwiftMir                        ReturnObjectReferenceNumbersReferenceNumberType = "swift_mir"
	ReturnObjectReferenceNumbersReferenceNumberTypeSwiftUetr                       ReturnObjectReferenceNumbersReferenceNumberType = "swift_uetr"
	ReturnObjectReferenceNumbersReferenceNumberTypeUsbankPaymentID                 ReturnObjectReferenceNumbersReferenceNumberType = "usbank_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeWellsFargoPaymentID             ReturnObjectReferenceNumbersReferenceNumberType = "wells_fargo_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeWellsFargoTraceNumber           ReturnObjectReferenceNumbersReferenceNumberType = "wells_fargo_trace_number"
)
