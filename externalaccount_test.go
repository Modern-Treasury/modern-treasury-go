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
		CounterpartyID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		AccountDetails: moderntreasury.F([]moderntreasury.ExternalAccountNewParamsAccountDetail{{
			AccountNumber:     moderntreasury.F("string"),
			AccountNumberType: moderntreasury.F(moderntreasury.ExternalAccountNewParamsAccountDetailsAccountNumberTypeIban),
		}, {
			AccountNumber:     moderntreasury.F("string"),
			AccountNumberType: moderntreasury.F(moderntreasury.ExternalAccountNewParamsAccountDetailsAccountNumberTypeIban),
		}, {
			AccountNumber:     moderntreasury.F("string"),
			AccountNumberType: moderntreasury.F(moderntreasury.ExternalAccountNewParamsAccountDetailsAccountNumberTypeIban),
		}}),
		AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
		ContactDetails: moderntreasury.F([]moderntreasury.ExternalAccountNewParamsContactDetail{{
			ContactIdentifier:     moderntreasury.F("string"),
			ContactIdentifierType: moderntreasury.F(moderntreasury.ExternalAccountNewParamsContactDetailsContactIdentifierTypeEmail),
		}, {
			ContactIdentifier:     moderntreasury.F("string"),
			ContactIdentifierType: moderntreasury.F(moderntreasury.ExternalAccountNewParamsContactDetailsContactIdentifierTypeEmail),
		}, {
			ContactIdentifier:     moderntreasury.F("string"),
			ContactIdentifierType: moderntreasury.F(moderntreasury.ExternalAccountNewParamsContactDetailsContactIdentifierTypeEmail),
		}}),
		LedgerAccount: moderntreasury.F(moderntreasury.ExternalAccountNewParamsLedgerAccount{
			Name:                     moderntreasury.F("string"),
			Description:              moderntreasury.F("string"),
			NormalBalance:            moderntreasury.F(shared.TransactionDirectionCredit),
			LedgerID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Currency:                 moderntreasury.F("string"),
			CurrencyExponent:         moderntreasury.F(int64(0)),
			LedgerAccountCategoryIDs: moderntreasury.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			LedgerableID:             moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LedgerableType:           moderntreasury.F(moderntreasury.ExternalAccountNewParamsLedgerAccountLedgerableTypeCounterparty),
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
		Name: moderntreasury.F("string"),
		PartyAddress: moderntreasury.F(moderntreasury.ExternalAccountNewParamsPartyAddress{
			Line1:      moderntreasury.F("string"),
			Line2:      moderntreasury.F("string"),
			Locality:   moderntreasury.F("string"),
			Region:     moderntreasury.F("string"),
			PostalCode: moderntreasury.F("string"),
			Country:    moderntreasury.F("string"),
		}),
		PartyIdentifier:     moderntreasury.F("string"),
		PartyName:           moderntreasury.F("string"),
		PartyType:           moderntreasury.F(moderntreasury.ExternalAccountNewParamsPartyTypeBusiness),
		PlaidProcessorToken: moderntreasury.F("string"),
		RoutingDetails: moderntreasury.F([]moderntreasury.ExternalAccountNewParamsRoutingDetail{{
			RoutingNumber:     moderntreasury.F("string"),
			RoutingNumberType: moderntreasury.F(moderntreasury.ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeAba),
			PaymentType:       moderntreasury.F(moderntreasury.ExternalAccountNewParamsRoutingDetailsPaymentTypeACH),
		}, {
			RoutingNumber:     moderntreasury.F("string"),
			RoutingNumberType: moderntreasury.F(moderntreasury.ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeAba),
			PaymentType:       moderntreasury.F(moderntreasury.ExternalAccountNewParamsRoutingDetailsPaymentTypeACH),
		}, {
			RoutingNumber:     moderntreasury.F("string"),
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
	_, err := client.ExternalAccounts.Get(context.TODO(), "string")
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
		"string",
		moderntreasury.ExternalAccountUpdateParams{
			AccountType:    moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
			CounterpartyID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Metadata: moderntreasury.F(map[string]string{
				"foo": "string",
			}),
			Name: moderntreasury.F("string"),
			PartyAddress: moderntreasury.F(moderntreasury.ExternalAccountUpdateParamsPartyAddress{
				Line1:      moderntreasury.F("string"),
				Line2:      moderntreasury.F("string"),
				Locality:   moderntreasury.F("string"),
				Region:     moderntreasury.F("string"),
				PostalCode: moderntreasury.F("string"),
				Country:    moderntreasury.F("string"),
			}),
			PartyName: moderntreasury.F("string"),
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
		AfterCursor:    moderntreasury.F("string"),
		CounterpartyID: moderntreasury.F("string"),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		PartyName: moderntreasury.F("string"),
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
	err := client.ExternalAccounts.Delete(context.TODO(), "string")
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
		"string",
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
		"string",
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
