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

func TestReturnNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Returns.New(context.TODO(), moderntreasury.ReturnNewParams{
		ReturnableID:          moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ReturnableType:        moderntreasury.F(moderntreasury.ReturnNewParamsReturnableTypeIncomingPaymentDetail),
		AdditionalInformation: moderntreasury.F("string"),
		Code:                  moderntreasury.F(moderntreasury.ReturnNewParamsCode901),
		DateOfDeath:           moderntreasury.F(time.Now()),
		Reason:                moderntreasury.F("string"),
		IdempotencyKey:        moderntreasury.F("string"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReturnGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Returns.Get(context.TODO(), "string")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReturnListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Returns.List(context.TODO(), moderntreasury.ReturnListParams{
		AfterCursor:       moderntreasury.F("string"),
		CounterpartyID:    moderntreasury.F("string"),
		InternalAccountID: moderntreasury.F("string"),
		PerPage:           moderntreasury.F(int64(0)),
		ReturnableID:      moderntreasury.F("string"),
		ReturnableType:    moderntreasury.F(moderntreasury.ReturnListParamsReturnableTypeIncomingPaymentDetail),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
