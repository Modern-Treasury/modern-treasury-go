// File generated from our OpenAPI spec by Stainless.

package moderntreasury_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Modern-Treasury/modern-treasury-go/v2"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

func TestVirtualAccountNewWithOptionalParams(t *testing.T) {
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
	_, err := client.VirtualAccounts.New(context.TODO(), moderntreasury.VirtualAccountNewParams{
		InternalAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Name:              moderntreasury.F("string"),
		AccountDetails: moderntreasury.F([]moderntreasury.VirtualAccountNewParamsAccountDetail{{
			AccountNumber:     moderntreasury.F("string"),
			AccountNumberType: moderntreasury.F(moderntreasury.VirtualAccountNewParamsAccountDetailsAccountNumberTypeClabe),
		}, {
			AccountNumber:     moderntreasury.F("string"),
			AccountNumberType: moderntreasury.F(moderntreasury.VirtualAccountNewParamsAccountDetailsAccountNumberTypeClabe),
		}, {
			AccountNumber:     moderntreasury.F("string"),
			AccountNumberType: moderntreasury.F(moderntreasury.VirtualAccountNewParamsAccountDetailsAccountNumberTypeClabe),
		}}),
		CounterpartyID:        moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CreditLedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		DebitLedgerAccountID:  moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Description:           moderntreasury.F("string"),
		LedgerAccount: moderntreasury.F(moderntreasury.VirtualAccountNewParamsLedgerAccount{
			Name:                     moderntreasury.F("string"),
			Description:              moderntreasury.F("string"),
			NormalBalance:            moderntreasury.F(shared.TransactionDirectionCredit),
			LedgerID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Currency:                 moderntreasury.F("string"),
			CurrencyExponent:         moderntreasury.F(int64(0)),
			LedgerAccountCategoryIDs: moderntreasury.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			LedgerableID:             moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LedgerableType:           moderntreasury.F(moderntreasury.VirtualAccountNewParamsLedgerAccountLedgerableTypeExternalAccount),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
		}),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		RoutingDetails: moderntreasury.F([]moderntreasury.VirtualAccountNewParamsRoutingDetail{{
			RoutingNumber:     moderntreasury.F("string"),
			RoutingNumberType: moderntreasury.F(moderntreasury.VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeAba),
			PaymentType:       moderntreasury.F(moderntreasury.VirtualAccountNewParamsRoutingDetailsPaymentTypeACH),
		}, {
			RoutingNumber:     moderntreasury.F("string"),
			RoutingNumberType: moderntreasury.F(moderntreasury.VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeAba),
			PaymentType:       moderntreasury.F(moderntreasury.VirtualAccountNewParamsRoutingDetailsPaymentTypeACH),
		}, {
			RoutingNumber:     moderntreasury.F("string"),
			RoutingNumberType: moderntreasury.F(moderntreasury.VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeAba),
			PaymentType:       moderntreasury.F(moderntreasury.VirtualAccountNewParamsRoutingDetailsPaymentTypeACH),
		}}),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVirtualAccountGet(t *testing.T) {
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
	_, err := client.VirtualAccounts.Get(context.TODO(), "string")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVirtualAccountUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.VirtualAccounts.Update(
		context.TODO(),
		"string",
		moderntreasury.VirtualAccountUpdateParams{
			CounterpartyID:  moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Metadata: moderntreasury.F(map[string]string{
				"foo": "string",
			}),
			Name: moderntreasury.F("string"),
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

func TestVirtualAccountListWithOptionalParams(t *testing.T) {
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
	_, err := client.VirtualAccounts.List(context.TODO(), moderntreasury.VirtualAccountListParams{
		AfterCursor:       moderntreasury.F("string"),
		CounterpartyID:    moderntreasury.F("string"),
		InternalAccountID: moderntreasury.F("string"),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		PerPage: moderntreasury.F(int64(0)),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVirtualAccountDelete(t *testing.T) {
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
	_, err := client.VirtualAccounts.Delete(context.TODO(), "string")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
