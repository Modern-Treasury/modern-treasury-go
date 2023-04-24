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

func TestInvoiceLineItemNewWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Invoices.LineItems.New(
		context.TODO(),
		"string",
		&requests.InvoiceLineItemNewParams{Name: moderntreasury.F("string"), Description: moderntreasury.F("string"), Quantity: moderntreasury.F(int64(0)), UnitAmount: moderntreasury.F(int64(0)), Direction: moderntreasury.F("string")},
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

func TestInvoiceLineItemGet(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Invoices.LineItems.Get(
		context.TODO(),
		"string",
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

func TestInvoiceLineItemUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Invoices.LineItems.Update(
		context.TODO(),
		"string",
		"string",
		&requests.InvoiceLineItemUpdateParams{ContactDetails: moderntreasury.F([]requests.InvoiceLineItemUpdateParamsContactDetails{{ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.InvoiceLineItemUpdateParamsContactDetailsContactIdentifierTypeEmail)}, {ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.InvoiceLineItemUpdateParamsContactDetailsContactIdentifierTypeEmail)}, {ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.InvoiceLineItemUpdateParamsContactDetailsContactIdentifierTypeEmail)}}), CounterpartyID: moderntreasury.F("string"), CounterpartyBillingAddress: moderntreasury.F(requests.InvoiceLineItemUpdateParamsCounterpartyBillingAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), CounterpartyShippingAddress: moderntreasury.F(requests.InvoiceLineItemUpdateParamsCounterpartyShippingAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), Currency: moderntreasury.F(requests.CurrencyAed), Description: moderntreasury.F("string"), DueDate: moderntreasury.F(time.Now()), InvoicerAddress: moderntreasury.F(requests.InvoiceLineItemUpdateParamsInvoicerAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), OriginatingAccountID: moderntreasury.F("string")},
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

func TestInvoiceLineItemListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Invoices.LineItems.List(
		context.TODO(),
		"string",
		&requests.InvoiceLineItemListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0))},
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

func TestInvoiceLineItemDelete(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Invoices.LineItems.Delete(
		context.TODO(),
		"string",
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
