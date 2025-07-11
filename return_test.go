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

func TestReturnNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Returns.New(context.TODO(), moderntreasury.ReturnNewParams{
		ReturnableID:          moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ReturnableType:        moderntreasury.F(moderntreasury.ReturnNewParamsReturnableTypeIncomingPaymentDetail),
		AdditionalInformation: moderntreasury.F("additional_information"),
		Code:                  moderntreasury.F(moderntreasury.ReturnNewParamsCode901),
		Corrections: moderntreasury.F(moderntreasury.ReturnNewParamsCorrections{
			AccountNumber:                  moderntreasury.F("account_number"),
			CompanyID:                      moderntreasury.F("company_id"),
			CompanyName:                    moderntreasury.F("company_name"),
			IndividualIdentificationNumber: moderntreasury.F("individual_identification_number"),
			RoutingNumber:                  moderntreasury.F("routing_number"),
			TransactionCode:                moderntreasury.F("transaction_code"),
		}),
		Data:        moderntreasury.F[any](map[string]interface{}{}),
		DateOfDeath: moderntreasury.F(time.Now()),
		Reason:      moderntreasury.F("reason"),
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
	_, err := client.Returns.Get(context.TODO(), "id")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReturnListWithOptionalParams(t *testing.T) {
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
	_, err := client.Returns.List(context.TODO(), moderntreasury.ReturnListParams{
		AfterCursor:       moderntreasury.F("after_cursor"),
		CounterpartyID:    moderntreasury.F("counterparty_id"),
		InternalAccountID: moderntreasury.F("internal_account_id"),
		PerPage:           moderntreasury.F(int64(0)),
		ReturnableID:      moderntreasury.F("returnable_id"),
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
