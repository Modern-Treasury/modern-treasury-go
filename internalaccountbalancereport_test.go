// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

func TestInternalAccountBalanceReportNew(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.InternalAccounts.BalanceReports.New(
		context.TODO(),
		"string",
		moderntreasury.BalanceReportNewParams{
			AsOfDate:          moderntreasury.F(time.Now()),
			AsOfTime:          moderntreasury.F("string"),
			BalanceReportType: moderntreasury.F(moderntreasury.BalanceReportNewParamsBalanceReportTypeIntraday),
			Balances: moderntreasury.F([]moderntreasury.BalanceReportNewParamsBalance{{
				Amount:         moderntreasury.F(int64(0)),
				BalanceType:    moderntreasury.F(moderntreasury.BalanceReportNewParamsBalancesBalanceTypeClosingAvailable),
				VendorCode:     moderntreasury.F("string"),
				VendorCodeType: moderntreasury.F("string"),
			}, {
				Amount:         moderntreasury.F(int64(0)),
				BalanceType:    moderntreasury.F(moderntreasury.BalanceReportNewParamsBalancesBalanceTypeClosingAvailable),
				VendorCode:     moderntreasury.F("string"),
				VendorCodeType: moderntreasury.F("string"),
			}, {
				Amount:         moderntreasury.F(int64(0)),
				BalanceType:    moderntreasury.F(moderntreasury.BalanceReportNewParamsBalancesBalanceTypeClosingAvailable),
				VendorCode:     moderntreasury.F("string"),
				VendorCodeType: moderntreasury.F("string"),
			}}),
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

func TestInternalAccountBalanceReportGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.InternalAccounts.BalanceReports.Get(
		context.TODO(),
		"string",
		"string",
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInternalAccountBalanceReportListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.InternalAccounts.BalanceReports.List(
		context.TODO(),
		"string",
		moderntreasury.BalanceReportListParams{
			AfterCursor:       moderntreasury.F("string"),
			AsOfDate:          moderntreasury.F(time.Now()),
			BalanceReportType: moderntreasury.F(moderntreasury.BalanceReportListParamsBalanceReportTypeIntraday),
			PerPage:           moderntreasury.F(int64(0)),
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

func TestInternalAccountBalanceReportDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithOrganizationID("my-organization-ID"),
	)
	err := client.InternalAccounts.BalanceReports.Delete(
		context.TODO(),
		"string",
		"string",
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
