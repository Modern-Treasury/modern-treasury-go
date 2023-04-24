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

func TestCounterpartyNewWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Counterparties.New(context.TODO(), &requests.CounterpartyNewParams{Name: moderntreasury.F("string"), Accounts: moderntreasury.F([]requests.CounterpartyNewParamsAccounts{{AccountType: moderntreasury.F(requests.ExternalAccountTypeCash), PartyType: moderntreasury.F(requests.CounterpartyNewParamsAccountsPartyTypeBusiness), PartyAddress: moderntreasury.F(requests.CounterpartyNewParamsAccountsPartyAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), Name: moderntreasury.F("string"), AccountDetails: moderntreasury.F([]requests.CounterpartyNewParamsAccountsAccountDetails{{AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(requests.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban)}, {AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(requests.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban)}, {AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(requests.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban)}}), RoutingDetails: moderntreasury.F([]requests.CounterpartyNewParamsAccountsRoutingDetails{{RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(requests.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(requests.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH)}, {RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(requests.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(requests.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH)}, {RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(requests.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(requests.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH)}}), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), PartyName: moderntreasury.F("string"), PartyIdentifier: moderntreasury.F("string"), PlaidProcessorToken: moderntreasury.F("string"), ContactDetails: moderntreasury.F([]requests.CounterpartyNewParamsAccountsContactDetails{{ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail)}, {ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail)}, {ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail)}})}, {AccountType: moderntreasury.F(requests.ExternalAccountTypeCash), PartyType: moderntreasury.F(requests.CounterpartyNewParamsAccountsPartyTypeBusiness), PartyAddress: moderntreasury.F(requests.CounterpartyNewParamsAccountsPartyAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), Name: moderntreasury.F("string"), AccountDetails: moderntreasury.F([]requests.CounterpartyNewParamsAccountsAccountDetails{{AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(requests.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban)}, {AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(requests.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban)}, {AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(requests.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban)}}), RoutingDetails: moderntreasury.F([]requests.CounterpartyNewParamsAccountsRoutingDetails{{RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(requests.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(requests.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH)}, {RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(requests.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(requests.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH)}, {RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(requests.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(requests.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH)}}), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), PartyName: moderntreasury.F("string"), PartyIdentifier: moderntreasury.F("string"), PlaidProcessorToken: moderntreasury.F("string"), ContactDetails: moderntreasury.F([]requests.CounterpartyNewParamsAccountsContactDetails{{ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail)}, {ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail)}, {ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail)}})}, {AccountType: moderntreasury.F(requests.ExternalAccountTypeCash), PartyType: moderntreasury.F(requests.CounterpartyNewParamsAccountsPartyTypeBusiness), PartyAddress: moderntreasury.F(requests.CounterpartyNewParamsAccountsPartyAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), Name: moderntreasury.F("string"), AccountDetails: moderntreasury.F([]requests.CounterpartyNewParamsAccountsAccountDetails{{AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(requests.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban)}, {AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(requests.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban)}, {AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(requests.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban)}}), RoutingDetails: moderntreasury.F([]requests.CounterpartyNewParamsAccountsRoutingDetails{{RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(requests.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(requests.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH)}, {RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(requests.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(requests.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH)}, {RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(requests.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba), PaymentType: moderntreasury.F(requests.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH)}}), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), PartyName: moderntreasury.F("string"), PartyIdentifier: moderntreasury.F("string"), PlaidProcessorToken: moderntreasury.F("string"), ContactDetails: moderntreasury.F([]requests.CounterpartyNewParamsAccountsContactDetails{{ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail)}, {ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail)}, {ContactIdentifier: moderntreasury.F("string"), ContactIdentifierType: moderntreasury.F(requests.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail)}})}}), Email: moderntreasury.F("dev@stainlessapi.com"), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), SendRemittanceAdvice: moderntreasury.F(true), Accounting: moderntreasury.F(requests.CounterpartyNewParamsAccounting{Type: moderntreasury.F(requests.CounterpartyNewParamsAccountingTypeCustomer)}), LedgerType: moderntreasury.F(requests.CounterpartyNewParamsLedgerTypeCustomer), TaxpayerIdentifier: moderntreasury.F("string")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCounterpartyGet(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Counterparties.Get(
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

func TestCounterpartyUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Counterparties.Update(
		context.TODO(),
		"string",
		&requests.CounterpartyUpdateParams{Name: moderntreasury.F("string"), Email: moderntreasury.F("dev@stainlessapi.com"), Metadata: moderntreasury.F(map[string]string{"foo": "string"}), SendRemittanceAdvice: moderntreasury.F(true), TaxpayerIdentifier: moderntreasury.F("string")},
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

func TestCounterpartyListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Counterparties.List(context.TODO(), &requests.CounterpartyListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0)), Name: moderntreasury.F("string"), Email: moderntreasury.F("dev@stainlessapi.com"), Metadata: moderntreasury.F(map[string]string{"foo": "string"}), CreatedAtLowerBound: moderntreasury.F(time.Now()), CreatedAtUpperBound: moderntreasury.F(time.Now())})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCounterpartyDelete(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	err := c.Counterparties.Delete(
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

func TestCounterpartyCollectAccountWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Counterparties.CollectAccount(
		context.TODO(),
		"string",
		&requests.CounterpartyCollectAccountParams{Direction: moderntreasury.F(requests.CounterpartyCollectAccountParamsDirectionCredit), SendEmail: moderntreasury.F(true), Fields: moderntreasury.F([]requests.CounterpartyCollectAccountParamsFields{requests.CounterpartyCollectAccountParamsFieldsName, requests.CounterpartyCollectAccountParamsFieldsName, requests.CounterpartyCollectAccountParamsFieldsName}), CustomRedirect: moderntreasury.F("https://example.com")},
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
