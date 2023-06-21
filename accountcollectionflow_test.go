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

func TestAccountCollectionFlowNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.AccountCollectionFlows.New(context.TODO(), moderntreasury.AccountCollectionFlowNewParams{
		CounterpartyID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PaymentTypes:   moderntreasury.F([]string{"string", "string", "string"}),
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

func TestAccountCollectionFlowGetWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.AccountCollectionFlows.Get(
		context.TODO(),
		"string",
		moderntreasury.AccountCollectionFlowGetParams{
			IdempotencyKey: moderntreasury.F("string"),
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

func TestAccountCollectionFlowUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.AccountCollectionFlows.Update(
		context.TODO(),
		"string",
		moderntreasury.AccountCollectionFlowUpdateParams{
			Status:         moderntreasury.F(moderntreasury.AccountCollectionFlowUpdateParamsStatusCancelled),
			IdempotencyKey: moderntreasury.F("string"),
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

func TestAccountCollectionFlowListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.AccountCollectionFlows.List(context.TODO(), moderntreasury.AccountCollectionFlowListParams{
		AfterCursor:       moderntreasury.F("string"),
		ClientToken:       moderntreasury.F("string"),
		CounterpartyID:    moderntreasury.F("string"),
		ExternalAccountID: moderntreasury.F("string"),
		PerPage:           moderntreasury.F(int64(0)),
		Status:            moderntreasury.F("string"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
