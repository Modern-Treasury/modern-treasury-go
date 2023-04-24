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

func TestPaymentOrderReversalNewWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.PaymentOrders.Reversals.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		&requests.PaymentOrderReversalNewParams{Reason: moderntreasury.F(requests.PaymentOrderReversalNewParamsReasonDuplicate), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), LedgerTransaction: moderntreasury.F(requests.PaymentOrderReversalNewParamsLedgerTransaction{Description: moderntreasury.F("string"), Status: moderntreasury.F(requests.PaymentOrderReversalNewParamsLedgerTransactionStatusArchived), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), EffectiveDate: moderntreasury.F(time.Now()), LedgerEntries: moderntreasury.F([]requests.PaymentOrderReversalNewParamsLedgerTransactionLedgerEntries{{Amount: moderntreasury.F(int64(0)), Direction: moderntreasury.F(requests.PaymentOrderReversalNewParamsLedgerTransactionLedgerEntriesDirectionCredit), LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), LockVersion: moderntreasury.F(int64(0)), PendingBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), PostedBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), AvailableBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), ShowResultingLedgerAccountBalances: moderntreasury.F(true)}, {Amount: moderntreasury.F(int64(0)), Direction: moderntreasury.F(requests.PaymentOrderReversalNewParamsLedgerTransactionLedgerEntriesDirectionCredit), LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), LockVersion: moderntreasury.F(int64(0)), PendingBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), PostedBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), AvailableBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), ShowResultingLedgerAccountBalances: moderntreasury.F(true)}, {Amount: moderntreasury.F(int64(0)), Direction: moderntreasury.F(requests.PaymentOrderReversalNewParamsLedgerTransactionLedgerEntriesDirectionCredit), LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), LockVersion: moderntreasury.F(int64(0)), PendingBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), PostedBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), AvailableBalanceAmount: moderntreasury.F(map[string]int64{"foo": int64(0)}), ShowResultingLedgerAccountBalances: moderntreasury.F(true)}}), ExternalID: moderntreasury.F("string"), LedgerableType: moderntreasury.F(requests.PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypeCounterparty), LedgerableID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")})},
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

func TestPaymentOrderReversalGet(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.PaymentOrders.Reversals.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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

func TestPaymentOrderReversalListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.PaymentOrders.Reversals.List(
		context.TODO(),
		"string",
		&requests.PaymentOrderReversalListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0))},
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
