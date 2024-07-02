// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

// ReturnService contains methods and other services that help with interacting
// with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReturnService] method instead.
type ReturnService struct {
	Options []option.RequestOption
}

// NewReturnService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewReturnService(opts ...option.RequestOption) (r *ReturnService) {
	r = &ReturnService{}
	r.Options = opts
	return
}

// Create a return.
func (r *ReturnService) New(ctx context.Context, body ReturnNewParams, opts ...option.RequestOption) (res *ReturnObject, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/returns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a single return.
func (r *ReturnService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ReturnObject, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/returns/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of returns.
func (r *ReturnService) List(ctx context.Context, query ReturnListParams, opts ...option.RequestOption) (res *pagination.Page[ReturnObject], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/returns"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get a list of returns.
func (r *ReturnService) ListAutoPaging(ctx context.Context, query ReturnListParams, opts ...option.RequestOption) *pagination.PageAutoPager[ReturnObject] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type ReturnObject struct {
	ID string `json:"id,required" format:"uuid"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount int64 `json:"amount,required"`
	// The return code. For ACH returns, this is the required ACH return code.
	Code      ReturnObjectCode `json:"code,required,nullable"`
	CreatedAt time.Time        `json:"created_at,required" format:"date-time"`
	// Currency that this transaction is denominated in.
	Currency shared.Currency `json:"currency,required"`
	// If the return's status is `returned`, this will include the return object's data
	// that is returning this return.
	CurrentReturn *ReturnObject `json:"current_return,required,nullable"`
	// If the return code is `R14` or `R15` this is the date the deceased counterparty
	// passed away.
	DateOfDeath time.Time `json:"date_of_death,required,nullable" format:"date"`
	// If an originating return failed to be processed by the bank, a description of
	// the failure reason will be available.
	FailureReason string `json:"failure_reason,required,nullable"`
	// The ID of the relevant Internal Account.
	InternalAccountID string `json:"internal_account_id,required,nullable" format:"uuid"`
	// The ID of the ledger transaction linked to the return.
	LedgerTransactionID string `json:"ledger_transaction_id,required,nullable" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode,required"`
	Object   string `json:"object,required"`
	// Often the bank will provide an explanation for the return, which is a short
	// human readable string.
	Reason string `json:"reason,required,nullable"`
	// An array of Payment Reference objects.
	ReferenceNumbers []ReturnObjectReferenceNumber `json:"reference_numbers,required"`
	// The ID of the object being returned or `null`.
	ReturnableID string `json:"returnable_id,required,nullable" format:"uuid"`
	// The type of object being returned or `null`.
	ReturnableType ReturnObjectReturnableType `json:"returnable_type,required,nullable"`
	// The role of the return, can be `originating` or `receiving`.
	Role ReturnObjectRole `json:"role,required"`
	// The current status of the return.
	Status ReturnObjectStatus `json:"status,required"`
	// The ID of the relevant Transaction or `null`.
	TransactionID string `json:"transaction_id,required,nullable" format:"uuid"`
	// The ID of the relevant Transaction Line Item or `null`.
	TransactionLineItemID string `json:"transaction_line_item_id,required,nullable" format:"uuid"`
	// The type of return. Can be one of: `ach`, `ach_noc`, `au_becs`, `bacs`, `eft`,
	// `interac`, `manual`, `paper_item`, `wire`.
	Type      ReturnObjectType `json:"type,required"`
	UpdatedAt time.Time        `json:"updated_at,required" format:"date-time"`
	// Some returns may include additional information from the bank. In these cases,
	// this string will be present.
	AdditionalInformation string           `json:"additional_information,nullable"`
	JSON                  returnObjectJSON `json:"-"`
}

// returnObjectJSON contains the JSON metadata for the struct [ReturnObject]
type returnObjectJSON struct {
	ID                    apijson.Field
	Amount                apijson.Field
	Code                  apijson.Field
	CreatedAt             apijson.Field
	Currency              apijson.Field
	CurrentReturn         apijson.Field
	DateOfDeath           apijson.Field
	FailureReason         apijson.Field
	InternalAccountID     apijson.Field
	LedgerTransactionID   apijson.Field
	LiveMode              apijson.Field
	Object                apijson.Field
	Reason                apijson.Field
	ReferenceNumbers      apijson.Field
	ReturnableID          apijson.Field
	ReturnableType        apijson.Field
	Role                  apijson.Field
	Status                apijson.Field
	TransactionID         apijson.Field
	TransactionLineItemID apijson.Field
	Type                  apijson.Field
	UpdatedAt             apijson.Field
	AdditionalInformation apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ReturnObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r returnObjectJSON) RawJSON() string {
	return r.raw
}

// The return code. For ACH returns, this is the required ACH return code.
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
	ReturnObjectCodeC01           ReturnObjectCode = "C01"
	ReturnObjectCodeC02           ReturnObjectCode = "C02"
	ReturnObjectCodeC03           ReturnObjectCode = "C03"
	ReturnObjectCodeC05           ReturnObjectCode = "C05"
	ReturnObjectCodeC06           ReturnObjectCode = "C06"
	ReturnObjectCodeC07           ReturnObjectCode = "C07"
	ReturnObjectCodeC08           ReturnObjectCode = "C08"
	ReturnObjectCodeC09           ReturnObjectCode = "C09"
	ReturnObjectCodeC13           ReturnObjectCode = "C13"
	ReturnObjectCodeC14           ReturnObjectCode = "C14"
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

func (r ReturnObjectCode) IsKnown() bool {
	switch r {
	case ReturnObjectCode901, ReturnObjectCode902, ReturnObjectCode903, ReturnObjectCode904, ReturnObjectCode905, ReturnObjectCode907, ReturnObjectCode908, ReturnObjectCode909, ReturnObjectCode910, ReturnObjectCode911, ReturnObjectCode912, ReturnObjectCode914, ReturnObjectCodeC01, ReturnObjectCodeC02, ReturnObjectCodeC03, ReturnObjectCodeC05, ReturnObjectCodeC06, ReturnObjectCodeC07, ReturnObjectCodeC08, ReturnObjectCodeC09, ReturnObjectCodeC13, ReturnObjectCodeC14, ReturnObjectCodeR01, ReturnObjectCodeR02, ReturnObjectCodeR03, ReturnObjectCodeR04, ReturnObjectCodeR05, ReturnObjectCodeR06, ReturnObjectCodeR07, ReturnObjectCodeR08, ReturnObjectCodeR09, ReturnObjectCodeR10, ReturnObjectCodeR11, ReturnObjectCodeR12, ReturnObjectCodeR14, ReturnObjectCodeR15, ReturnObjectCodeR16, ReturnObjectCodeR17, ReturnObjectCodeR20, ReturnObjectCodeR21, ReturnObjectCodeR22, ReturnObjectCodeR23, ReturnObjectCodeR24, ReturnObjectCodeR29, ReturnObjectCodeR31, ReturnObjectCodeR33, ReturnObjectCodeR37, ReturnObjectCodeR38, ReturnObjectCodeR39, ReturnObjectCodeR51, ReturnObjectCodeR52, ReturnObjectCodeR53, ReturnObjectCodeCurrencycloud:
		return true
	}
	return false
}

type ReturnObjectReferenceNumber struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode,required"`
	Object   string `json:"object,required"`
	// The vendor reference number.
	ReferenceNumber string `json:"reference_number,required"`
	// The type of the reference number. Referring to the vendor payment id.
	ReferenceNumberType ReturnObjectReferenceNumbersReferenceNumberType `json:"reference_number_type,required"`
	UpdatedAt           time.Time                                       `json:"updated_at,required" format:"date-time"`
	JSON                returnObjectReferenceNumberJSON                 `json:"-"`
}

// returnObjectReferenceNumberJSON contains the JSON metadata for the struct
// [ReturnObjectReferenceNumber]
type returnObjectReferenceNumberJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	LiveMode            apijson.Field
	Object              apijson.Field
	ReferenceNumber     apijson.Field
	ReferenceNumberType apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ReturnObjectReferenceNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r returnObjectReferenceNumberJSON) RawJSON() string {
	return r.raw
}

// The type of the reference number. Referring to the vendor payment id.
type ReturnObjectReferenceNumbersReferenceNumberType string

const (
	ReturnObjectReferenceNumbersReferenceNumberTypeACHOriginalTraceNumber                         ReturnObjectReferenceNumbersReferenceNumberType = "ach_original_trace_number"
	ReturnObjectReferenceNumbersReferenceNumberTypeACHTraceNumber                                 ReturnObjectReferenceNumbersReferenceNumberType = "ach_trace_number"
	ReturnObjectReferenceNumbersReferenceNumberTypeBankprovPaymentActivityDate                    ReturnObjectReferenceNumbersReferenceNumberType = "bankprov_payment_activity_date"
	ReturnObjectReferenceNumbersReferenceNumberTypeBankprovPaymentID                              ReturnObjectReferenceNumbersReferenceNumberType = "bankprov_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeBnkDevPrenotificationID                        ReturnObjectReferenceNumbersReferenceNumberType = "bnk_dev_prenotification_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeBnkDevTransferID                               ReturnObjectReferenceNumbersReferenceNumberType = "bnk_dev_transfer_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeBofaEndToEndID                                 ReturnObjectReferenceNumbersReferenceNumberType = "bofa_end_to_end_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeBofaTransactionID                              ReturnObjectReferenceNumbersReferenceNumberType = "bofa_transaction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeCheckNumber                                    ReturnObjectReferenceNumbersReferenceNumberType = "check_number"
	ReturnObjectReferenceNumbersReferenceNumberTypeCitibankReferenceNumber                        ReturnObjectReferenceNumbersReferenceNumberType = "citibank_reference_number"
	ReturnObjectReferenceNumbersReferenceNumberTypeCitibankWorldlinkClearingSystemReferenceNumber ReturnObjectReferenceNumbersReferenceNumberType = "citibank_worldlink_clearing_system_reference_number"
	ReturnObjectReferenceNumbersReferenceNumberTypeColumnFxQuoteID                                ReturnObjectReferenceNumbersReferenceNumberType = "column_fx_quote_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeColumnReversalPairTransferID                   ReturnObjectReferenceNumbersReferenceNumberType = "column_reversal_pair_transfer_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeColumnTransferID                               ReturnObjectReferenceNumbersReferenceNumberType = "column_transfer_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeCrossRiverPaymentID                            ReturnObjectReferenceNumbersReferenceNumberType = "cross_river_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeCrossRiverServiceMessage                       ReturnObjectReferenceNumbersReferenceNumberType = "cross_river_service_message"
	ReturnObjectReferenceNumbersReferenceNumberTypeCrossRiverTransactionID                        ReturnObjectReferenceNumbersReferenceNumberType = "cross_river_transaction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeCurrencycloudConversionID                      ReturnObjectReferenceNumbersReferenceNumberType = "currencycloud_conversion_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeCurrencycloudPaymentID                         ReturnObjectReferenceNumbersReferenceNumberType = "currencycloud_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeDcBankTransactionID                            ReturnObjectReferenceNumbersReferenceNumberType = "dc_bank_transaction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeDwollaTransactionID                            ReturnObjectReferenceNumbersReferenceNumberType = "dwolla_transaction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeEftTraceNumber                                 ReturnObjectReferenceNumbersReferenceNumberType = "eft_trace_number"
	ReturnObjectReferenceNumbersReferenceNumberTypeEvolveTransactionID                            ReturnObjectReferenceNumbersReferenceNumberType = "evolve_transaction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeFedwireImad                                    ReturnObjectReferenceNumbersReferenceNumberType = "fedwire_imad"
	ReturnObjectReferenceNumbersReferenceNumberTypeFedwireOmad                                    ReturnObjectReferenceNumbersReferenceNumberType = "fedwire_omad"
	ReturnObjectReferenceNumbersReferenceNumberTypeFirstRepublicInternalID                        ReturnObjectReferenceNumbersReferenceNumberType = "first_republic_internal_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsCollectionRequestID                ReturnObjectReferenceNumbersReferenceNumberType = "goldman_sachs_collection_request_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsEndToEndID                         ReturnObjectReferenceNumbersReferenceNumberType = "goldman_sachs_end_to_end_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsPaymentRequestID                   ReturnObjectReferenceNumbersReferenceNumberType = "goldman_sachs_payment_request_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsRequestID                          ReturnObjectReferenceNumbersReferenceNumberType = "goldman_sachs_request_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsUniquePaymentID                    ReturnObjectReferenceNumbersReferenceNumberType = "goldman_sachs_unique_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeInteracMessageID                               ReturnObjectReferenceNumbersReferenceNumberType = "interac_message_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcCcn                                        ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_ccn"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcClearingSystemReference                    ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_clearing_system_reference"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcCustomerReferenceID                        ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_customer_reference_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcEndToEndID                                 ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_end_to_end_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcFirmRootID                                 ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_firm_root_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcP3ID                                       ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_p3_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcPaymentBatchID                             ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_payment_batch_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcPaymentInformationID                       ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_payment_information_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeJpmcPaymentReturnedDatetime                    ReturnObjectReferenceNumbersReferenceNumberType = "jpmc_payment_returned_datetime"
	ReturnObjectReferenceNumbersReferenceNumberTypeLobCheckID                                     ReturnObjectReferenceNumbersReferenceNumberType = "lob_check_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeOther                                          ReturnObjectReferenceNumbersReferenceNumberType = "other"
	ReturnObjectReferenceNumbersReferenceNumberTypePartialSwiftMir                                ReturnObjectReferenceNumbersReferenceNumberType = "partial_swift_mir"
	ReturnObjectReferenceNumbersReferenceNumberTypePncClearingReference                           ReturnObjectReferenceNumbersReferenceNumberType = "pnc_clearing_reference"
	ReturnObjectReferenceNumbersReferenceNumberTypePncInstructionID                               ReturnObjectReferenceNumbersReferenceNumberType = "pnc_instruction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypePncMultipaymentID                              ReturnObjectReferenceNumbersReferenceNumberType = "pnc_multipayment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypePncPaymentTraceID                              ReturnObjectReferenceNumbersReferenceNumberType = "pnc_payment_trace_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeRspecVendorPaymentID                           ReturnObjectReferenceNumbersReferenceNumberType = "rspec_vendor_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeRtpInstructionID                               ReturnObjectReferenceNumbersReferenceNumberType = "rtp_instruction_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeSignetAPIReferenceID                           ReturnObjectReferenceNumbersReferenceNumberType = "signet_api_reference_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeSignetConfirmationID                           ReturnObjectReferenceNumbersReferenceNumberType = "signet_confirmation_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeSignetRequestID                                ReturnObjectReferenceNumbersReferenceNumberType = "signet_request_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeSilvergatePaymentID                            ReturnObjectReferenceNumbersReferenceNumberType = "silvergate_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeSvbEndToEndID                                  ReturnObjectReferenceNumbersReferenceNumberType = "svb_end_to_end_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeSvbPaymentID                                   ReturnObjectReferenceNumbersReferenceNumberType = "svb_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeSvbTransactionClearedForSanctionsReview        ReturnObjectReferenceNumbersReferenceNumberType = "svb_transaction_cleared_for_sanctions_review"
	ReturnObjectReferenceNumbersReferenceNumberTypeSvbTransactionHeldForSanctionsReview           ReturnObjectReferenceNumbersReferenceNumberType = "svb_transaction_held_for_sanctions_review"
	ReturnObjectReferenceNumbersReferenceNumberTypeSwiftMir                                       ReturnObjectReferenceNumbersReferenceNumberType = "swift_mir"
	ReturnObjectReferenceNumbersReferenceNumberTypeSwiftUetr                                      ReturnObjectReferenceNumbersReferenceNumberType = "swift_uetr"
	ReturnObjectReferenceNumbersReferenceNumberTypeUmbProductPartnerAccountNumber                 ReturnObjectReferenceNumbersReferenceNumberType = "umb_product_partner_account_number"
	ReturnObjectReferenceNumbersReferenceNumberTypeUsbankPaymentID                                ReturnObjectReferenceNumbersReferenceNumberType = "usbank_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeWellsFargoPaymentID                            ReturnObjectReferenceNumbersReferenceNumberType = "wells_fargo_payment_id"
	ReturnObjectReferenceNumbersReferenceNumberTypeWellsFargoTraceNumber                          ReturnObjectReferenceNumbersReferenceNumberType = "wells_fargo_trace_number"
)

func (r ReturnObjectReferenceNumbersReferenceNumberType) IsKnown() bool {
	switch r {
	case ReturnObjectReferenceNumbersReferenceNumberTypeACHOriginalTraceNumber, ReturnObjectReferenceNumbersReferenceNumberTypeACHTraceNumber, ReturnObjectReferenceNumbersReferenceNumberTypeBankprovPaymentActivityDate, ReturnObjectReferenceNumbersReferenceNumberTypeBankprovPaymentID, ReturnObjectReferenceNumbersReferenceNumberTypeBnkDevPrenotificationID, ReturnObjectReferenceNumbersReferenceNumberTypeBnkDevTransferID, ReturnObjectReferenceNumbersReferenceNumberTypeBofaEndToEndID, ReturnObjectReferenceNumbersReferenceNumberTypeBofaTransactionID, ReturnObjectReferenceNumbersReferenceNumberTypeCheckNumber, ReturnObjectReferenceNumbersReferenceNumberTypeCitibankReferenceNumber, ReturnObjectReferenceNumbersReferenceNumberTypeCitibankWorldlinkClearingSystemReferenceNumber, ReturnObjectReferenceNumbersReferenceNumberTypeColumnFxQuoteID, ReturnObjectReferenceNumbersReferenceNumberTypeColumnReversalPairTransferID, ReturnObjectReferenceNumbersReferenceNumberTypeColumnTransferID, ReturnObjectReferenceNumbersReferenceNumberTypeCrossRiverPaymentID, ReturnObjectReferenceNumbersReferenceNumberTypeCrossRiverServiceMessage, ReturnObjectReferenceNumbersReferenceNumberTypeCrossRiverTransactionID, ReturnObjectReferenceNumbersReferenceNumberTypeCurrencycloudConversionID, ReturnObjectReferenceNumbersReferenceNumberTypeCurrencycloudPaymentID, ReturnObjectReferenceNumbersReferenceNumberTypeDcBankTransactionID, ReturnObjectReferenceNumbersReferenceNumberTypeDwollaTransactionID, ReturnObjectReferenceNumbersReferenceNumberTypeEftTraceNumber, ReturnObjectReferenceNumbersReferenceNumberTypeEvolveTransactionID, ReturnObjectReferenceNumbersReferenceNumberTypeFedwireImad, ReturnObjectReferenceNumbersReferenceNumberTypeFedwireOmad, ReturnObjectReferenceNumbersReferenceNumberTypeFirstRepublicInternalID, ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsCollectionRequestID, ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsEndToEndID, ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsPaymentRequestID, ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsRequestID, ReturnObjectReferenceNumbersReferenceNumberTypeGoldmanSachsUniquePaymentID, ReturnObjectReferenceNumbersReferenceNumberTypeInteracMessageID, ReturnObjectReferenceNumbersReferenceNumberTypeJpmcCcn, ReturnObjectReferenceNumbersReferenceNumberTypeJpmcClearingSystemReference, ReturnObjectReferenceNumbersReferenceNumberTypeJpmcCustomerReferenceID, ReturnObjectReferenceNumbersReferenceNumberTypeJpmcEndToEndID, ReturnObjectReferenceNumbersReferenceNumberTypeJpmcFirmRootID, ReturnObjectReferenceNumbersReferenceNumberTypeJpmcP3ID, ReturnObjectReferenceNumbersReferenceNumberTypeJpmcPaymentBatchID, ReturnObjectReferenceNumbersReferenceNumberTypeJpmcPaymentInformationID, ReturnObjectReferenceNumbersReferenceNumberTypeJpmcPaymentReturnedDatetime, ReturnObjectReferenceNumbersReferenceNumberTypeLobCheckID, ReturnObjectReferenceNumbersReferenceNumberTypeOther, ReturnObjectReferenceNumbersReferenceNumberTypePartialSwiftMir, ReturnObjectReferenceNumbersReferenceNumberTypePncClearingReference, ReturnObjectReferenceNumbersReferenceNumberTypePncInstructionID, ReturnObjectReferenceNumbersReferenceNumberTypePncMultipaymentID, ReturnObjectReferenceNumbersReferenceNumberTypePncPaymentTraceID, ReturnObjectReferenceNumbersReferenceNumberTypeRspecVendorPaymentID, ReturnObjectReferenceNumbersReferenceNumberTypeRtpInstructionID, ReturnObjectReferenceNumbersReferenceNumberTypeSignetAPIReferenceID, ReturnObjectReferenceNumbersReferenceNumberTypeSignetConfirmationID, ReturnObjectReferenceNumbersReferenceNumberTypeSignetRequestID, ReturnObjectReferenceNumbersReferenceNumberTypeSilvergatePaymentID, ReturnObjectReferenceNumbersReferenceNumberTypeSvbEndToEndID, ReturnObjectReferenceNumbersReferenceNumberTypeSvbPaymentID, ReturnObjectReferenceNumbersReferenceNumberTypeSvbTransactionClearedForSanctionsReview, ReturnObjectReferenceNumbersReferenceNumberTypeSvbTransactionHeldForSanctionsReview, ReturnObjectReferenceNumbersReferenceNumberTypeSwiftMir, ReturnObjectReferenceNumbersReferenceNumberTypeSwiftUetr, ReturnObjectReferenceNumbersReferenceNumberTypeUmbProductPartnerAccountNumber, ReturnObjectReferenceNumbersReferenceNumberTypeUsbankPaymentID, ReturnObjectReferenceNumbersReferenceNumberTypeWellsFargoPaymentID, ReturnObjectReferenceNumbersReferenceNumberTypeWellsFargoTraceNumber:
		return true
	}
	return false
}

// The type of object being returned or `null`.
type ReturnObjectReturnableType string

const (
	ReturnObjectReturnableTypeIncomingPaymentDetail ReturnObjectReturnableType = "incoming_payment_detail"
	ReturnObjectReturnableTypePaperItem             ReturnObjectReturnableType = "paper_item"
	ReturnObjectReturnableTypePaymentOrder          ReturnObjectReturnableType = "payment_order"
	ReturnObjectReturnableTypeReturn                ReturnObjectReturnableType = "return"
	ReturnObjectReturnableTypeReversal              ReturnObjectReturnableType = "reversal"
)

func (r ReturnObjectReturnableType) IsKnown() bool {
	switch r {
	case ReturnObjectReturnableTypeIncomingPaymentDetail, ReturnObjectReturnableTypePaperItem, ReturnObjectReturnableTypePaymentOrder, ReturnObjectReturnableTypeReturn, ReturnObjectReturnableTypeReversal:
		return true
	}
	return false
}

// The role of the return, can be `originating` or `receiving`.
type ReturnObjectRole string

const (
	ReturnObjectRoleOriginating ReturnObjectRole = "originating"
	ReturnObjectRoleReceiving   ReturnObjectRole = "receiving"
)

func (r ReturnObjectRole) IsKnown() bool {
	switch r {
	case ReturnObjectRoleOriginating, ReturnObjectRoleReceiving:
		return true
	}
	return false
}

// The current status of the return.
type ReturnObjectStatus string

const (
	ReturnObjectStatusCompleted  ReturnObjectStatus = "completed"
	ReturnObjectStatusFailed     ReturnObjectStatus = "failed"
	ReturnObjectStatusPending    ReturnObjectStatus = "pending"
	ReturnObjectStatusProcessing ReturnObjectStatus = "processing"
	ReturnObjectStatusReturned   ReturnObjectStatus = "returned"
	ReturnObjectStatusSent       ReturnObjectStatus = "sent"
)

func (r ReturnObjectStatus) IsKnown() bool {
	switch r {
	case ReturnObjectStatusCompleted, ReturnObjectStatusFailed, ReturnObjectStatusPending, ReturnObjectStatusProcessing, ReturnObjectStatusReturned, ReturnObjectStatusSent:
		return true
	}
	return false
}

// The type of return. Can be one of: `ach`, `ach_noc`, `au_becs`, `bacs`, `eft`,
// `interac`, `manual`, `paper_item`, `wire`.
type ReturnObjectType string

const (
	ReturnObjectTypeACH         ReturnObjectType = "ach"
	ReturnObjectTypeACHNoc      ReturnObjectType = "ach_noc"
	ReturnObjectTypeAuBecs      ReturnObjectType = "au_becs"
	ReturnObjectTypeBacs        ReturnObjectType = "bacs"
	ReturnObjectTypeBook        ReturnObjectType = "book"
	ReturnObjectTypeCheck       ReturnObjectType = "check"
	ReturnObjectTypeCrossBorder ReturnObjectType = "cross_border"
	ReturnObjectTypeEft         ReturnObjectType = "eft"
	ReturnObjectTypeInterac     ReturnObjectType = "interac"
	ReturnObjectTypeManual      ReturnObjectType = "manual"
	ReturnObjectTypePaperItem   ReturnObjectType = "paper_item"
	ReturnObjectTypeSepa        ReturnObjectType = "sepa"
	ReturnObjectTypeWire        ReturnObjectType = "wire"
)

func (r ReturnObjectType) IsKnown() bool {
	switch r {
	case ReturnObjectTypeACH, ReturnObjectTypeACHNoc, ReturnObjectTypeAuBecs, ReturnObjectTypeBacs, ReturnObjectTypeBook, ReturnObjectTypeCheck, ReturnObjectTypeCrossBorder, ReturnObjectTypeEft, ReturnObjectTypeInterac, ReturnObjectTypeManual, ReturnObjectTypePaperItem, ReturnObjectTypeSepa, ReturnObjectTypeWire:
		return true
	}
	return false
}

type ReturnNewParams struct {
	// The ID of the object being returned or `null`.
	ReturnableID param.Field[string] `json:"returnable_id,required" format:"uuid"`
	// The type of object being returned. Currently, this may only be
	// incoming_payment_detail.
	ReturnableType param.Field[ReturnNewParamsReturnableType] `json:"returnable_type,required"`
	// Some returns may include additional information from the bank. In these cases,
	// this string will be present.
	AdditionalInformation param.Field[string] `json:"additional_information"`
	// The return code. For ACH returns, this is the required ACH return code.
	Code param.Field[ReturnNewParamsCode] `json:"code"`
	// If the return code is `R14` or `R15` this is the date the deceased counterparty
	// passed away.
	DateOfDeath param.Field[time.Time] `json:"date_of_death" format:"date"`
	// An optional description of the reason for the return. This is for internal usage
	// and will not be transmitted to the bank.”
	Reason param.Field[string] `json:"reason"`
}

func (r ReturnNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of object being returned. Currently, this may only be
// incoming_payment_detail.
type ReturnNewParamsReturnableType string

const (
	ReturnNewParamsReturnableTypeIncomingPaymentDetail ReturnNewParamsReturnableType = "incoming_payment_detail"
)

func (r ReturnNewParamsReturnableType) IsKnown() bool {
	switch r {
	case ReturnNewParamsReturnableTypeIncomingPaymentDetail:
		return true
	}
	return false
}

// The return code. For ACH returns, this is the required ACH return code.
type ReturnNewParamsCode string

const (
	ReturnNewParamsCode901           ReturnNewParamsCode = "901"
	ReturnNewParamsCode902           ReturnNewParamsCode = "902"
	ReturnNewParamsCode903           ReturnNewParamsCode = "903"
	ReturnNewParamsCode904           ReturnNewParamsCode = "904"
	ReturnNewParamsCode905           ReturnNewParamsCode = "905"
	ReturnNewParamsCode907           ReturnNewParamsCode = "907"
	ReturnNewParamsCode908           ReturnNewParamsCode = "908"
	ReturnNewParamsCode909           ReturnNewParamsCode = "909"
	ReturnNewParamsCode910           ReturnNewParamsCode = "910"
	ReturnNewParamsCode911           ReturnNewParamsCode = "911"
	ReturnNewParamsCode912           ReturnNewParamsCode = "912"
	ReturnNewParamsCode914           ReturnNewParamsCode = "914"
	ReturnNewParamsCodeC01           ReturnNewParamsCode = "C01"
	ReturnNewParamsCodeC02           ReturnNewParamsCode = "C02"
	ReturnNewParamsCodeC03           ReturnNewParamsCode = "C03"
	ReturnNewParamsCodeC05           ReturnNewParamsCode = "C05"
	ReturnNewParamsCodeC06           ReturnNewParamsCode = "C06"
	ReturnNewParamsCodeC07           ReturnNewParamsCode = "C07"
	ReturnNewParamsCodeC08           ReturnNewParamsCode = "C08"
	ReturnNewParamsCodeC09           ReturnNewParamsCode = "C09"
	ReturnNewParamsCodeC13           ReturnNewParamsCode = "C13"
	ReturnNewParamsCodeC14           ReturnNewParamsCode = "C14"
	ReturnNewParamsCodeR01           ReturnNewParamsCode = "R01"
	ReturnNewParamsCodeR02           ReturnNewParamsCode = "R02"
	ReturnNewParamsCodeR03           ReturnNewParamsCode = "R03"
	ReturnNewParamsCodeR04           ReturnNewParamsCode = "R04"
	ReturnNewParamsCodeR05           ReturnNewParamsCode = "R05"
	ReturnNewParamsCodeR06           ReturnNewParamsCode = "R06"
	ReturnNewParamsCodeR07           ReturnNewParamsCode = "R07"
	ReturnNewParamsCodeR08           ReturnNewParamsCode = "R08"
	ReturnNewParamsCodeR09           ReturnNewParamsCode = "R09"
	ReturnNewParamsCodeR10           ReturnNewParamsCode = "R10"
	ReturnNewParamsCodeR11           ReturnNewParamsCode = "R11"
	ReturnNewParamsCodeR12           ReturnNewParamsCode = "R12"
	ReturnNewParamsCodeR14           ReturnNewParamsCode = "R14"
	ReturnNewParamsCodeR15           ReturnNewParamsCode = "R15"
	ReturnNewParamsCodeR16           ReturnNewParamsCode = "R16"
	ReturnNewParamsCodeR17           ReturnNewParamsCode = "R17"
	ReturnNewParamsCodeR20           ReturnNewParamsCode = "R20"
	ReturnNewParamsCodeR21           ReturnNewParamsCode = "R21"
	ReturnNewParamsCodeR22           ReturnNewParamsCode = "R22"
	ReturnNewParamsCodeR23           ReturnNewParamsCode = "R23"
	ReturnNewParamsCodeR24           ReturnNewParamsCode = "R24"
	ReturnNewParamsCodeR29           ReturnNewParamsCode = "R29"
	ReturnNewParamsCodeR31           ReturnNewParamsCode = "R31"
	ReturnNewParamsCodeR33           ReturnNewParamsCode = "R33"
	ReturnNewParamsCodeR37           ReturnNewParamsCode = "R37"
	ReturnNewParamsCodeR38           ReturnNewParamsCode = "R38"
	ReturnNewParamsCodeR39           ReturnNewParamsCode = "R39"
	ReturnNewParamsCodeR51           ReturnNewParamsCode = "R51"
	ReturnNewParamsCodeR52           ReturnNewParamsCode = "R52"
	ReturnNewParamsCodeR53           ReturnNewParamsCode = "R53"
	ReturnNewParamsCodeCurrencycloud ReturnNewParamsCode = "currencycloud"
)

func (r ReturnNewParamsCode) IsKnown() bool {
	switch r {
	case ReturnNewParamsCode901, ReturnNewParamsCode902, ReturnNewParamsCode903, ReturnNewParamsCode904, ReturnNewParamsCode905, ReturnNewParamsCode907, ReturnNewParamsCode908, ReturnNewParamsCode909, ReturnNewParamsCode910, ReturnNewParamsCode911, ReturnNewParamsCode912, ReturnNewParamsCode914, ReturnNewParamsCodeC01, ReturnNewParamsCodeC02, ReturnNewParamsCodeC03, ReturnNewParamsCodeC05, ReturnNewParamsCodeC06, ReturnNewParamsCodeC07, ReturnNewParamsCodeC08, ReturnNewParamsCodeC09, ReturnNewParamsCodeC13, ReturnNewParamsCodeC14, ReturnNewParamsCodeR01, ReturnNewParamsCodeR02, ReturnNewParamsCodeR03, ReturnNewParamsCodeR04, ReturnNewParamsCodeR05, ReturnNewParamsCodeR06, ReturnNewParamsCodeR07, ReturnNewParamsCodeR08, ReturnNewParamsCodeR09, ReturnNewParamsCodeR10, ReturnNewParamsCodeR11, ReturnNewParamsCodeR12, ReturnNewParamsCodeR14, ReturnNewParamsCodeR15, ReturnNewParamsCodeR16, ReturnNewParamsCodeR17, ReturnNewParamsCodeR20, ReturnNewParamsCodeR21, ReturnNewParamsCodeR22, ReturnNewParamsCodeR23, ReturnNewParamsCodeR24, ReturnNewParamsCodeR29, ReturnNewParamsCodeR31, ReturnNewParamsCodeR33, ReturnNewParamsCodeR37, ReturnNewParamsCodeR38, ReturnNewParamsCodeR39, ReturnNewParamsCodeR51, ReturnNewParamsCodeR52, ReturnNewParamsCodeR53, ReturnNewParamsCodeCurrencycloud:
		return true
	}
	return false
}

type ReturnListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// Specify `counterparty_id` if you wish to see returns that occurred with a
	// specific counterparty.
	CounterpartyID param.Field[string] `query:"counterparty_id"`
	// Specify `internal_account_id` if you wish to see returns to/from a specific
	// account.
	InternalAccountID param.Field[string] `query:"internal_account_id"`
	PerPage           param.Field[int64]  `query:"per_page"`
	// The ID of a valid returnable. Must be accompanied by `returnable_type`.
	ReturnableID param.Field[string] `query:"returnable_id"`
	// One of `payment_order`, `paper_item`, `reversal`, or `incoming_payment_detail`.
	// Must be accompanied by `returnable_id`.
	ReturnableType param.Field[ReturnListParamsReturnableType] `query:"returnable_type"`
}

// URLQuery serializes [ReturnListParams]'s query parameters as `url.Values`.
func (r ReturnListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// One of `payment_order`, `paper_item`, `reversal`, or `incoming_payment_detail`.
// Must be accompanied by `returnable_id`.
type ReturnListParamsReturnableType string

const (
	ReturnListParamsReturnableTypeIncomingPaymentDetail ReturnListParamsReturnableType = "incoming_payment_detail"
	ReturnListParamsReturnableTypePaperItem             ReturnListParamsReturnableType = "paper_item"
	ReturnListParamsReturnableTypePaymentOrder          ReturnListParamsReturnableType = "payment_order"
	ReturnListParamsReturnableTypeReturn                ReturnListParamsReturnableType = "return"
	ReturnListParamsReturnableTypeReversal              ReturnListParamsReturnableType = "reversal"
)

func (r ReturnListParamsReturnableType) IsKnown() bool {
	switch r {
	case ReturnListParamsReturnableTypeIncomingPaymentDetail, ReturnListParamsReturnableTypePaperItem, ReturnListParamsReturnableTypePaymentOrder, ReturnListParamsReturnableTypeReturn, ReturnListParamsReturnableTypeReversal:
		return true
	}
	return false
}
