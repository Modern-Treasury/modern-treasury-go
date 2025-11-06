// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/packages/pagination"
)

// PaymentReferenceService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentReferenceService] method instead.
type PaymentReferenceService struct {
	Options []option.RequestOption
}

// NewPaymentReferenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPaymentReferenceService(opts ...option.RequestOption) (r *PaymentReferenceService) {
	r = &PaymentReferenceService{}
	r.Options = opts
	return
}

// get payment_reference
func (r *PaymentReferenceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PaymentReference, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/payment_references/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// list payment_references
func (r *PaymentReferenceService) List(ctx context.Context, query PaymentReferenceListParams, opts ...option.RequestOption) (res *pagination.Page[PaymentReference], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/payment_references"
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

// list payment_references
func (r *PaymentReferenceService) ListAutoPaging(ctx context.Context, query PaymentReferenceListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PaymentReference] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// get payment_reference
//
// Deprecated: use `Get` instead
func (r *PaymentReferenceService) Retireve(ctx context.Context, id string, opts ...option.RequestOption) (res *PaymentReference, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/payment_references/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PaymentReference struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode,required"`
	Object   string `json:"object,required"`
	// The actual reference number assigned by the bank.
	ReferenceNumber string `json:"reference_number,required"`
	// The type of reference number.
	ReferenceNumberType PaymentReferenceReferenceNumberType `json:"reference_number_type,required"`
	// The id of the referenceable to search for. Must be accompanied by the
	// referenceable_type or will return an error.
	ReferenceableID string `json:"referenceable_id,required"`
	// One of the referenceable types. This must be accompanied by the id of the
	// referenceable or will return an error.
	ReferenceableType PaymentReferenceReferenceableType `json:"referenceable_type,required"`
	UpdatedAt         time.Time                         `json:"updated_at,required" format:"date-time"`
	JSON              paymentReferenceJSON              `json:"-"`
}

// paymentReferenceJSON contains the JSON metadata for the struct
// [PaymentReference]
type paymentReferenceJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	LiveMode            apijson.Field
	Object              apijson.Field
	ReferenceNumber     apijson.Field
	ReferenceNumberType apijson.Field
	ReferenceableID     apijson.Field
	ReferenceableType   apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PaymentReference) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentReferenceJSON) RawJSON() string {
	return r.raw
}

// The type of reference number.
type PaymentReferenceReferenceNumberType string

