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

func TestLedgerAccountPayoutNewWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerAccountPayouts.New(context.TODO(), &requests.LedgerAccountPayoutNewParams{Description: moderntreasury.F("string"), Status: moderntreasury.F(requests.LedgerAccountPayoutNewParamsStatusPending), PayoutLedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), FundingLedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), EffectiveAtUpperBound: moderntreasury.F("14:15:22Z"), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerAccountPayoutUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerAccountPayouts.Update(
		context.TODO(),
		"string",
		&requests.LedgerAccountPayoutUpdateParams{Description: moderntreasury.F("string"), Status: moderntreasury.F(requests.LedgerAccountPayoutUpdateParamsStatusPosted), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"})},
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

func TestLedgerAccountPayoutListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerAccountPayouts.List(context.TODO(), &requests.LedgerAccountPayoutListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0)), PayoutLedgerAccountID: moderntreasury.F("string")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerAccountPayoutRetireve(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerAccountPayouts.Retireve(
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
