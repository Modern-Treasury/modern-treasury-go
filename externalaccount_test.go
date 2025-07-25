// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Modern-Treasury/modern-treasury-go/v2"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

func TestExternalAccountNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ExternalAccounts.New(context.TODO(), moderntreasury.ExternalAccountNewParams{
		CounterpartyID:  moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		QueryExternalID: moderntreasury.F("external_id"),
		AccountDetails: moderntreasury.F([]moderntreasury.ExternalAccountNewParamsAccountDetail{{
			AccountNumber:     moderntreasury.F("account_number"),
			AccountNumberType: moderntreasury.F(moderntreasury.ExternalAccountNewParamsAccountDetailsAccountNumberTypeAuNumber),
		}}),
		AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeBaseWallet),
		ContactDetails: moderntreasury.F([]moderntreasury.ContactDetailCreateRequestParam{{
			ContactIdentifier:     moderntreasury.F("contact_identifier"),
			ContactIdentifierType: moderntreasury.F(moderntreasury.ContactDetailCreateRequestContactIdentifierTypeEmail),
		}}),
		BodyExternalID: moderntreasury.F("external_id"),
		LedgerAccount: moderntreasury.F(shared.LedgerAccountCreateRequestParam{
			Currency:                 moderntreasury.F("currency"),
			LedgerID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:                     moderntreasury.F("name"),
			NormalBalance:            moderntreasury.F(shared.TransactionDirectionCredit),
			CurrencyExponent:         moderntreasury.F(int64(0)),
			Description:              moderntreasury.F("description"),
			LedgerAccountCategoryIDs: moderntreasury.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			LedgerableID:             moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LedgerableType:           moderntreasury.F(shared.LedgerAccountCreateRequestLedgerableTypeCounterparty),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
		}),
		Metadata: moderntreasury.F(map[string]string{
			"key":    "value",
			"foo":    "bar",
			"modern": "treasury",
		}),
		Name: moderntreasury.F("name"),
		PartyAddress: moderntreasury.F(shared.AddressRequestParam{
			Country:    moderntreasury.F("country"),
			Line1:      moderntreasury.F("line1"),
			Line2:      moderntreasury.F("line2"),
			Locality:   moderntreasury.F("locality"),
			PostalCode: moderntreasury.F("postal_code"),
			Region:     moderntreasury.F("region"),
		}),
		PartyIdentifier:     moderntreasury.F("party_identifier"),
		PartyName:           moderntreasury.F("party_name"),
		PartyType:           moderntreasury.F(moderntreasury.ExternalAccountNewParamsPartyTypeBusiness),
		PlaidProcessorToken: moderntreasury.F("plaid_processor_token"),
		RoutingDetails: moderntreasury.F([]moderntreasury.ExternalAccountNewParamsRoutingDetail{{
			RoutingNumber:     moderntreasury.F("routing_number"),
			RoutingNumberType: moderntreasury.F(moderntreasury.ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeAba),
			PaymentType:       moderntreasury.F(moderntreasury.ExternalAccountNewParamsRoutingDetailsPaymentTypeACH),
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

func TestExternalAccountGet(t *testing.T) {
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
	_, err := client.ExternalAccounts.Get(context.TODO(), "id")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExternalAccountUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ExternalAccounts.Update(
		context.TODO(),
		"id",
		moderntreasury.ExternalAccountUpdateParams{
			AccountType:    moderntreasury.F(moderntreasury.ExternalAccountTypeBaseWallet),
			CounterpartyID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Metadata: moderntreasury.F(map[string]string{
				"foo": "string",
			}),
			Name: moderntreasury.F("name"),
			PartyAddress: moderntreasury.F(shared.AddressRequestParam{
				Country:    moderntreasury.F("country"),
				Line1:      moderntreasury.F("line1"),
				Line2:      moderntreasury.F("line2"),
				Locality:   moderntreasury.F("locality"),
				PostalCode: moderntreasury.F("postal_code"),
				Region:     moderntreasury.F("region"),
			}),
			PartyName: moderntreasury.F("party_name"),
			PartyType: moderntreasury.F(moderntreasury.ExternalAccountUpdateParamsPartyTypeBusiness),
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

func TestExternalAccountListWithOptionalParams(t *testing.T) {
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
	_, err := client.ExternalAccounts.List(context.TODO(), moderntreasury.ExternalAccountListParams{
		AfterCursor:    moderntreasury.F("after_cursor"),
		CounterpartyID: moderntreasury.F("counterparty_id"),
		ExternalID:     moderntreasury.F("external_id"),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		PartyName: moderntreasury.F("party_name"),
		PerPage:   moderntreasury.F(int64(0)),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExternalAccountDelete(t *testing.T) {
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
	err := client.ExternalAccounts.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExternalAccountCompleteVerificationWithOptionalParams(t *testing.T) {
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
	_, err := client.ExternalAccounts.CompleteVerification(
		context.TODO(),
		"id",
		moderntreasury.ExternalAccountCompleteVerificationParams{
			Amounts: moderntreasury.F([]int64{int64(2), int64(4)}),
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

func TestExternalAccountVerifyWithOptionalParams(t *testing.T) {
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
	_, err := client.ExternalAccounts.Verify(
		context.TODO(),
		"id",
		moderntreasury.ExternalAccountVerifyParams{
			OriginatingAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			PaymentType:          moderntreasury.F(moderntreasury.ExternalAccountVerifyParamsPaymentTypeACH),
			Currency:             moderntreasury.F(shared.CurrencyAed),
			FallbackType:         moderntreasury.F(moderntreasury.ExternalAccountVerifyParamsFallbackTypeACH),
			Priority:             moderntreasury.F(moderntreasury.ExternalAccountVerifyParamsPriorityHigh),
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
