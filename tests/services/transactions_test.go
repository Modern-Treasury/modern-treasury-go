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

func TestTransactionGet(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Transactions.Get(
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

func TestTransactionUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Transactions.Update(
		context.TODO(),
		"string",
		&requests.TransactionUpdateParams{Metadata: moderntreasury.F(map[string]string{"foo": "string"})},
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

func TestTransactionListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Transactions.List(context.TODO(), &requests.TransactionListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0)), InternalAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), VirtualAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Posted: moderntreasury.F(true), AsOfDateStart: moderntreasury.F(time.Now()), AsOfDateEnd: moderntreasury.F(time.Now()), Direction: moderntreasury.F("string"), CounterpartyID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), PaymentType: moderntreasury.F("string"), TransactableType: moderntreasury.F("string"), Description: moderntreasury.F("string"), Metadata: moderntreasury.F(map[string]string{"foo": "string"})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
