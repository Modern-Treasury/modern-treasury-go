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

func TestInternalAccountNewWithOptionalParams(t *testing.T) {
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
	_, err := client.InternalAccounts.New(context.TODO(), moderntreasury.InternalAccountNewParams{
		ConnectionID: moderntreasury.F("connection_id"),
		Currency:     moderntreasury.F(moderntreasury.InternalAccountNewParamsCurrencyUsd),
		Name:         moderntreasury.F("name"),
		PartyName:    moderntreasury.F("party_name"),
		AccountCapabilities: moderntreasury.F([]moderntreasury.InternalAccountNewParamsAccountCapability{{
			ID:          moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			CreatedAt:   moderntreasury.F(time.Now()),
			Direction:   moderntreasury.F(shared.TransactionDirectionCredit),
			DiscardedAt: moderntreasury.F(time.Now()),
			Identifier:  moderntreasury.F("identifier"),
			LiveMode:    moderntreasury.F(true),
			Object:      moderntreasury.F("object"),
			PaymentType: moderntreasury.F(moderntreasury.InternalAccountNewParamsAccountCapabilitiesPaymentTypeACH),
			UpdatedAt:   moderntreasury.F(time.Now()),
		}}),
		AccountType:     moderntreasury.F(moderntreasury.InternalAccountNewParamsAccountTypeChecking),
		CounterpartyID:  moderntreasury.F("counterparty_id"),
		LegalEntityID:   moderntreasury.F("legal_entity_id"),
		ParentAccountID: moderntreasury.F("parent_account_id"),
		PartyAddress: moderntreasury.F(moderntreasury.InternalAccountNewParamsPartyAddress{
			Country:    moderntreasury.F("country"),
			Line1:      moderntreasury.F("line1"),
			Locality:   moderntreasury.F("locality"),
			PostalCode: moderntreasury.F("postal_code"),
			Region:     moderntreasury.F("region"),
			Line2:      moderntreasury.F("line2"),
		}),
		VendorAttributes: moderntreasury.F(map[string]string{
			"key":    "value",
			"foo":    "bar",
			"modern": "treasury",
		}),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInternalAccountGet(t *testing.T) {
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
	_, err := client.InternalAccounts.Get(context.TODO(), "id")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInternalAccountUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.InternalAccounts.Update(
		context.TODO(),
		"id",
		moderntreasury.InternalAccountUpdateParams{
			CounterpartyID:  moderntreasury.F("counterparty_id"),
			LedgerAccountID: moderntreasury.F("ledger_account_id"),
			Metadata: moderntreasury.F(map[string]string{
				"foo": "string",
			}),
			Name:            moderntreasury.F("name"),
			ParentAccountID: moderntreasury.F("parent_account_id"),
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

func TestInternalAccountListWithOptionalParams(t *testing.T) {
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
	_, err := client.InternalAccounts.List(context.TODO(), moderntreasury.InternalAccountListParams{
		AfterCursor:    moderntreasury.F("after_cursor"),
		CounterpartyID: moderntreasury.F("counterparty_id"),
		Currency:       moderntreasury.F(shared.CurrencyAed),
		LegalEntityID:  moderntreasury.F("legal_entity_id"),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		PaymentDirection: moderntreasury.F(shared.TransactionDirectionCredit),
		PaymentType:      moderntreasury.F(moderntreasury.InternalAccountListParamsPaymentTypeACH),
		PerPage:          moderntreasury.F(int64(0)),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
