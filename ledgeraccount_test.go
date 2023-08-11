// File generated from our OpenAPI spec by Stainless.

package moderntreasury_test

import (
	"context"
	"errors"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestLedgerAccountNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.LedgerAccounts.New(context.TODO(), moderntreasury.LedgerAccountNewParams{
		Currency:         moderntreasury.F("string"),
		LedgerID:         moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Name:             moderntreasury.F("string"),
		NormalBalance:    moderntreasury.F(moderntreasury.LedgerAccountNewParamsNormalBalanceCredit),
		CurrencyExponent: moderntreasury.F(int64(0)),
		Description:      moderntreasury.F("string"),
		LedgerableID:     moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		LedgerableType:   moderntreasury.F(moderntreasury.LedgerAccountNewParamsLedgerableTypeExternalAccount),
		Metadata: moderntreasury.F(map[string]string{
			"key":    "value",
			"foo":    "bar",
			"modern": "treasury",
		}),
		IdempotencyKey: moderntreasury.F("string"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerAccountGetWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.LedgerAccounts.Get(
		context.TODO(),
		"string",
		moderntreasury.LedgerAccountGetParams{
			Balances: moderntreasury.F(moderntreasury.LedgerAccountGetParamsBalances{
				AsOfDate:              moderntreasury.F(time.Now()),
				EffectiveAt:           moderntreasury.F(time.Now()),
				EffectiveAtLowerBound: moderntreasury.F(time.Now()),
				EffectiveAtUpperBound: moderntreasury.F(time.Now()),
				AsOfLockVersion:       moderntreasury.F(int64(0)),
			}),
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

func TestLedgerAccountUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.LedgerAccounts.Update(
		context.TODO(),
		"string",
		moderntreasury.LedgerAccountUpdateParams{
			Description: moderntreasury.F("string"),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
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

func TestLedgerAccountListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.LedgerAccounts.List(context.TODO(), moderntreasury.LedgerAccountListParams{
		ID:          moderntreasury.F([]string{"string", "string", "string"}),
		AfterCursor: moderntreasury.F("string"),
		Balances: moderntreasury.F(moderntreasury.LedgerAccountListParamsBalances{
			AsOfDate:              moderntreasury.F(time.Now()),
			EffectiveAt:           moderntreasury.F(time.Now()),
			EffectiveAtLowerBound: moderntreasury.F(time.Now()),
			EffectiveAtUpperBound: moderntreasury.F(time.Now()),
		}),
		CreatedAt: moderntreasury.F(map[string]time.Time{
			"foo": time.Now(),
		}),
		Currency:                moderntreasury.F("string"),
		LedgerAccountCategoryID: moderntreasury.F("string"),
		LedgerID:                moderntreasury.F("string"),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		Name:    moderntreasury.F("string"),
		PerPage: moderntreasury.F(int64(0)),
		UpdatedAt: moderntreasury.F(map[string]time.Time{
			"foo": time.Now(),
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

func TestLedgerAccountDelete(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.LedgerAccounts.Delete(context.TODO(), "string")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
