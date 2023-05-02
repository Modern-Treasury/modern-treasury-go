package moderntreasury_test

import (
	"context"
	"errors"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestIncomingPaymentDetailGet(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.IncomingPaymentDetails.Get(
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

func TestIncomingPaymentDetailUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.IncomingPaymentDetails.Update(
		context.TODO(),
		"string",
		moderntreasury.IncomingPaymentDetailUpdateParams{Metadata: moderntreasury.F(map[string]string{"foo": "string"})},
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIncomingPaymentDetailListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.IncomingPaymentDetails.List(context.TODO(), moderntreasury.IncomingPaymentDetailListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0)), Direction: moderntreasury.F(moderntreasury.IncomingPaymentDetailListParamsDirectionCredit), Status: moderntreasury.F(moderntreasury.IncomingPaymentDetailListParamsStatusCompleted), Type: moderntreasury.F(moderntreasury.IncomingPaymentDetailListParamsTypeACH), AsOfDateStart: moderntreasury.F(time.Now()), AsOfDateEnd: moderntreasury.F(time.Now()), Metadata: moderntreasury.F(map[string]string{"foo": "string"}), VirtualAccountID: moderntreasury.F("string")})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIncomingPaymentDetailNewAsyncWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.IncomingPaymentDetails.NewAsync(context.TODO(), moderntreasury.IncomingPaymentDetailNewAsyncParams{Type: moderntreasury.F(moderntreasury.IncomingPaymentDetailNewAsyncParamsTypeACH), Direction: moderntreasury.F(moderntreasury.IncomingPaymentDetailNewAsyncParamsDirectionCredit), Amount: moderntreasury.F(int64(0)), Currency: moderntreasury.F(shared.CurrencyAed), InternalAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), VirtualAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), AsOfDate: moderntreasury.F(time.Now())})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
