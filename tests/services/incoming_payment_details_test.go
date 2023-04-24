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

func TestIncomingPaymentDetailGet(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.IncomingPaymentDetails.Get(
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

func TestIncomingPaymentDetailUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.IncomingPaymentDetails.Update(
		context.TODO(),
		"string",
		&requests.IncomingPaymentDetailUpdateParams{Metadata: moderntreasury.F(map[string]string{"foo": "string"})},
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

func TestIncomingPaymentDetailListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.IncomingPaymentDetails.List(context.TODO(), &requests.IncomingPaymentDetailListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0)), Direction: moderntreasury.F(requests.IncomingPaymentDetailListParamsDirectionCredit), Status: moderntreasury.F(requests.IncomingPaymentDetailListParamsStatusCompleted), Type: moderntreasury.F(requests.IncomingPaymentDetailListParamsTypeACH), AsOfDateStart: moderntreasury.F(time.Now()), AsOfDateEnd: moderntreasury.F(time.Now()), Metadata: moderntreasury.F(map[string]string{"foo": "string"}), VirtualAccountID: moderntreasury.F("string")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIncomingPaymentDetailNewAsyncWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.IncomingPaymentDetails.NewAsync(context.TODO(), &requests.IncomingPaymentDetailNewAsyncParams{Type: moderntreasury.F(requests.IncomingPaymentDetailNewAsyncParamsTypeACH), Direction: moderntreasury.F(requests.IncomingPaymentDetailNewAsyncParamsDirectionCredit), Amount: moderntreasury.F(int64(0)), Currency: moderntreasury.F(requests.CurrencyAed), InternalAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), VirtualAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), AsOfDate: moderntreasury.F(time.Now())})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
