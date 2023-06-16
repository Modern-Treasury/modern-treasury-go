// File generated from our OpenAPI spec by Stainless.

package moderntreasury_test

import (
	"context"
	"errors"
	"testing"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestInternalAccountNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.InternalAccounts.New(context.TODO(), moderntreasury.InternalAccountNewParams{
		ConnectionID:    moderntreasury.F("string"),
		Currency:        moderntreasury.F(moderntreasury.InternalAccountNewParamsCurrencyUsd),
		Name:            moderntreasury.F("string"),
		PartyName:       moderntreasury.F("string"),
		CounterpartyID:  moderntreasury.F("string"),
		EntityID:        moderntreasury.F("string"),
		ParentAccountID: moderntreasury.F("string"),
		PartyAddress:    moderntreasury.F(moderntreasury.InternalAccountNewParamsPartyAddress{Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")}),
		IdempotencyKey:  moderntreasury.F("string"),
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
	if !testutil.CheckTestServer(t) {
		return
	}
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.InternalAccounts.Get(context.TODO(), "string")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInternalAccountUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.InternalAccounts.Update(
		context.TODO(),
		"string",
		moderntreasury.InternalAccountUpdateParams{
			CounterpartyID:  moderntreasury.F("string"),
			Metadata:        moderntreasury.F(map[string]string{"foo": "string"}),
			Name:            moderntreasury.F("string"),
			ParentAccountID: moderntreasury.F("string"),
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
	if !testutil.CheckTestServer(t) {
		return
	}
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.InternalAccounts.List(context.TODO(), moderntreasury.InternalAccountListParams{
		AfterCursor:      moderntreasury.F("string"),
		CounterpartyID:   moderntreasury.F("string"),
		Currency:         moderntreasury.F(shared.CurrencyAed),
		Metadata:         moderntreasury.F(map[string]string{"foo": "string"}),
		PaymentDirection: moderntreasury.F(moderntreasury.InternalAccountListParamsPaymentDirectionCredit),
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
