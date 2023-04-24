package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/core"
	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
)

func TestInvoiceNewWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Invoices.New(context.TODO(), &requests.InvoiceNewParams{ContactDetails: moderntreasury.F([]requests.InvoiceNewParamsContactDetails{{ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.InvoiceNewParamsContactDetailsContactIdentifierTypeEmail)}, {ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.InvoiceNewParamsContactDetailsContactIdentifierTypeEmail)}, {ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.InvoiceNewParamsContactDetailsContactIdentifierTypeEmail)}}), CounterpartyID: moderntreasury.F("string"), CounterpartyBillingAddress: moderntreasury.F(requests.InvoiceNewParamsCounterpartyBillingAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), CounterpartyShippingAddress: moderntreasury.F(requests.InvoiceNewParamsCounterpartyShippingAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), Currency: moderntreasury.F(requests.CurrencyAed), Description: moderntreasury.F("string"), DueDate: moderntreasury.F(time.Now()), InvoicerAddress: moderntreasury.F(requests.InvoiceNewParamsInvoicerAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), OriginatingAccountID: moderntreasury.F("string")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceGet(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Invoices.Get(
		context.TODO(),
		"string",
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Invoices.Update(
		context.TODO(),
		"string",
		&requests.InvoiceUpdateParams{ContactDetails: moderntreasury.F([]requests.InvoiceUpdateParamsContactDetails{{ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.InvoiceUpdateParamsContactDetailsContactIdentifierTypeEmail)}, {ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.InvoiceUpdateParamsContactDetailsContactIdentifierTypeEmail)}, {ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.InvoiceUpdateParamsContactDetailsContactIdentifierTypeEmail)}}), CounterpartyID: moderntreasury.F("string"), CounterpartyBillingAddress: moderntreasury.F(requests.InvoiceUpdateParamsCounterpartyBillingAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), CounterpartyShippingAddress: moderntreasury.F(requests.InvoiceUpdateParamsCounterpartyShippingAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), Currency: moderntreasury.F(requests.CurrencyAed), Description: moderntreasury.F("string"), DueDate: moderntreasury.F(time.Now()), InvoicerAddress: moderntreasury.F(requests.InvoiceUpdateParamsInvoicerAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), OriginatingAccountID: moderntreasury.F("string"), IncludePaymentUi: moderntreasury.F(true), Status: moderntreasury.F("string")},
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Invoices.List(context.TODO(), &requests.InvoiceListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0))})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
