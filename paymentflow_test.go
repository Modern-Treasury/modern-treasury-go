package moderntreasury_test

import (
	"context"
	"errors"
	"testing"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestPaymentFlowNewWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.PaymentFlows.New(context.TODO(), moderntreasury.PaymentFlowNewParams{
		Amount:               moderntreasury.F(int64(0)),
		CounterpartyID:       moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Currency:             moderntreasury.F("string"),
		Direction:            moderntreasury.F(moderntreasury.PaymentFlowNewParamsDirectionCredit),
		OriginatingAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		IdempotencyKey:       moderntreasury.F("string"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentFlowGetWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.PaymentFlows.Get(
		context.TODO(),
		"string",
		moderntreasury.PaymentFlowGetParams{
			IdempotencyKey: moderntreasury.F("string"),
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

func TestPaymentFlowUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.PaymentFlows.Update(
		context.TODO(),
		"string",
		moderntreasury.PaymentFlowUpdateParams{
			Status:         moderntreasury.F(moderntreasury.PaymentFlowUpdateParamsStatusCancelled),
			IdempotencyKey: moderntreasury.F("string"),
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

func TestPaymentFlowListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.PaymentFlows.List(context.TODO(), moderntreasury.PaymentFlowListParams{
		AfterCursor:          moderntreasury.F("string"),
		ClientToken:          moderntreasury.F("string"),
		CounterpartyID:       moderntreasury.F("string"),
		OriginatingAccountID: moderntreasury.F("string"),
		PaymentOrderID:       moderntreasury.F("string"),
		PerPage:              moderntreasury.F(int64(0)),
		ReceivingAccountID:   moderntreasury.F("string"),
		Status:               moderntreasury.F("string"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