const (
	PaymentReferenceReferenceNumberTypeACHOriginalTraceNumber                         PaymentReferenceReferenceNumberType = "ach_original_trace_number"
	PaymentReferenceReferenceNumberTypeACHTraceNumber                                 PaymentReferenceReferenceNumberType = "ach_trace_number"
	PaymentReferenceReferenceNumberTypeBankprovPaymentActivityDate                    PaymentReferenceReferenceNumberType = "bankprov_payment_activity_date"
	PaymentReferenceReferenceNumberTypeBankprovPaymentID                              PaymentReferenceReferenceNumberType = "bankprov_payment_id"
	PaymentReferenceReferenceNumberTypeBnkDevPrenotificationID                        PaymentReferenceReferenceNumberType = "bnk_dev_prenotification_id"
	PaymentReferenceReferenceNumberTypeBnkDevTransferID                               PaymentReferenceReferenceNumberType = "bnk_dev_transfer_id"
	PaymentReferenceReferenceNumberTypeBnyMellonTransactionReferenceNumber            PaymentReferenceReferenceNumberType = "bny_mellon_transaction_reference_number"
	PaymentReferenceReferenceNumberTypeBofaEndToEndID                                 PaymentReferenceReferenceNumberType = "bofa_end_to_end_id"
	PaymentReferenceReferenceNumberTypeBofaTransactionID                              PaymentReferenceReferenceNumberType = "bofa_transaction_id"
	PaymentReferenceReferenceNumberTypeBraleTransferID                                PaymentReferenceReferenceNumberType = "brale_transfer_id"
	PaymentReferenceReferenceNumberTypeCheckNumber                                    PaymentReferenceReferenceNumberType = "check_number"
	PaymentReferenceReferenceNumberTypeCitibankReferenceNumber                        PaymentReferenceReferenceNumberType = "citibank_reference_number"
	PaymentReferenceReferenceNumberTypeCitibankWorldlinkClearingSystemReferenceNumber PaymentReferenceReferenceNumberType = "citibank_worldlink_clearing_system_reference_number"
	PaymentReferenceReferenceNumberTypeColumnFxQuoteID                                PaymentReferenceReferenceNumberType = "column_fx_quote_id"
	PaymentReferenceReferenceNumberTypeColumnReversalPairTransferID                   PaymentReferenceReferenceNumberType = "column_reversal_pair_transfer_id"
	PaymentReferenceReferenceNumberTypeColumnTransferID                               PaymentReferenceReferenceNumberType = "column_transfer_id"
	PaymentReferenceReferenceNumberTypeCrossRiverCoreTransactionID                    PaymentReferenceReferenceNumberType = "cross_river_core_transaction_id"
	PaymentReferenceReferenceNumberTypeCrossRiverFedBatchID                           PaymentReferenceReferenceNumberType = "cross_river_fed_batch_id"
	PaymentReferenceReferenceNumberTypeCrossRiverPaymentID                            PaymentReferenceReferenceNumberType = "cross_river_payment_id"
	PaymentReferenceReferenceNumberTypeCrossRiverServiceMessage                       PaymentReferenceReferenceNumberType = "cross_river_service_message"
	PaymentReferenceReferenceNumberTypeCrossRiverTransactionID                        PaymentReferenceReferenceNumberType = "cross_river_transaction_id"
	PaymentReferenceReferenceNumberTypeCurrencycloudConversionID                      PaymentReferenceReferenceNumberType = "currencycloud_conversion_id"
	PaymentReferenceReferenceNumberTypeCurrencycloudPaymentID                         PaymentReferenceReferenceNumberType = "currencycloud_payment_id"
	PaymentReferenceReferenceNumberTypeDcBankTransactionID                            PaymentReferenceReferenceNumberType = "dc_bank_transaction_id"
	PaymentReferenceReferenceNumberTypeEftTraceNumber                                 PaymentReferenceReferenceNumberType = "eft_trace_number"
	PaymentReferenceReferenceNumberTypeEvolveCoreBatch                                PaymentReferenceReferenceNumberType = "evolve_core_batch"
	PaymentReferenceReferenceNumberTypeEvolveCoreFileKey                              PaymentReferenceReferenceNumberType = "evolve_core_file_key"
	PaymentReferenceReferenceNumberTypeEvolveCoreSeq                                  PaymentReferenceReferenceNumberType = "evolve_core_seq"
	PaymentReferenceReferenceNumberTypeEvolveTransactionID                            PaymentReferenceReferenceNumberType = "evolve_transaction_id"
	PaymentReferenceReferenceNumberTypeFakeVendorPaymentID                            PaymentReferenceReferenceNumberType = "fake_vendor_payment_id"
	PaymentReferenceReferenceNumberTypeFedwireImad                                    PaymentReferenceReferenceNumberType = "fedwire_imad"
	PaymentReferenceReferenceNumberTypeFedwireOmad                                    PaymentReferenceReferenceNumberType = "fedwire_omad"
	PaymentReferenceReferenceNumberTypeFirstRepublicInternalID                        PaymentReferenceReferenceNumberType = "first_republic_internal_id"
	PaymentReferenceReferenceNumberTypeGoldmanSachsCollectionRequestID                PaymentReferenceReferenceNumberType = "goldman_sachs_collection_request_id"
	PaymentReferenceReferenceNumberTypeGoldmanSachsEndToEndID                         PaymentReferenceReferenceNumberType = "goldman_sachs_end_to_end_id"
	PaymentReferenceReferenceNumberTypeGoldmanSachsPaymentRequestID                   PaymentReferenceReferenceNumberType = "goldman_sachs_payment_request_id"
	PaymentReferenceReferenceNumberTypeGoldmanSachsRequestID                          PaymentReferenceReferenceNumberType = "goldman_sachs_request_id"
	PaymentReferenceReferenceNumberTypeGoldmanSachsUniquePaymentID                    PaymentReferenceReferenceNumberType = "goldman_sachs_unique_payment_id"
	PaymentReferenceReferenceNumberTypeHifiOfframpID                                  PaymentReferenceReferenceNumberType = "hifi_offramp_id"
	PaymentReferenceReferenceNumberTypeHifiTransferID                                 PaymentReferenceReferenceNumberType = "hifi_transfer_id"
	PaymentReferenceReferenceNumberTypeInteracMessageID                               PaymentReferenceReferenceNumberType = "interac_message_id"
	PaymentReferenceReferenceNumberTypeJpmcCcn                                        PaymentReferenceReferenceNumberType = "jpmc_ccn"
	PaymentReferenceReferenceNumberTypeJpmcClearingSystemReference                    PaymentReferenceReferenceNumberType = "jpmc_clearing_system_reference"
	PaymentReferenceReferenceNumberTypeJpmcCustomerReferenceID                        PaymentReferenceReferenceNumberType = "jpmc_customer_reference_id"
	PaymentReferenceReferenceNumberTypeJpmcEndToEndID                                 PaymentReferenceReferenceNumberType = "jpmc_end_to_end_id"
	PaymentReferenceReferenceNumberTypeJpmcFirmRootID                                 PaymentReferenceReferenceNumberType = "jpmc_firm_root_id"
	PaymentReferenceReferenceNumberTypeJpmcFxTrnID                                    PaymentReferenceReferenceNumberType = "jpmc_fx_trn_id"
	PaymentReferenceReferenceNumberTypeJpmcP3ID                                       PaymentReferenceReferenceNumberType = "jpmc_p3_id"
	PaymentReferenceReferenceNumberTypeJpmcPaymentBatchID                             PaymentReferenceReferenceNumberType = "jpmc_payment_batch_id"
	PaymentReferenceReferenceNumberTypeJpmcPaymentInformationID                       PaymentReferenceReferenceNumberType = "jpmc_payment_information_id"
	PaymentReferenceReferenceNumberTypeJpmcPaymentReturnedDatetime                    PaymentReferenceReferenceNumberType = "jpmc_payment_returned_datetime"
	PaymentReferenceReferenceNumberTypeJpmcTransactionReferenceNumber                 PaymentReferenceReferenceNumberType = "jpmc_transaction_reference_number"
	PaymentReferenceReferenceNumberTypeLobCheckID                                     PaymentReferenceReferenceNumberType = "lob_check_id"
	PaymentReferenceReferenceNumberTypeMtFlowTransferID                               PaymentReferenceReferenceNumberType = "mt_flow_transfer_id"
	PaymentReferenceReferenceNumberTypeOther                                          PaymentReferenceReferenceNumberType = "other"
	PaymentReferenceReferenceNumberTypePartialSwiftMir                                PaymentReferenceReferenceNumberType = "partial_swift_mir"
	PaymentReferenceReferenceNumberTypePncClearingReference                           PaymentReferenceReferenceNumberType = "pnc_clearing_reference"
	PaymentReferenceReferenceNumberTypePncInstructionID                               PaymentReferenceReferenceNumberType = "pnc_instruction_id"
	PaymentReferenceReferenceNumberTypePncMultipaymentID                              PaymentReferenceReferenceNumberType = "pnc_multipayment_id"
	PaymentReferenceReferenceNumberTypePncPaymentTraceID                              PaymentReferenceReferenceNumberType = "pnc_payment_trace_id"
	PaymentReferenceReferenceNumberTypePncRequestForPaymentID                         PaymentReferenceReferenceNumberType = "pnc_request_for_payment_id"
	PaymentReferenceReferenceNumberTypePncTransactionReferenceNumber                  PaymentReferenceReferenceNumberType = "pnc_transaction_reference_number"
	PaymentReferenceReferenceNumberTypeRbcWireReferenceID                             PaymentReferenceReferenceNumberType = "rbc_wire_reference_id"
	PaymentReferenceReferenceNumberTypeRtpInstructionID                               PaymentReferenceReferenceNumberType = "rtp_instruction_id"
	PaymentReferenceReferenceNumberTypeSignetAPIReferenceID                           PaymentReferenceReferenceNumberType = "signet_api_reference_id"
	PaymentReferenceReferenceNumberTypeSignetConfirmationID                           PaymentReferenceReferenceNumberType = "signet_confirmation_id"
	PaymentReferenceReferenceNumberTypeSignetRequestID                                PaymentReferenceReferenceNumberType = "signet_request_id"
	PaymentReferenceReferenceNumberTypeSilvergatePaymentID                            PaymentReferenceReferenceNumberType = "silvergate_payment_id"
	PaymentReferenceReferenceNumberTypeSvbEndToEndID                                  PaymentReferenceReferenceNumberType = "svb_end_to_end_id"
	PaymentReferenceReferenceNumberTypeSvbPaymentID                                   PaymentReferenceReferenceNumberType = "svb_payment_id"
	PaymentReferenceReferenceNumberTypeSvbTransactionClearedForSanctionsReview        PaymentReferenceReferenceNumberType = "svb_transaction_cleared_for_sanctions_review"
	PaymentReferenceReferenceNumberTypeSvbTransactionHeldForSanctionsReview           PaymentReferenceReferenceNumberType = "svb_transaction_held_for_sanctions_review"
	PaymentReferenceReferenceNumberTypeSwiftMir                                       PaymentReferenceReferenceNumberType = "swift_mir"
	PaymentReferenceReferenceNumberTypeSwiftUetr                                      PaymentReferenceReferenceNumberType = "swift_uetr"
	PaymentReferenceReferenceNumberTypeUmbProductPartnerAccountNumber                 PaymentReferenceReferenceNumberType = "umb_product_partner_account_number"
	PaymentReferenceReferenceNumberTypeUsbankPaymentApplicationReferenceID            PaymentReferenceReferenceNumberType = "usbank_payment_application_reference_id"
	PaymentReferenceReferenceNumberTypeUsbankPaymentID                                PaymentReferenceReferenceNumberType = "usbank_payment_id"
	PaymentReferenceReferenceNumberTypeUsbankPendingRtpPaymentID                      PaymentReferenceReferenceNumberType = "usbank_pending_rtp_payment_id"
	PaymentReferenceReferenceNumberTypeUsbankPostedRtpPaymentID                       PaymentReferenceReferenceNumberType = "usbank_posted_rtp_payment_id"
	PaymentReferenceReferenceNumberTypeWellsFargoEndToEndID                           PaymentReferenceReferenceNumberType = "wells_fargo_end_to_end_id"
	PaymentReferenceReferenceNumberTypeWellsFargoPaymentID                            PaymentReferenceReferenceNumberType = "wells_fargo_payment_id"
	PaymentReferenceReferenceNumberTypeWellsFargoTraceNumber                          PaymentReferenceReferenceNumberType = "wells_fargo_trace_number"
	PaymentReferenceReferenceNumberTypeWellsFargoUetr                                 PaymentReferenceReferenceNumberType = "wells_fargo_uetr"
	PaymentReferenceReferenceNumberTypeWesternAlliancePaymentID                       PaymentReferenceReferenceNumberType = "western_alliance_payment_id"
	PaymentReferenceReferenceNumberTypeWesternAllianceTransactionID                   PaymentReferenceReferenceNumberType = "western_alliance_transaction_id"
	PaymentReferenceReferenceNumberTypeWesternAllianceWireConfirmationNumber          PaymentReferenceReferenceNumberType = "western_alliance_wire_confirmation_number"
)

