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

func TestInternalAccountNewWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.InternalAccounts.New(context.TODO(), &requests.InternalAccountNewParams{ConnectionID: moderntreasury.F("string"), Name: moderntreasury.F("string"), PartyName: moderntreasury.F("string"), PartyAddress: moderntreasury.F(requests.InternalAccountNewParamsPartyAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), Currency: moderntreasury.F(requests.InternalAccountNewParamsCurrencyUsd), EntityID: moderntreasury.F("string"), ParentAccountID: moderntreasury.F("string"), CounterpartyID: moderntreasury.F("string")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInternalAccountGet(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.InternalAccounts.Get(
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

func TestInternalAccountUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.InternalAccounts.Update(
		context.TODO(),
		"string",
		&requests.InternalAccountUpdateParams{Name: moderntreasury.F("string"), Metadata: moderntreasury.F(map[string]string{"foo": "string"}), ParentAccountID: moderntreasury.F("string"), CounterpartyID: moderntreasury.F("string")},
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

func TestInternalAccountListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.InternalAccounts.List(context.TODO(), &requests.InternalAccountListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0)), Currency: moderntreasury.F(requests.CurrencyAed), CounterpartyID: moderntreasury.F("string"), PaymentType: moderntreasury.F(requests.InternalAccountListParamsPaymentTypeACH), PaymentDirection: moderntreasury.F(requests.InternalAccountListParamsPaymentDirectionCredit), Metadata: moderntreasury.F(map[string]string{"foo": "string"})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
