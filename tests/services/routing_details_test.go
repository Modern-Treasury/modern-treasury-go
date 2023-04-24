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

func TestRoutingDetailNewWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.RoutingDetails.New(
		context.TODO(),
		requests.RoutingDetailNewParamsAccountsTypeExternalAccounts,
		"string",
		&requests.RoutingDetailNewParams{RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(requests.RoutingDetailNewParamsRoutingNumberTypeAba), PaymentType: moderntreasury.F(requests.RoutingDetailNewParamsPaymentTypeACH)},
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

func TestRoutingDetailGet(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.RoutingDetails.Get(
		context.TODO(),
		requests.AccountsTypeExternalAccounts,
		"string",
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

func TestRoutingDetailListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.RoutingDetails.List(
		context.TODO(),
		requests.AccountsTypeExternalAccounts,
		"string",
		&requests.RoutingDetailListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0))},
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

func TestRoutingDetailDelete(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	err := c.RoutingDetails.Delete(
		context.TODO(),
		requests.RoutingDetailDeleteParamsAccountsTypeExternalAccounts,
		"string",
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
