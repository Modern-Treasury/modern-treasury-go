// File generated from our OpenAPI spec by Stainless.

package moderntreasury_test

import (
	"context"
	"errors"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestIncomingPaymentDetailGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.IncomingPaymentDetails.Get(context.TODO(), "string")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIncomingPaymentDetailUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.IncomingPaymentDetails.Update(
		context.TODO(),
		"string",
		moderntreasury.IncomingPaymentDetailUpdateParams{
			Metadata: moderntreasury.F(map[string]string{
				"foo": "string",
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

func TestIncomingPaymentDetailListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.IncomingPaymentDetails.List(context.TODO(), moderntreasury.IncomingPaymentDetailListParams{
		AfterCursor:   moderntreasury.F("string"),
		AsOfDateEnd:   moderntreasury.F(time.Now()),
		AsOfDateStart: moderntreasury.F(time.Now()),
		Direction:     moderntreasury.F(moderntreasury.IncomingPaymentDetailListParamsDirectionCredit),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		PerPage:          moderntreasury.F(int64(0)),
		Status:           moderntreasury.F(moderntreasury.IncomingPaymentDetailListParamsStatusCompleted),
		Type:             moderntreasury.F(moderntreasury.IncomingPaymentDetailListParamsTypeACH),
		VirtualAccountID: moderntreasury.F("string"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIncomingPaymentDetailNewAsyncWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.IncomingPaymentDetails.NewAsync(context.TODO(), moderntreasury.IncomingPaymentDetailNewAsyncParams{
		Amount:            moderntreasury.F(int64(0)),
		AsOfDate:          moderntreasury.F(time.Now()),
		Currency:          moderntreasury.F(shared.CurrencyAed),
		Description:       moderntreasury.F("string"),
		Direction:         moderntreasury.F(moderntreasury.IncomingPaymentDetailNewAsyncParamsDirectionCredit),
		InternalAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Type:              moderntreasury.F(moderntreasury.IncomingPaymentDetailNewAsyncParamsTypeACH),
		VirtualAccountID:  moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		IdempotencyKey:    moderntreasury.F("string"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
