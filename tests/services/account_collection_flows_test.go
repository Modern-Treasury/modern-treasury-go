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

func TestAccountCollectionFlowNew(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.AccountCollectionFlows.New(context.TODO(), &requests.AccountCollectionFlowNewParams{CounterpartyID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), PaymentTypes: moderntreasury.F([]string{"string", "string", "string"})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountCollectionFlowGet(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.AccountCollectionFlows.Get(
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

func TestAccountCollectionFlowUpdate(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.AccountCollectionFlows.Update(
		context.TODO(),
		"string",
		&requests.AccountCollectionFlowUpdateParams{Status: moderntreasury.F(requests.AccountCollectionFlowUpdateParamsStatusCancelled)},
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

func TestAccountCollectionFlowListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.AccountCollectionFlows.List(context.TODO(), &requests.AccountCollectionFlowListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0)), ClientToken: moderntreasury.F("string"), Status: moderntreasury.F("string"), CounterpartyID: moderntreasury.F("string"), ExternalAccountID: moderntreasury.F("string")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
