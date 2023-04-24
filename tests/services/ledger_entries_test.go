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

func TestLedgerEntryGet(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerEntries.Get(
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

func TestLedgerEntryListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerEntries.List(context.TODO(), &requests.LedgerEntryListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0)), ID: moderntreasury.F(map[string]string{"foo": "string"}), LedgerAccountID: moderntreasury.F("string"), LedgerTransactionID: moderntreasury.F("string"), EffectiveDate: moderntreasury.F(map[string]time.Time{"foo": time.Now()}), EffectiveAt: moderntreasury.F(map[string]string{"foo": "string"}), UpdatedAt: moderntreasury.F(map[string]time.Time{"foo": time.Now()}), AsOfLockVersion: moderntreasury.F(int64(0)), LedgerAccountLockVersion: moderntreasury.F(map[string]int64{"foo": int64(0)}), LedgerAccountCategoryID: moderntreasury.F("string"), ShowDeleted: moderntreasury.F(true), Direction: moderntreasury.F(requests.LedgerEntryListParamsDirectionCredit), Status: moderntreasury.F(requests.LedgerEntryListParamsStatusPending), OrderBy: moderntreasury.F(requests.LedgerEntryListParamsOrderBy{CreatedAt: moderntreasury.F(requests.LedgerEntryListParamsOrderByCreatedAtAsc), EffectiveAt: moderntreasury.F(requests.LedgerEntryListParamsOrderByEffectiveAtAsc)})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
