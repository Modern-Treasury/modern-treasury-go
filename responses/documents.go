package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type Document struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// A category given to the document, can be `null`.
	DocumentType string `json:"document_type,required,nullable"`
	// The unique identifier for the associated object.
	DocumentableID string `json:"documentable_id,required" format:"uuid"`
	// The type of the associated object. Currently can be one of `payment_order`,
	// `transaction`, `paper_item`, `expected_payment`, `counterparty`, `organization`,
	// `case`, `internal_account` or `external_account`.
	DocumentableType DocumentDocumentableType  `json:"documentable_type,required"`
	DocumentDetails  []DocumentDocumentDetails `json:"document_details,required"`
	File             DocumentFile              `json:"file,required"`
	JSON             DocumentJSON
}

type DocumentJSON struct {
	ID               pjson.Metadata
	Object           pjson.Metadata
	LiveMode         pjson.Metadata
	CreatedAt        pjson.Metadata
	UpdatedAt        pjson.Metadata
	DiscardedAt      pjson.Metadata
	DocumentType     pjson.Metadata
	DocumentableID   pjson.Metadata
	DocumentableType pjson.Metadata
	DocumentDetails  pjson.Metadata
	File             pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Document using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Document) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type DocumentDocumentableType string

const (
	DocumentDocumentableTypeCase            DocumentDocumentableType = "case"
	DocumentDocumentableTypeCounterparty    DocumentDocumentableType = "counterparty"
	DocumentDocumentableTypeExpectedPayment DocumentDocumentableType = "expected_payment"
	DocumentDocumentableTypeExternalAccount DocumentDocumentableType = "external_account"
	DocumentDocumentableTypeInternalAccount DocumentDocumentableType = "internal_account"
	DocumentDocumentableTypeOrganization    DocumentDocumentableType = "organization"
	DocumentDocumentableTypePaperItem       DocumentDocumentableType = "paper_item"
	DocumentDocumentableTypePaymentOrder    DocumentDocumentableType = "payment_order"
	DocumentDocumentableTypeTransaction     DocumentDocumentableType = "transaction"
)

type DocumentDocumentDetails struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode               bool      `json:"live_mode,required"`
	CreatedAt              time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt              time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt            time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	DocumentIdentifierType string    `json:"document_identifier_type,required"`
	DocumentIdentifier     string    `json:"document_identifier,required"`
	JSON                   DocumentDocumentDetailsJSON
}

type DocumentDocumentDetailsJSON struct {
	ID                     pjson.Metadata
	Object                 pjson.Metadata
	LiveMode               pjson.Metadata
	CreatedAt              pjson.Metadata
	UpdatedAt              pjson.Metadata
	DiscardedAt            pjson.Metadata
	DocumentIdentifierType pjson.Metadata
	DocumentIdentifier     pjson.Metadata
	Raw                    []byte
	Extras                 map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into DocumentDocumentDetails using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DocumentDocumentDetails) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type DocumentFile struct {
	// The size of the document in bytes.
	Size int64 `json:"size"`
	// The original filename of the document.
	Filename string `json:"filename"`
	// The MIME content type of the document.
	ContentType string `json:"content_type"`
	JSON        DocumentFileJSON
}

type DocumentFileJSON struct {
	Size        pjson.Metadata
	Filename    pjson.Metadata
	ContentType pjson.Metadata
	Raw         []byte
	Extras      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into DocumentFile using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DocumentFile) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
