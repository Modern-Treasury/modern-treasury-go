// File generated from our OpenAPI spec by Stainless.

package moderntreasury_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go/v2"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

func TestInvoiceNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.Invoices.New(context.TODO(), moderntreasury.InvoiceNewParams{
		CounterpartyID:       moderntreasury.F("string"),
		DueDate:              moderntreasury.F(time.Now()),
		OriginatingAccountID: moderntreasury.F("string"),
		ContactDetails: moderntreasury.F([]moderntreasury.InvoiceNewParamsContactDetail{{
			ID:                    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Object:                moderntreasury.F("string"),
			LiveMode:              moderntreasury.F(true),
			CreatedAt:             moderntreasury.F(time.Now()),
			UpdatedAt:             moderntreasury.F(time.Now()),
			DiscardedAt:           moderntreasury.F(time.Now()),
			ContactIdentifier:     moderntreasury.F("string"),
			ContactIdentifierType: moderntreasury.F(moderntreasury.InvoiceNewParamsContactDetailsContactIdentifierTypeEmail),
		}, {
			ID:                    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Object:                moderntreasury.F("string"),
			LiveMode:              moderntreasury.F(true),
			CreatedAt:             moderntreasury.F(time.Now()),
			UpdatedAt:             moderntreasury.F(time.Now()),
			DiscardedAt:           moderntreasury.F(time.Now()),
			ContactIdentifier:     moderntreasury.F("string"),
			ContactIdentifierType: moderntreasury.F(moderntreasury.InvoiceNewParamsContactDetailsContactIdentifierTypeEmail),
		}, {
			ID:                    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Object:                moderntreasury.F("string"),
			LiveMode:              moderntreasury.F(true),
			CreatedAt:             moderntreasury.F(time.Now()),
			UpdatedAt:             moderntreasury.F(time.Now()),
			DiscardedAt:           moderntreasury.F(time.Now()),
			ContactIdentifier:     moderntreasury.F("string"),
			ContactIdentifierType: moderntreasury.F(moderntreasury.InvoiceNewParamsContactDetailsContactIdentifierTypeEmail),
		}}),
		CounterpartyBillingAddress: moderntreasury.F(moderntreasury.InvoiceNewParamsCounterpartyBillingAddress{
			Line1:      moderntreasury.F("string"),
			Line2:      moderntreasury.F("string"),
			Locality:   moderntreasury.F("string"),
			Region:     moderntreasury.F("string"),
			PostalCode: moderntreasury.F("string"),
			Country:    moderntreasury.F("string"),
		}),
		CounterpartyShippingAddress: moderntreasury.F(moderntreasury.InvoiceNewParamsCounterpartyShippingAddress{
			Line1:      moderntreasury.F("string"),
			Line2:      moderntreasury.F("string"),
			Locality:   moderntreasury.F("string"),
			Region:     moderntreasury.F("string"),
			PostalCode: moderntreasury.F("string"),
			Country:    moderntreasury.F("string"),
		}),
		Currency:    moderntreasury.F(shared.CurrencyAed),
		Description: moderntreasury.F("string"),
		InvoicerAddress: moderntreasury.F(moderntreasury.InvoiceNewParamsInvoicerAddress{
			Line1:      moderntreasury.F("string"),
			Line2:      moderntreasury.F("string"),
			Locality:   moderntreasury.F("string"),
			Region:     moderntreasury.F("string"),
			PostalCode: moderntreasury.F("string"),
			Country:    moderntreasury.F("string"),
		}),
		NotificationEmailAddresses: moderntreasury.F([]string{"string", "string", "string"}),
		NotificationsEnabled:       moderntreasury.F(true),
		PaymentEffectiveDate:       moderntreasury.F(time.Now()),
		PaymentMethod:              moderntreasury.F(moderntreasury.InvoiceNewParamsPaymentMethodUi),
		PaymentType:                moderntreasury.F(moderntreasury.PaymentOrderTypeACH),
		ReceivingAccountID:         moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		RecipientEmail:             moderntreasury.F("string"),
		RecipientName:              moderntreasury.F("string"),
		VirtualAccountID:           moderntreasury.F("string"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.Invoices.Get(context.TODO(), "string")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.Invoices.Update(
		context.TODO(),
		"string",
		moderntreasury.InvoiceUpdateParams{
			ContactDetails: moderntreasury.F([]moderntreasury.InvoiceUpdateParamsContactDetail{{
				ID:                    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Object:                moderntreasury.F("string"),
				LiveMode:              moderntreasury.F(true),
				CreatedAt:             moderntreasury.F(time.Now()),
				UpdatedAt:             moderntreasury.F(time.Now()),
				DiscardedAt:           moderntreasury.F(time.Now()),
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.InvoiceUpdateParamsContactDetailsContactIdentifierTypeEmail),
			}, {
				ID:                    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Object:                moderntreasury.F("string"),
				LiveMode:              moderntreasury.F(true),
				CreatedAt:             moderntreasury.F(time.Now()),
				UpdatedAt:             moderntreasury.F(time.Now()),
				DiscardedAt:           moderntreasury.F(time.Now()),
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.InvoiceUpdateParamsContactDetailsContactIdentifierTypeEmail),
			}, {
				ID:                    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Object:                moderntreasury.F("string"),
				LiveMode:              moderntreasury.F(true),
				CreatedAt:             moderntreasury.F(time.Now()),
				UpdatedAt:             moderntreasury.F(time.Now()),
				DiscardedAt:           moderntreasury.F(time.Now()),
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.InvoiceUpdateParamsContactDetailsContactIdentifierTypeEmail),
			}}),
			CounterpartyBillingAddress: moderntreasury.F(moderntreasury.InvoiceUpdateParamsCounterpartyBillingAddress{
				Line1:      moderntreasury.F("string"),
				Line2:      moderntreasury.F("string"),
				Locality:   moderntreasury.F("string"),
				Region:     moderntreasury.F("string"),
				PostalCode: moderntreasury.F("string"),
				Country:    moderntreasury.F("string"),
			}),
			CounterpartyID: moderntreasury.F("string"),
			CounterpartyShippingAddress: moderntreasury.F(moderntreasury.InvoiceUpdateParamsCounterpartyShippingAddress{
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
			InvoicerAddress: moderntreasury.F(moderntreasury.InvoiceUpdateParamsInvoicerAddress{
				Line1:      moderntreasury.F("string"),
				Line2:      moderntreasury.F("string"),
				Locality:   moderntreasury.F("string"),
				Region:     moderntreasury.F("string"),
				PostalCode: moderntreasury.F("string"),
				Country:    moderntreasury.F("string"),
			}),
			NotificationEmailAddresses: moderntreasury.F([]string{"string", "string", "string"}),
			NotificationsEnabled:       moderntreasury.F(true),
			OriginatingAccountID:       moderntreasury.F("string"),
			PaymentEffectiveDate:       moderntreasury.F(time.Now()),
			PaymentMethod:              moderntreasury.F(moderntreasury.InvoiceUpdateParamsPaymentMethodUi),
			PaymentType:                moderntreasury.F(moderntreasury.PaymentOrderTypeACH),
			ReceivingAccountID:         moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			RecipientEmail:             moderntreasury.F("string"),
			RecipientName:              moderntreasury.F("string"),
			Status:                     moderntreasury.F("string"),
			VirtualAccountID:           moderntreasury.F("string"),
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

func TestInvoiceListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.Invoices.List(context.TODO(), moderntreasury.InvoiceListParams{
		AfterCursor: moderntreasury.F("string"),
		PerPage:     moderntreasury.F(int64(0)),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceAddPaymentOrder(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithOrganizationID("my-organization-ID"),
	)
	err := client.Invoices.AddPaymentOrder(
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
