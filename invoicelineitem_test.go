// File generated from our OpenAPI spec by Stainless.

package moderntreasury_test

import (
	"context"
	"errors"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestInvoiceLineItemNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.Invoices.LineItems.New(
		context.TODO(),
		"string",
		moderntreasury.InvoiceLineItemNewParams{
			Name:           moderntreasury.F("string"),
			UnitAmount:     moderntreasury.F(int64(0)),
			Description:    moderntreasury.F("string"),
			Direction:      moderntreasury.F("string"),
			Quantity:       moderntreasury.F(int64(0)),
			IdempotencyKey: moderntreasury.F("string"),
		},
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceLineItemGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.Invoices.LineItems.Get(
		context.TODO(),
		"string",
		"string",
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceLineItemUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.Invoices.LineItems.Update(
		context.TODO(),
		"string",
		"string",
		moderntreasury.InvoiceLineItemUpdateParams{
			ContactDetails: moderntreasury.F([]moderntreasury.InvoiceLineItemUpdateParamsContactDetail{{
				ID:                    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Object:                moderntreasury.F("string"),
				LiveMode:              moderntreasury.F(true),
				CreatedAt:             moderntreasury.F(time.Now()),
				UpdatedAt:             moderntreasury.F(time.Now()),
				DiscardedAt:           moderntreasury.F(time.Now()),
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.InvoiceLineItemUpdateParamsContactDetailsContactIdentifierTypeEmail),
			}, {
				ID:                    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Object:                moderntreasury.F("string"),
				LiveMode:              moderntreasury.F(true),
				CreatedAt:             moderntreasury.F(time.Now()),
				UpdatedAt:             moderntreasury.F(time.Now()),
				DiscardedAt:           moderntreasury.F(time.Now()),
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.InvoiceLineItemUpdateParamsContactDetailsContactIdentifierTypeEmail),
			}, {
				ID:                    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Object:                moderntreasury.F("string"),
				LiveMode:              moderntreasury.F(true),
				CreatedAt:             moderntreasury.F(time.Now()),
				UpdatedAt:             moderntreasury.F(time.Now()),
				DiscardedAt:           moderntreasury.F(time.Now()),
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.InvoiceLineItemUpdateParamsContactDetailsContactIdentifierTypeEmail),
			}}),
			CounterpartyBillingAddress: moderntreasury.F(moderntreasury.InvoiceLineItemUpdateParamsCounterpartyBillingAddress{
				Line1:      moderntreasury.F("string"),
				Line2:      moderntreasury.F("string"),
				Locality:   moderntreasury.F("string"),
				Region:     moderntreasury.F("string"),
				PostalCode: moderntreasury.F("string"),
				Country:    moderntreasury.F("string"),
			}),
			CounterpartyID: moderntreasury.F("string"),
			CounterpartyShippingAddress: moderntreasury.F(moderntreasury.InvoiceLineItemUpdateParamsCounterpartyShippingAddress{
				Line1:      moderntreasury.F("string"),
				Line2:      moderntreasury.F("string"),
				Locality:   moderntreasury.F("string"),
				Region:     moderntreasury.F("string"),
				PostalCode: moderntreasury.F("string"),
				Country:    moderntreasury.F("string"),
			}),
			Currency:    moderntreasury.F(shared.CurrencyAed),
			Description: moderntreasury.F("string"),
			DueDate:     moderntreasury.F(time.Now()),
			InvoicerAddress: moderntreasury.F(moderntreasury.InvoiceLineItemUpdateParamsInvoicerAddress{
				Line1:      moderntreasury.F("string"),
				Line2:      moderntreasury.F("string"),
				Locality:   moderntreasury.F("string"),
				Region:     moderntreasury.F("string"),
				PostalCode: moderntreasury.F("string"),
				Country:    moderntreasury.F("string"),
			}),
			OriginatingAccountID: moderntreasury.F("string"),
			PaymentEffectiveDate: moderntreasury.F(time.Now()),
			PaymentMethod:        moderntreasury.F(moderntreasury.InvoiceLineItemUpdateParamsPaymentMethodUi),
			PaymentType:          moderntreasury.F(moderntreasury.InvoiceLineItemUpdateParamsPaymentTypeACH),
			ReceivingAccountID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceLineItemListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.Invoices.LineItems.List(
		context.TODO(),
		"string",
		moderntreasury.InvoiceLineItemListParams{
			AfterCursor: moderntreasury.F("string"),
			PerPage:     moderntreasury.F(int64(0)),
		},
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceLineItemDelete(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.Invoices.LineItems.Delete(
		context.TODO(),
		"string",
		"string",
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