func (r PaymentReferenceReferenceNumberType) IsKnown() bool {
	switch r {
	case PaymentReferenceReferenceNumberTypeACHOriginalTraceNumber, PaymentReferenceReferenceNumberTypeACHTraceNumber, PaymentReferenceReferenceNumberTypeBankprovPaymentActivityDate, PaymentReferenceReferenceNumberTypeBankprovPaymentID, PaymentReferenceReferenceNumberTypeBnkDevPrenotificationID, PaymentReferenceReferenceNumberTypeBnkDevTransferID, PaymentReferenceReferenceNumberTypeBnyMellonTransactionReferenceNumber, PaymentReferenceReferenceNumberTypeBofaEndToEndID, PaymentReferenceReferenceNumberTypeBofaTransactionID, PaymentReferenceReferenceNumberTypeBraleTransferID, PaymentReferenceReferenceNumberTypeCheckNumber, PaymentReferenceReferenceNumberTypeCitibankReferenceNumber, PaymentReferenceReferenceNumberTypeCitibankWorldlinkClearingSystemReferenceNumber, PaymentReferenceReferenceNumberTypeColumnFxQuoteID, PaymentReferenceReferenceNumberTypeColumnReversalPairTransferID, PaymentReferenceReferenceNumberTypeColumnTransferID, PaymentReferenceReferenceNumberTypeCrossRiverCoreTransactionID, PaymentReferenceReferenceNumberTypeCrossRiverFedBatchID, PaymentReferenceReferenceNumberTypeCrossRiverPaymentID, PaymentReferenceReferenceNumberTypeCrossRiverServiceMessage, PaymentReferenceReferenceNumberTypeCrossRiverTransactionID, PaymentReferenceReferenceNumberTypeCurrencycloudConversionID, PaymentReferenceReferenceNumberTypeCurrencycloudPaymentID, PaymentReferenceReferenceNumberTypeDcBankTransactionID, PaymentReferenceReferenceNumberTypeEftTraceNumber, PaymentReferenceReferenceNumberTypeEvolveCoreBatch, PaymentReferenceReferenceNumberTypeEvolveCoreFileKey, PaymentReferenceReferenceNumberTypeEvolveCoreSeq, PaymentReferenceReferenceNumberTypeEvolveTransactionID, PaymentReferenceReferenceNumberTypeFakeVendorPaymentID, PaymentReferenceReferenceNumberTypeFedwireImad, PaymentReferenceReferenceNumberTypeFedwireOmad, PaymentReferenceReferenceNumberTypeFirstRepublicInternalID, PaymentReferenceReferenceNumberTypeGoldmanSachsCollectionRequestID, PaymentReferenceReferenceNumberTypeGoldmanSachsEndToEndID, PaymentReferenceReferenceNumberTypeGoldmanSachsPaymentRequestID, PaymentReferenceReferenceNumberTypeGoldmanSachsRequestID, PaymentReferenceReferenceNumberTypeGoldmanSachsUniquePaymentID, PaymentReferenceReferenceNumberTypeHifiOfframpID, PaymentReferenceReferenceNumberTypeHifiTransferID, PaymentReferenceReferenceNumberTypeInteracMessageID, PaymentReferenceReferenceNumberTypeJpmcCcn, PaymentReferenceReferenceNumberTypeJpmcClearingSystemReference, PaymentReferenceReferenceNumberTypeJpmcCustomerReferenceID, PaymentReferenceReferenceNumberTypeJpmcEndToEndID, PaymentReferenceReferenceNumberTypeJpmcFirmRootID, PaymentReferenceReferenceNumberTypeJpmcFxTrnID, PaymentReferenceReferenceNumberTypeJpmcP3ID, PaymentReferenceReferenceNumberTypeJpmcPaymentBatchID, PaymentReferenceReferenceNumberTypeJpmcPaymentInformationID, PaymentReferenceReferenceNumberTypeJpmcPaymentReturnedDatetime, PaymentReferenceReferenceNumberTypeJpmcTransactionReferenceNumber, PaymentReferenceReferenceNumberTypeLobCheckID, PaymentReferenceReferenceNumberTypeMtFlowTransferID, PaymentReferenceReferenceNumberTypeOther, PaymentReferenceReferenceNumberTypePartialSwiftMir, PaymentReferenceReferenceNumberTypePncClearingReference, PaymentReferenceReferenceNumberTypePncInstructionID, PaymentReferenceReferenceNumberTypePncMultipaymentID, PaymentReferenceReferenceNumberTypePncPaymentTraceID, PaymentReferenceReferenceNumberTypePncRequestForPaymentID, PaymentReferenceReferenceNumberTypePncTransactionReferenceNumber, PaymentReferenceReferenceNumberTypeRbcWireReferenceID, PaymentReferenceReferenceNumberTypeRtpInstructionID, PaymentReferenceReferenceNumberTypeSignetAPIReferenceID, PaymentReferenceReferenceNumberTypeSignetConfirmationID, PaymentReferenceReferenceNumberTypeSignetRequestID, PaymentReferenceReferenceNumberTypeSilvergatePaymentID, PaymentReferenceReferenceNumberTypeSvbEndToEndID, PaymentReferenceReferenceNumberTypeSvbPaymentID, PaymentReferenceReferenceNumberTypeSvbTransactionClearedForSanctionsReview, PaymentReferenceReferenceNumberTypeSvbTransactionHeldForSanctionsReview, PaymentReferenceReferenceNumberTypeSwiftMir, PaymentReferenceReferenceNumberTypeSwiftUetr, PaymentReferenceReferenceNumberTypeUmbProductPartnerAccountNumber, PaymentReferenceReferenceNumberTypeUsbankPaymentApplicationReferenceID, PaymentReferenceReferenceNumberTypeUsbankPaymentID, PaymentReferenceReferenceNumberTypeUsbankPendingRtpPaymentID, PaymentReferenceReferenceNumberTypeUsbankPostedRtpPaymentID, PaymentReferenceReferenceNumberTypeWellsFargoEndToEndID, PaymentReferenceReferenceNumberTypeWellsFargoPaymentID, PaymentReferenceReferenceNumberTypeWellsFargoTraceNumber, PaymentReferenceReferenceNumberTypeWellsFargoUetr, PaymentReferenceReferenceNumberTypeWesternAlliancePaymentID, PaymentReferenceReferenceNumberTypeWesternAllianceTransactionID, PaymentReferenceReferenceNumberTypeWesternAllianceWireConfirmationNumber:
		return true
	}
	return false
}

