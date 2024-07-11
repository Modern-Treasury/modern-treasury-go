// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
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
		CounterpartyID:       moderntreasury.F("counterparty_id"),
		DueDate:              moderntreasury.F(time.Now()),
		OriginatingAccountID: moderntreasury.F("originating_account_id"),
		ContactDetails: moderntreasury.F([]moderntreasury.InvoiceNewParamsContactDetail{{
			ID:                    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Object:                moderntreasury.F("object"),
			LiveMode:              moderntreasury.F(true),
			CreatedAt:             moderntreasury.F(time.Now()),
			UpdatedAt:             moderntreasury.F(time.Now()),
			DiscardedAt:           moderntreasury.F(time.Now()),
			ContactIdentifier:     moderntreasury.F("contact_identifier"),
			ContactIdentifierType: moderntreasury.F(moderntreasury.InvoiceNewParamsContactDetailsContactIdentifierTypeEmail),
		}, {
			ID:                    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Object:                moderntreasury.F("object"),
			LiveMode:              moderntreasury.F(true),
			CreatedAt:             moderntreasury.F(time.Now()),
			UpdatedAt:             moderntreasury.F(time.Now()),
			DiscardedAt:           moderntreasury.F(time.Now()),
			ContactIdentifier:     moderntreasury.F("contact_identifier"),
			ContactIdentifierType: moderntreasury.F(moderntreasury.InvoiceNewParamsContactDetailsContactIdentifierTypeEmail),
		}, {
			ID:                    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Object:                moderntreasury.F("object"),
			LiveMode:              moderntreasury.F(true),
			CreatedAt:             moderntreasury.F(time.Now()),
			UpdatedAt:             moderntreasury.F(time.Now()),
			DiscardedAt:           moderntreasury.F(time.Now()),
			ContactIdentifier:     moderntreasury.F("contact_identifier"),
			ContactIdentifierType: moderntreasury.F(moderntreasury.InvoiceNewParamsContactDetailsContactIdentifierTypeEmail),
		}}),
		CounterpartyBillingAddress: moderntreasury.F(moderntreasury.InvoiceNewParamsCounterpartyBillingAddress{
			Line1:      moderntreasury.F("line1"),
			Line2:      moderntreasury.F("line2"),
			Locality:   moderntreasury.F("locality"),
			Region:     moderntreasury.F("region"),
			PostalCode: moderntreasury.F("postal_code"),
			Country:    moderntreasury.F("country"),
		}),
		CounterpartyShippingAddress: moderntreasury.F(moderntreasury.InvoiceNewParamsCounterpartyShippingAddress{
			Line1:      moderntreasury.F("line1"),
			Line2:      moderntreasury.F("line2"),
			Locality:   moderntreasury.F("locality"),
			Region:     moderntreasury.F("region"),
			PostalCode: moderntreasury.F("postal_code"),
			Country:    moderntreasury.F("country"),
		}),
		Currency:              moderntreasury.F(shared.CurrencyAed),
		Description:           moderntreasury.F("description"),
		FallbackPaymentMethod: moderntreasury.F("fallback_payment_method"),
		IngestLedgerEntries:   moderntreasury.F(true),
		InvoiceLineItems: moderntreasury.F([]moderntreasury.InvoiceNewParamsInvoiceLineItem{{
			Name:              moderntreasury.F("name"),
			Description:       moderntreasury.F("description"),
			Quantity:          moderntreasury.F(int64(0)),
			UnitAmount:        moderntreasury.F(int64(0)),
			UnitAmountDecimal: moderntreasury.F("unit_amount_decimal"),
			Direction:         moderntreasury.F("direction"),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
		}, {
			Name:              moderntreasury.F("name"),
			Description:       moderntreasury.F("description"),
			Quantity:          moderntreasury.F(int64(0)),
			UnitAmount:        moderntreasury.F(int64(0)),
			UnitAmountDecimal: moderntreasury.F("unit_amount_decimal"),
			Direction:         moderntreasury.F("direction"),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
		}, {
			Name:              moderntreasury.F("name"),
			Description:       moderntreasury.F("description"),
			Quantity:          moderntreasury.F(int64(0)),
			UnitAmount:        moderntreasury.F(int64(0)),
			UnitAmountDecimal: moderntreasury.F("unit_amount_decimal"),
			Direction:         moderntreasury.F("direction"),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
		}}),
		InvoicerAddress: moderntreasury.F(moderntreasury.InvoiceNewParamsInvoicerAddress{
			Line1:      moderntreasury.F("line1"),
			Line2:      moderntreasury.F("line2"),
			Locality:   moderntreasury.F("locality"),
			Region:     moderntreasury.F("region"),
			PostalCode: moderntreasury.F("postal_code"),
			Country:    moderntreasury.F("country"),
		}),
		LedgerAccountSettlementID:  moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		NotificationEmailAddresses: moderntreasury.F([]string{"string", "string", "string"}),
		NotificationsEnabled:       moderntreasury.F(true),
		PaymentEffectiveDate:       moderntreasury.F(time.Now()),
		PaymentMethod:              moderntreasury.F(moderntreasury.InvoiceNewParamsPaymentMethodUi),
		PaymentType:                moderntreasury.F(moderntreasury.PaymentOrderTypeACH),
		ReceivingAccountID:         moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		RecipientEmail:             moderntreasury.F("recipient_email"),
		RecipientName:              moderntreasury.F("recipient_name"),
		RemindAfterOverdueDays:     moderntreasury.F([]int64{int64(0), int64(0), int64(0)}),
		VirtualAccountID:           moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
	_, err := client.Invoices.Get(context.TODO(), "id")
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
		"id",
		moderntreasury.InvoiceUpdateParams{
			ContactDetails: moderntreasury.F([]moderntreasury.InvoiceUpdateParamsContactDetail{{
				ID:                    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Object:                moderntreasury.F("object"),
				LiveMode:              moderntreasury.F(true),
				CreatedAt:             moderntreasury.F(time.Now()),
				UpdatedAt:             moderntreasury.F(time.Now()),
				DiscardedAt:           moderntreasury.F(time.Now()),
				ContactIdentifier:     moderntreasury.F("contact_identifier"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.InvoiceUpdateParamsContactDetailsContactIdentifierTypeEmail),
			}, {
				ID:                    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Object:                moderntreasury.F("object"),
				LiveMode:              moderntreasury.F(true),
				CreatedAt:             moderntreasury.F(time.Now()),
				UpdatedAt:             moderntreasury.F(time.Now()),
				DiscardedAt:           moderntreasury.F(time.Now()),
				ContactIdentifier:     moderntreasury.F("contact_identifier"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.InvoiceUpdateParamsContactDetailsContactIdentifierTypeEmail),
			}, {
				ID:                    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Object:                moderntreasury.F("object"),
				LiveMode:              moderntreasury.F(true),
				CreatedAt:             moderntreasury.F(time.Now()),
				UpdatedAt:             moderntreasury.F(time.Now()),
				DiscardedAt:           moderntreasury.F(time.Now()),
				ContactIdentifier:     moderntreasury.F("contact_identifier"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.InvoiceUpdateParamsContactDetailsContactIdentifierTypeEmail),
			}}),
			CounterpartyBillingAddress: moderntreasury.F(moderntreasury.InvoiceUpdateParamsCounterpartyBillingAddress{
				Line1:      moderntreasury.F("line1"),
				Line2:      moderntreasury.F("line2"),
				Locality:   moderntreasury.F("locality"),
				Region:     moderntreasury.F("region"),
				PostalCode: moderntreasury.F("postal_code"),
				Country:    moderntreasury.F("country"),
			}),
			CounterpartyID: moderntreasury.F("counterparty_id"),
			CounterpartyShippingAddress: moderntreasury.F(moderntreasury.InvoiceUpdateParamsCounterpartyShippingAddress{
				Line1:      moderntreasury.F("line1"),
				Line2:      moderntreasury.F("line2"),
				Locality:   moderntreasury.F("locality"),
				Region:     moderntreasury.F("region"),
				PostalCode: moderntreasury.F("postal_code"),
				Country:    moderntreasury.F("country"),
			}),
			Currency:              moderntreasury.F(shared.CurrencyAed),
			Description:           moderntreasury.F("description"),
			DueDate:               moderntreasury.F(time.Now()),
			FallbackPaymentMethod: moderntreasury.F("fallback_payment_method"),
			IngestLedgerEntries:   moderntreasury.F(true),
			InvoiceLineItems: moderntreasury.F([]moderntreasury.InvoiceUpdateParamsInvoiceLineItem{{
				Name:              moderntreasury.F("name"),
				Description:       moderntreasury.F("description"),
				Quantity:          moderntreasury.F(int64(0)),
				UnitAmount:        moderntreasury.F(int64(0)),
				UnitAmountDecimal: moderntreasury.F("unit_amount_decimal"),
				Direction:         moderntreasury.F("direction"),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
			}, {
				Name:              moderntreasury.F("name"),
				Description:       moderntreasury.F("description"),
				Quantity:          moderntreasury.F(int64(0)),
				UnitAmount:        moderntreasury.F(int64(0)),
				UnitAmountDecimal: moderntreasury.F("unit_amount_decimal"),
				Direction:         moderntreasury.F("direction"),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
			}, {
				Name:              moderntreasury.F("name"),
				Description:       moderntreasury.F("description"),
				Quantity:          moderntreasury.F(int64(0)),
				UnitAmount:        moderntreasury.F(int64(0)),
				UnitAmountDecimal: moderntreasury.F("unit_amount_decimal"),
				Direction:         moderntreasury.F("direction"),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
			}}),
			InvoicerAddress: moderntreasury.F(moderntreasury.InvoiceUpdateParamsInvoicerAddress{
				Line1:      moderntreasury.F("line1"),
				Line2:      moderntreasury.F("line2"),
				Locality:   moderntreasury.F("locality"),
				Region:     moderntreasury.F("region"),
				PostalCode: moderntreasury.F("postal_code"),
				Country:    moderntreasury.F("country"),
			}),
			LedgerAccountSettlementID:  moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			NotificationEmailAddresses: moderntreasury.F([]string{"string", "string", "string"}),
			NotificationsEnabled:       moderntreasury.F(true),
			OriginatingAccountID:       moderntreasury.F("originating_account_id"),
			PaymentEffectiveDate:       moderntreasury.F(time.Now()),
			PaymentMethod:              moderntreasury.F(moderntreasury.InvoiceUpdateParamsPaymentMethodUi),
			PaymentType:                moderntreasury.F(moderntreasury.PaymentOrderTypeACH),
			ReceivingAccountID:         moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			RecipientEmail:             moderntreasury.F("recipient_email"),
			RecipientName:              moderntreasury.F("recipient_name"),
			RemindAfterOverdueDays:     moderntreasury.F([]int64{int64(0), int64(0), int64(0)}),
			Status:                     moderntreasury.F("status"),
			VirtualAccountID:           moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
		AfterCursor:       moderntreasury.F("after_cursor"),
		CounterpartyID:    moderntreasury.F("counterparty_id"),
		DueDateEnd:        moderntreasury.F(time.Now()),
		DueDateStart:      moderntreasury.F(time.Now()),
		ExpectedPaymentID: moderntreasury.F("expected_payment_id"),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		Number:               moderntreasury.F("number"),
		OriginatingAccountID: moderntreasury.F("originating_account_id"),
		PaymentOrderID:       moderntreasury.F("payment_order_id"),
		PerPage:              moderntreasury.F(int64(0)),
		Status:               moderntreasury.F(moderntreasury.InvoiceListParamsStatusDraft),
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
		"id",
		"payment_order_id",
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
