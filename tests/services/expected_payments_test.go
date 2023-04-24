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

func TestExpectedPaymentNewWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.ExpectedPayments.New(context.TODO(), &requests.ExpectedPaymentNewParams{AmountUpperBound: moderntreasury.F(int64(0)), AmountLowerBound: moderntreasury.F(int64(0)), Direction: moderntreasury.F(requests.ExpectedPaymentNewParamsDirectionCredit), InternalAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Type: moderntreasury.F(requests.ExpectedPaymentTypeACH), Currency: moderntreasury.F(requests.CurrencyAed), DateUpperBound: moderntreasury.F(time.Now()), DateLowerBound: moderntreasury.F(time.Now()), Description: moderntreasury.F("string"), StatementDescriptor: moderntreasury.F("string"), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), CounterpartyID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), RemittanceInformation: moderntreasury.F("string"), LineItems: moderntreasury.F([]requests.ExpectedPaymentNewParamsLineItems{{Amount: moderntreasury.F(int64(0)), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), Description: moderntreasury.F("string"), AccountingCategoryID: moderntreasury.F("string")}, {Amount: moderntreasury.F(int64(0)), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), Description: moderntreasury.F("string"), AccountingCategoryID: moderntreasury.F("string")}, {Amount: moderntreasury.F(int64(0)), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), Description: moderntreasury.F("string"), AccountingCategoryID: moderntreasury.F("string")}})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExpectedPaymentGet(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.ExpectedPayments.Get(
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

func TestExpectedPaymentUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.ExpectedPayments.Update(
		context.TODO(),
		"string",
		&requests.ExpectedPaymentUpdateParams{AmountUpperBound: moderntreasury.F(int64(0)), AmountLowerBound: moderntreasury.F(int64(0)), Direction: moderntreasury.F(requests.ExpectedPaymentUpdateParamsDirectionCredit), InternalAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Type: moderntreasury.F(requests.ExpectedPaymentTypeACH), Currency: moderntreasury.F(requests.CurrencyAed), DateUpperBound: moderntreasury.F(time.Now()), DateLowerBound: moderntreasury.F(time.Now()), Description: moderntreasury.F("string"), StatementDescriptor: moderntreasury.F("string"), Metadata: moderntreasury.F(map[string]string{"key": "value", "foo": "bar", "modern": "treasury"}), CounterpartyID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), RemittanceInformation: moderntreasury.F("string")},
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

func TestExpectedPaymentListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.ExpectedPayments.List(context.TODO(), &requests.ExpectedPaymentListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0)), Status: moderntreasury.F(requests.ExpectedPaymentListParamsStatusArchived), InternalAccountID: moderntreasury.F("string"), Direction: moderntreasury.F(requests.ExpectedPaymentListParamsDirectionCredit), Type: moderntreasury.F(requests.ExpectedPaymentListParamsTypeACH), CounterpartyID: moderntreasury.F("string"), Metadata: moderntreasury.F(map[string]string{"foo": "string"}), CreatedAtLowerBound: moderntreasury.F(time.Now()), CreatedAtUpperBound: moderntreasury.F(time.Now())})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExpectedPaymentDelete(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.ExpectedPayments.Delete(
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
