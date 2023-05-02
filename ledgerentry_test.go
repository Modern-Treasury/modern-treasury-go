package moderntreasury_test

import (
	"context"
	"errors"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestLedgerEntryGet(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerEntries.Get(
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

func TestLedgerEntryListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerEntries.List(context.TODO(), moderntreasury.LedgerEntryListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0)), ID: moderntreasury.F(map[string]string{"foo": "string"}), LedgerAccountID: moderntreasury.F("string"), LedgerTransactionID: moderntreasury.F("string"), EffectiveDate: moderntreasury.F(map[string]time.Time{"foo": time.Now()}), EffectiveAt: moderntreasury.F(map[string]string{"foo": "string"}), UpdatedAt: moderntreasury.F(map[string]time.Time{"foo": time.Now()}), AsOfLockVersion: moderntreasury.F(int64(0)), LedgerAccountLockVersion: moderntreasury.F(map[string]int64{"foo": int64(0)}), LedgerAccountCategoryID: moderntreasury.F("string"), ShowDeleted: moderntreasury.F(true), Direction: moderntreasury.F(moderntreasury.LedgerEntryListParamsDirectionCredit), Status: moderntreasury.F(moderntreasury.LedgerEntryListParamsStatusPending), OrderBy: moderntreasury.F(moderntreasury.LedgerEntryListParamsOrderBy{CreatedAt: moderntreasury.F(moderntreasury.LedgerEntryListParamsOrderByCreatedAtAsc), EffectiveAt: moderntreasury.F(moderntreasury.LedgerEntryListParamsOrderByEffectiveAtAsc)})})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
