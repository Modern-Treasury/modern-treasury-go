package requests

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/url"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type DocumentNewParams struct {
	// A category given to the document, can be `null`.
	DocumentType field.Field[string]    `form:"document_type"`
	File         field.Field[io.Reader] `form:"file,required" format:"binary"`
}

func (r DocumentNewParams) MarshalMultipart() (data []byte, err error) {
	body := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(body)
	defer writer.Close()
	{
		bdy, err := pjson.Marshal(r.DocumentType)
		if err != nil {
			return nil, err
		}
		writer.WriteField("document_type", string(bdy))
	}
	{
		name := "anonymous_file"
		if nameable, ok := r.File.Value.(interface{ Name() string }); ok {
			name = nameable.Name()
		}
		part, err := writer.CreateFormFile("file", name)
		if err != nil {
			return nil, err
		}
		io.Copy(part, r.File.Value)
	}
	return body.Bytes(), nil
}

type DocumentNewParamsDocumentableType string

const (
	DocumentNewParamsDocumentableTypeCases            DocumentNewParamsDocumentableType = "cases"
	DocumentNewParamsDocumentableTypeCounterparties   DocumentNewParamsDocumentableType = "counterparties"
	DocumentNewParamsDocumentableTypeExpectedPayments DocumentNewParamsDocumentableType = "expected_payments"
	DocumentNewParamsDocumentableTypeExternalAccounts DocumentNewParamsDocumentableType = "external_accounts"
	DocumentNewParamsDocumentableTypeInternalAccounts DocumentNewParamsDocumentableType = "internal_accounts"
	DocumentNewParamsDocumentableTypeOrganizations    DocumentNewParamsDocumentableType = "organizations"
	DocumentNewParamsDocumentableTypePaperItems       DocumentNewParamsDocumentableType = "paper_items"
	DocumentNewParamsDocumentableTypePaymentOrders    DocumentNewParamsDocumentableType = "payment_orders"
	DocumentNewParamsDocumentableTypeTransactions     DocumentNewParamsDocumentableType = "transactions"
)

type DocumentGetParamsDocumentableType string

const (
	DocumentGetParamsDocumentableTypeCases            DocumentGetParamsDocumentableType = "cases"
	DocumentGetParamsDocumentableTypeCounterparties   DocumentGetParamsDocumentableType = "counterparties"
	DocumentGetParamsDocumentableTypeExpectedPayments DocumentGetParamsDocumentableType = "expected_payments"
	DocumentGetParamsDocumentableTypeExternalAccounts DocumentGetParamsDocumentableType = "external_accounts"
	DocumentGetParamsDocumentableTypeInternalAccounts DocumentGetParamsDocumentableType = "internal_accounts"
	DocumentGetParamsDocumentableTypeOrganizations    DocumentGetParamsDocumentableType = "organizations"
	DocumentGetParamsDocumentableTypePaperItems       DocumentGetParamsDocumentableType = "paper_items"
	DocumentGetParamsDocumentableTypePaymentOrders    DocumentGetParamsDocumentableType = "payment_orders"
	DocumentGetParamsDocumentableTypeTransactions     DocumentGetParamsDocumentableType = "transactions"
)

type DocumentListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
}

// URLQuery serializes DocumentListParams into a url.Values of the query parameters
// associated with this value
func (r DocumentListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type DocumentListParamsDocumentableType string

const (
	DocumentListParamsDocumentableTypeCases            DocumentListParamsDocumentableType = "cases"
	DocumentListParamsDocumentableTypeCounterparties   DocumentListParamsDocumentableType = "counterparties"
	DocumentListParamsDocumentableTypeExpectedPayments DocumentListParamsDocumentableType = "expected_payments"
	DocumentListParamsDocumentableTypeExternalAccounts DocumentListParamsDocumentableType = "external_accounts"
	DocumentListParamsDocumentableTypeInternalAccounts DocumentListParamsDocumentableType = "internal_accounts"
	DocumentListParamsDocumentableTypeOrganizations    DocumentListParamsDocumentableType = "organizations"
	DocumentListParamsDocumentableTypePaperItems       DocumentListParamsDocumentableType = "paper_items"
	DocumentListParamsDocumentableTypePaymentOrders    DocumentListParamsDocumentableType = "payment_orders"
	DocumentListParamsDocumentableTypeTransactions     DocumentListParamsDocumentableType = "transactions"
)
