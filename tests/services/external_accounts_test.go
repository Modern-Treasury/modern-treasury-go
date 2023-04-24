package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/core"
	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
)

func TestExternalAccountNewWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.ExternalAccounts.New(context.TODO(), &requests.ExternalAccountNewParams{AccountType: moderntreasury.F(requests.ExternalAccountTypeCash), PartyType: moderntreasury.F(requests.ExternalAccountNewParamsPartyTypeBusiness), PartyAddress: moderntreasury.F(requests.ExternalAccountNewParamsPartyAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), Name: moderntreasury.F("string"), CounterpartyID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), AccountDetails: moderntreasury.F([]requests.ExternalAccountNewParamsAccountDetails{{AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(requests.ExternalAccountNewParamsAccountDetailsAccountNumberTypeIban)}, {AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(requests.ExternalAccountNewParamsAccountDetailsAccountNumberTypeIban)}, {AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(requests.ExternalAccountNewParamsAccountDetailsAccountNumberTypeIban)}}), RoutingDetails: moderntreasury.F([]requests.ExternalAccountNewParamsRoutingDetails{{RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(requests.ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(requests.ExternalAccountNewParamsRoutingDetailsPaymentTypeACH)}, {RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(requests.ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(requests.ExternalAccountNewParamsRoutingDetailsPaymentTypeACH)}, {RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(requests.ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(requests.ExternalAccountNewParamsRoutingDetailsPaymentTypeACH)}}), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), PartyName: moderntreasury.F("string"), PartyIdentifier: moderntreasury.F("string"), PlaidProcessorToken: moderntreasury.F("string"), ContactDetails: moderntreasury.F([]requests.ExternalAccountNewParamsContactDetails{{ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.ExternalAccountNewParamsContactDetailsContactIdentifierTypeEmail)}, {ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.ExternalAccountNewParamsContactDetailsContactIdentifierTypeEmail)}, {ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.ExternalAccountNewParamsContactDetailsContactIdentifierTypeEmail)}})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExternalAccountGet(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.ExternalAccounts.Get(
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

func TestExternalAccountUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.ExternalAccounts.Update(
		context.TODO(),
		"string",
		&requests.ExternalAccountUpdateParams{PartyType: moderntreasury.F(requests.ExternalAccountUpdateParamsPartyTypeBusiness), AccountType: moderntreasury.F(requests.ExternalAccountTypeCash), CounterpartyID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Name: moderntreasury.F("string"), PartyName: moderntreasury.F("string"), PartyAddress: moderntreasury.F(requests.ExternalAccountUpdateParamsPartyAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), Metadata: moderntreasury.F(map[string]string{"foo": "string"})},
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

func TestExternalAccountListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.ExternalAccounts.List(context.TODO(), &requests.ExternalAccountListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0)), PartyName: moderntreasury.F("string"), CounterpartyID: moderntreasury.F("string"), Metadata: moderntreasury.F(map[string]string{"foo": "string"})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExternalAccountDelete(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	err := c.ExternalAccounts.Delete(
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

func TestExternalAccountCompleteVerificationWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.ExternalAccounts.CompleteVerification(
		context.TODO(),
		"string",
		&requests.ExternalAccountCompleteVerificationParams{Amounts: moderntreasury.F([]int64{int64(0), int64(0)})},
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

func TestExternalAccountVerifyWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.ExternalAccounts.Verify(
		context.TODO(),
		"string",
		&requests.ExternalAccountVerifyParams{OriginatingAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), PaymentType: moderntreasury.F(requests.ExternalAccountVerifyParamsPaymentTypeACH), Currency: moderntreasury.F(requests.CurrencyAed)},
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
