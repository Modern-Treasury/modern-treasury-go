// File generated from our OpenAPI spec by Stainless.

package moderntreasury_test

import (
	"context"
	"errors"
	"testing"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestLedgerAccountPayoutNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerAccountPayouts.New(context.TODO(), moderntreasury.LedgerAccountPayoutNewParams{
		FundingLedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PayoutLedgerAccountID:  moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Description:            moderntreasury.F("string"),
		EffectiveAtUpperBound:  moderntreasury.F("14:15:22Z"),
		Metadata:               moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}),
		Status:                 moderntreasury.F(moderntreasury.LedgerAccountPayoutNewParamsStatusPending),
		IdempotencyKey:         moderntreasury.F("string"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerAccountPayoutUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerAccountPayouts.Update(
		context.TODO(),
		"string",
		moderntreasury.LedgerAccountPayoutUpdateParams{
			Description: moderntreasury.F("string"),
			Metadata:    moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}),
			Status:      moderntreasury.F(moderntreasury.LedgerAccountPayoutUpdateParamsStatusPosted),
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

func TestLedgerAccountPayoutListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerAccountPayouts.List(context.TODO(), moderntreasury.LedgerAccountPayoutListParams{
		AfterCursor:           moderntreasury.F("string"),
		PayoutLedgerAccountID: moderntreasury.F("string"),
		PerPage:               moderntreasury.F(int64(0)),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerAccountPayoutRetireve(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerAccountPayouts.Retireve(context.TODO(), "string")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
