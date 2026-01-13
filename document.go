// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiform"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/packages/pagination"
)

// DocumentService contains methods and other services that help with interacting
// with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocumentService] method instead.
type DocumentService struct {
	Options []option.RequestOption
}

// NewDocumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDocumentService(opts ...option.RequestOption) (r *DocumentService) {
	r = &DocumentService{}
	r.Options = opts
	return
}

// Create a document.
func (r *DocumentService) New(ctx context.Context, body DocumentNewParams, opts ...option.RequestOption) (res *Document, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/documents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an existing document.
func (r *DocumentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Document, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/documents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of documents.
func (r *DocumentService) List(ctx context.Context, query DocumentListParams, opts ...option.RequestOption) (res *pagination.Page[Document], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/documents"
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

// Get a list of documents.
func (r *DocumentService) ListAutoPaging(ctx context.Context, query DocumentListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Document] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type Document struct {
	ID              string                   `json:"id,required" format:"uuid"`
	CreatedAt       time.Time                `json:"created_at,required" format:"date-time"`
	DiscardedAt     time.Time                `json:"discarded_at,required,nullable" format:"date-time"`
	DocumentDetails []DocumentDocumentDetail `json:"document_details,required"`
	// A category given to the document, can be `null`.
	DocumentType string `json:"document_type,required,nullable"`
	// The unique identifier for the associated object.
	DocumentableID string `json:"documentable_id,required,nullable" format:"uuid"`
	// The type of the associated object. Currently can be one of `payment_order`,
	// `transaction`, `expected_payment`, `counterparty`, `organization`, `case`,
	// `internal_account`, `decision`, or `external_account`.
	DocumentableType DocumentDocumentableType `json:"documentable_type,required,nullable"`
	File             DocumentFile             `json:"file,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode,required"`
	Object   string `json:"object,required"`
	// The source of the document. Can be `vendor`, `customer`, or `modern_treasury`.
	Source    string       `json:"source,required"`
	UpdatedAt time.Time    `json:"updated_at,required" format:"date-time"`
	JSON      documentJSON `json:"-"`
}

// documentJSON contains the JSON metadata for the struct [Document]
type documentJSON struct {
	ID               apijson.Field
	CreatedAt        apijson.Field
	DiscardedAt      apijson.Field
	DocumentDetails  apijson.Field
	DocumentType     apijson.Field
	DocumentableID   apijson.Field
	DocumentableType apijson.Field
	File             apijson.Field
	LiveMode         apijson.Field
	Object           apijson.Field
	Source           apijson.Field
	UpdatedAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Document) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r documentJSON) RawJSON() string {
	return r.raw
}

type DocumentDocumentDetail struct {
	ID                     string    `json:"id,required" format:"uuid"`
	CreatedAt              time.Time `json:"created_at,required" format:"date-time"`
	DiscardedAt            time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	DocumentIdentifier     string    `json:"document_identifier,required"`
	DocumentIdentifierType string    `json:"document_identifier_type,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool                       `json:"live_mode,required"`
	Object    string                     `json:"object,required"`
	UpdatedAt time.Time                  `json:"updated_at,required" format:"date-time"`
	JSON      documentDocumentDetailJSON `json:"-"`
}

// documentDocumentDetailJSON contains the JSON metadata for the struct
// [DocumentDocumentDetail]
type documentDocumentDetailJSON struct {
	ID                     apijson.Field
	CreatedAt              apijson.Field
	DiscardedAt            apijson.Field
	DocumentIdentifier     apijson.Field
	DocumentIdentifierType apijson.Field
	LiveMode               apijson.Field
	Object                 apijson.Field
	UpdatedAt              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *DocumentDocumentDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r documentDocumentDetailJSON) RawJSON() string {
	return r.raw
}

// The type of the associated object. Currently can be one of `payment_order`,
// `transaction`, `expected_payment`, `counterparty`, `organization`, `case`,
// `internal_account`, `decision`, or `external_account`.
type DocumentDocumentableType string

const (
	DocumentDocumentableTypeConnection            DocumentDocumentableType = "connection"
	DocumentDocumentableTypeCounterparty          DocumentDocumentableType = "counterparty"
	DocumentDocumentableTypeExpectedPayment       DocumentDocumentableType = "expected_payment"
	DocumentDocumentableTypeExternalAccount       DocumentDocumentableType = "external_account"
	DocumentDocumentableTypeIdentification        DocumentDocumentableType = "identification"
	DocumentDocumentableTypeIncomingPaymentDetail DocumentDocumentableType = "incoming_payment_detail"
	DocumentDocumentableTypeInternalAccount       DocumentDocumentableType = "internal_account"
	DocumentDocumentableTypeLegalEntity           DocumentDocumentableType = "legal_entity"
	DocumentDocumentableTypeOrganization          DocumentDocumentableType = "organization"
	DocumentDocumentableTypePaymentOrder          DocumentDocumentableType = "payment_order"
	DocumentDocumentableTypeTransaction           DocumentDocumentableType = "transaction"
)

func (r DocumentDocumentableType) IsKnown() bool {
	switch r {
	case DocumentDocumentableTypeConnection, DocumentDocumentableTypeCounterparty, DocumentDocumentableTypeExpectedPayment, DocumentDocumentableTypeExternalAccount, DocumentDocumentableTypeIdentification, DocumentDocumentableTypeIncomingPaymentDetail, DocumentDocumentableTypeInternalAccount, DocumentDocumentableTypeLegalEntity, DocumentDocumentableTypeOrganization, DocumentDocumentableTypePaymentOrder, DocumentDocumentableTypeTransaction:
		return true
	}
	return false
}

type DocumentFile struct {
	// The MIME content type of the document.
	ContentType string `json:"content_type"`
	// The original filename of the document.
	Filename string `json:"filename"`
	// The size of the document in bytes.
	Size int64            `json:"size"`
	JSON documentFileJSON `json:"-"`
}

// documentFileJSON contains the JSON metadata for the struct [DocumentFile]
type documentFileJSON struct {
	ContentType apijson.Field
	Filename    apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DocumentFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r documentFileJSON) RawJSON() string {
	return r.raw
}

type DocumentNewParams struct {
	File param.Field[io.Reader] `json:"file,required" format:"binary"`
	// A category given to the document, can be `null`.
	DocumentType param.Field[string] `json:"document_type"`
	// The unique identifier for the associated object.
	DocumentableID   param.Field[string]                            `json:"documentable_id"`
	DocumentableType param.Field[DocumentNewParamsDocumentableType] `json:"documentable_type"`
}

func (r DocumentNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type DocumentNewParamsDocumentableType string

const (
	DocumentNewParamsDocumentableTypeConnections            DocumentNewParamsDocumentableType = "connections"
	DocumentNewParamsDocumentableTypeCounterparties         DocumentNewParamsDocumentableType = "counterparties"
	DocumentNewParamsDocumentableTypeExpectedPayments       DocumentNewParamsDocumentableType = "expected_payments"
	DocumentNewParamsDocumentableTypeExternalAccounts       DocumentNewParamsDocumentableType = "external_accounts"
	DocumentNewParamsDocumentableTypeIdentifications        DocumentNewParamsDocumentableType = "identifications"
	DocumentNewParamsDocumentableTypeIncomingPaymentDetails DocumentNewParamsDocumentableType = "incoming_payment_details"
	DocumentNewParamsDocumentableTypeInternalAccounts       DocumentNewParamsDocumentableType = "internal_accounts"
	DocumentNewParamsDocumentableTypeLegalEntities          DocumentNewParamsDocumentableType = "legal_entities"
	DocumentNewParamsDocumentableTypeOrganizations          DocumentNewParamsDocumentableType = "organizations"
	DocumentNewParamsDocumentableTypePaymentOrders          DocumentNewParamsDocumentableType = "payment_orders"
	DocumentNewParamsDocumentableTypeTransactions           DocumentNewParamsDocumentableType = "transactions"
)

func (r DocumentNewParamsDocumentableType) IsKnown() bool {
	switch r {
	case DocumentNewParamsDocumentableTypeConnections, DocumentNewParamsDocumentableTypeCounterparties, DocumentNewParamsDocumentableTypeExpectedPayments, DocumentNewParamsDocumentableTypeExternalAccounts, DocumentNewParamsDocumentableTypeIdentifications, DocumentNewParamsDocumentableTypeIncomingPaymentDetails, DocumentNewParamsDocumentableTypeInternalAccounts, DocumentNewParamsDocumentableTypeLegalEntities, DocumentNewParamsDocumentableTypeOrganizations, DocumentNewParamsDocumentableTypePaymentOrders, DocumentNewParamsDocumentableTypeTransactions:
		return true
	}
	return false
}

type DocumentListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// The unique identifier for the associated object.
	DocumentableID param.Field[string] `query:"documentable_id"`
	// The type of the associated object. Currently can be one of `payment_order`,
	// `transaction`, `expected_payment`, `counterparty`, `organization`, `case`,
	// `internal_account`, `decision`, or `external_account`.
	DocumentableType param.Field[DocumentListParamsDocumentableType] `query:"documentable_type"`
	PerPage          param.Field[int64]                              `query:"per_page"`
}

// URLQuery serializes [DocumentListParams]'s query parameters as `url.Values`.
func (r DocumentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of the associated object. Currently can be one of `payment_order`,
// `transaction`, `expected_payment`, `counterparty`, `organization`, `case`,
// `internal_account`, `decision`, or `external_account`.
type DocumentListParamsDocumentableType string

const (
	DocumentListParamsDocumentableTypeConnections            DocumentListParamsDocumentableType = "connections"
	DocumentListParamsDocumentableTypeCounterparties         DocumentListParamsDocumentableType = "counterparties"
	DocumentListParamsDocumentableTypeExpectedPayments       DocumentListParamsDocumentableType = "expected_payments"
	DocumentListParamsDocumentableTypeExternalAccounts       DocumentListParamsDocumentableType = "external_accounts"
	DocumentListParamsDocumentableTypeIdentifications        DocumentListParamsDocumentableType = "identifications"
	DocumentListParamsDocumentableTypeIncomingPaymentDetails DocumentListParamsDocumentableType = "incoming_payment_details"
	DocumentListParamsDocumentableTypeInternalAccounts       DocumentListParamsDocumentableType = "internal_accounts"
	DocumentListParamsDocumentableTypeLegalEntities          DocumentListParamsDocumentableType = "legal_entities"
	DocumentListParamsDocumentableTypeOrganizations          DocumentListParamsDocumentableType = "organizations"
	DocumentListParamsDocumentableTypePaymentOrders          DocumentListParamsDocumentableType = "payment_orders"
	DocumentListParamsDocumentableTypeTransactions           DocumentListParamsDocumentableType = "transactions"
)

func (r DocumentListParamsDocumentableType) IsKnown() bool {
	switch r {
	case DocumentListParamsDocumentableTypeConnections, DocumentListParamsDocumentableTypeCounterparties, DocumentListParamsDocumentableTypeExpectedPayments, DocumentListParamsDocumentableTypeExternalAccounts, DocumentListParamsDocumentableTypeIdentifications, DocumentListParamsDocumentableTypeIncomingPaymentDetails, DocumentListParamsDocumentableTypeInternalAccounts, DocumentListParamsDocumentableTypeLegalEntities, DocumentListParamsDocumentableTypeOrganizations, DocumentListParamsDocumentableTypePaymentOrders, DocumentListParamsDocumentableTypeTransactions:
		return true
	}
	return false
}
