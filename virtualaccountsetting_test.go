// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Modern-Treasury/modern-treasury-go/v2"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

func TestVirtualAccountSettingNewWithOptionalParams(t *testing.T) {
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
	_, err := client.VirtualAccountSettings.New(context.TODO(), moderntreasury.VirtualAccountSettingNewParams{
		AllocationType:                      moderntreasury.F("allocation_type"),
		InternalAccountID:                   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		AllocationIdentifier:                moderntreasury.F("allocation_identifier"),
		AllocationLength:                    moderntreasury.F(int64(0)),
		AllocationRangeEnd:                  moderntreasury.F("allocation_range_end"),
		AllocationRangeStart:                moderntreasury.F("allocation_range_start"),
		ExternalID:                          moderntreasury.F("external_id"),
		GeneratedAllocationIdentifierLength: moderntreasury.F(int64(0)),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVirtualAccountSettingListWithOptionalParams(t *testing.T) {
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
	_, err := client.VirtualAccountSettings.List(context.TODO(), moderntreasury.VirtualAccountSettingListParams{
		AfterCursor: moderntreasury.F("after_cursor"),
		ExternalID:  moderntreasury.F("external_id"),
		PerPage:     moderntreasury.F(int64(0)),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
