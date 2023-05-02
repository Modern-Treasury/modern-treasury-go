package moderntreasury_test

import (
	"context"
	"errors"
	"testing"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestInternalAccountNewWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.InternalAccounts.New(context.TODO(), moderntreasury.InternalAccountNewParams{ConnectionID: moderntreasury.F("string"), Name: moderntreasury.F("string"), PartyName: moderntreasury.F("string"), PartyAddress: moderntreasury.F(moderntreasury.InternalAccountNewParamsPartyAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}), Currency: moderntreasury.F(moderntreasury.InternalAccountNewParamsCurrencyUsd), EntityID: moderntreasury.F("string"), ParentAccountID: moderntreasury.F("string"), CounterpartyID: moderntreasury.F("string")})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInternalAccountGet(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.InternalAccounts.Get(
		context.TODO(),
		"string",
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInternalAccountUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.InternalAccounts.Update(
		context.TODO(),
		"string",
		moderntreasury.InternalAccountUpdateParams{Name: moderntreasury.F("string"), Metadata: moderntreasury.F(map[string]string{"foo": "string"}), ParentAccountID: moderntreasury.F("string"), CounterpartyID: moderntreasury.F("string")},
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInternalAccountListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.InternalAccounts.List(context.TODO(), moderntreasury.InternalAccountListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0)), Currency: moderntreasury.F(shared.CurrencyAed), CounterpartyID: moderntreasury.F("string"), PaymentType: moderntreasury.F(moderntreasury.InternalAccountListParamsPaymentTypeACH), PaymentDirection: moderntreasury.F(moderntreasury.InternalAccountListParamsPaymentDirectionCredit), Metadata: moderntreasury.F(map[string]string{"foo": "string"})})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