// One of the referenceable types. This must be accompanied by the id of the
// referenceable or will return an error.
type PaymentReferenceReferenceableType string

const (
	PaymentReferenceReferenceableTypePaymentOrder PaymentReferenceReferenceableType = "payment_order"
	PaymentReferenceReferenceableTypeReversal     PaymentReferenceReferenceableType = "reversal"
	PaymentReferenceReferenceableTypeReturn       PaymentReferenceReferenceableType = "return"
)

func (r PaymentReferenceReferenceableType) IsKnown() bool {
	switch r {
	case PaymentReferenceReferenceableTypePaymentOrder, PaymentReferenceReferenceableTypeReversal, PaymentReferenceReferenceableTypeReturn:
		return true
	}
	return false
}

type PaymentReferenceListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	PerPage     param.Field[int64]  `query:"per_page"`
	// The actual reference number assigned by the bank.
	ReferenceNumber param.Field[string] `query:"reference_number"`
	// The id of the referenceable to search for. Must be accompanied by the
	// referenceable_type or will return an error.
	ReferenceableID param.Field[string] `query:"referenceable_id"`
	// One of the referenceable types. This must be accompanied by the id of the
	// referenceable or will return an error.
	ReferenceableType param.Field[PaymentReferenceListParamsReferenceableType] `query:"referenceable_type"`
}

// URLQuery serializes [PaymentReferenceListParams]'s query parameters as
// `url.Values`.
func (r PaymentReferenceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// One of the referenceable types. This must be accompanied by the id of the
// referenceable or will return an error.
type PaymentReferenceListParamsReferenceableType string

const (
	PaymentReferenceListParamsReferenceableTypePaymentOrder PaymentReferenceListParamsReferenceableType = "payment_order"
	PaymentReferenceListParamsReferenceableTypeReturn       PaymentReferenceListParamsReferenceableType = "return"
	PaymentReferenceListParamsReferenceableTypeReversal     PaymentReferenceListParamsReferenceableType = "reversal"
)

func (r PaymentReferenceListParamsReferenceableType) IsKnown() bool {
	switch r {
	case PaymentReferenceListParamsReferenceableTypePaymentOrder, PaymentReferenceListParamsReferenceableTypeReturn, PaymentReferenceListParamsReferenceableTypeReversal:
		return true
	}
	return false
}
