// File generated from our OpenAPI spec by Stainless.

package moderntreasury_test

import (
	"context"
	"errors"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestLedgerTransactionNewWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerTransactions.New(context.TODO(), moderntreasury.LedgerTransactionNewParams{
		EffectiveDate:  moderntreasury.F(time.Now()),
		LedgerEntries:  moderntreasury.F([]moderntreasury.LedgerTransactionNewParamsLedgerEntries{{Amount: moderntreasury.F(int64(0)), Direction: moderntreasury.F(moderntreasury.LedgerTransactionNewParamsLedgerEntriesDirectionCredit), LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), LockVersion: moderntreasury.F(int64(0)), PendingBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), PostedBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), AvailableBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), ShowResultingLedgerAccountBalances: moderntreasury.F(true)}, {Amount: moderntreasury.F(int64(0)), Direction: moderntreasury.F(moderntreasury.LedgerTransactionNewParamsLedgerEntriesDirectionCredit), LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), LockVersion: moderntreasury.F(int64(0)), PendingBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), PostedBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), AvailableBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), ShowResultingLedgerAccountBalances: moderntreasury.F(true)}, {Amount: moderntreasury.F(int64(0)), Direction: moderntreasury.F(moderntreasury.LedgerTransactionNewParamsLedgerEntriesDirectionCredit), LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), LockVersion: moderntreasury.F(int64(0)), PendingBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), PostedBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), AvailableBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), ShowResultingLedgerAccountBalances: moderntreasury.F(true)}}),
		Description:    moderntreasury.F("string"),
		ExternalID:     moderntreasury.F("string"),
		LedgerableID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		LedgerableType: moderntreasury.F(moderntreasury.LedgerTransactionNewParamsLedgerableTypeCounterparty),
		Metadata:       moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}),
		Status:         moderntreasury.F(moderntreasury.LedgerTransactionNewParamsStatusArchived),
		IdempotencyKey: moderntreasury.F("string"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerTransactionGet(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerTransactions.Get(
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

func TestLedgerTransactionUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerTransactions.Update(
		context.TODO(),
		"string",
		moderntreasury.LedgerTransactionUpdateParams{
			Description:   moderntreasury.F("string"),
			LedgerEntries: moderntreasury.F([]moderntreasury.LedgerTransactionUpdateParamsLedgerEntries{{Amount: moderntreasury.F(int64(0)), Direction: moderntreasury.F(moderntreasury.LedgerTransactionUpdateParamsLedgerEntriesDirectionCredit), LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), LockVersion: moderntreasury.F(int64(0)), PendingBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), PostedBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), AvailableBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), ShowResultingLedgerAccountBalances: moderntreasury.F(true)}, {Amount: moderntreasury.F(int64(0)), Direction: moderntreasury.F(moderntreasury.LedgerTransactionUpdateParamsLedgerEntriesDirectionCredit), LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), LockVersion: moderntreasury.F(int64(0)), PendingBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), PostedBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), AvailableBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), ShowResultingLedgerAccountBalances: moderntreasury.F(true)}, {Amount: moderntreasury.F(int64(0)), Direction: moderntreasury.F(moderntreasury.LedgerTransactionUpdateParamsLedgerEntriesDirectionCredit), LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), LockVersion: moderntreasury.F(int64(0)), PendingBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), PostedBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), AvailableBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), ShowResultingLedgerAccountBalances: moderntreasury.F(true)}}),
			Metadata:      moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}),
			Status:        moderntreasury.F(moderntreasury.LedgerTransactionUpdateParamsStatusArchived),
		},
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerTransactionListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerTransactions.List(context.TODO(), moderntreasury.LedgerTransactionListParams{
		AfterCursor:             moderntreasury.F("string"),
		EffectiveAt:             moderntreasury.F(map[string]string{"foo": "string"}),
		EffectiveDate:           moderntreasury.F(map[string]time.Time{"foo": time.Now()}),
		ExternalID:              moderntreasury.F("string"),
		ID:                      moderntreasury.F(map[string]string{"foo": "string"}),
		LedgerAccountCategoryID: moderntreasury.F("string"),
		LedgerAccountID:         moderntreasury.F("string"),
		LedgerID:                moderntreasury.F("string"),
		Metadata:                moderntreasury.F(map[string]string{"foo": "string"}),
		OrderBy:                 moderntreasury.F(moderntreasury.LedgerTransactionListParamsOrderBy{CreatedAt: moderntreasury.F(moderntreasury.LedgerTransactionListParamsOrderByCreatedAtAsc), EffectiveAt: moderntreasury.F(moderntreasury.LedgerTransactionListParamsOrderByEffectiveAtAsc)}),
		PerPage:                 moderntreasury.F(int64(0)),
		PostedAt:                moderntreasury.F(map[string]time.Time{"foo": time.Now()}),
		Status:                  moderntreasury.F(moderntreasury.LedgerTransactionListParamsStatusPending),
		UpdatedAt:               moderntreasury.F(map[string]time.Time{"foo": time.Now()}),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
