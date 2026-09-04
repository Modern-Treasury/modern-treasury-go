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

func TestIdentificationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Identifications.New(context.TODO(), moderntreasury.IdentificationNewParams{
		IDNumber:      moderntreasury.F("id_number"),
		IDType:        moderntreasury.F(moderntreasury.IdentificationNewParamsIDTypeArCuil),
		LegalEntityID: moderntreasury.F("legal_entity_id"),
		Documents: moderntreasury.F([]moderntreasury.IdentificationNewParamsDocument{{
			DocumentType: moderntreasury.F(moderntreasury.IdentificationNewParamsDocumentsDocumentTypeArticlesOfIncorporation),
			FileData:     moderntreasury.F("file_data"),
			Filename:     moderntreasury.F("filename"),
		}}),
		ExpirationDate: moderntreasury.F(time.Now()),
		IssuingCountry: moderntreasury.F("issuing_country"),
		IssuingRegion:  moderntreasury.F("issuing_region"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIdentificationGet(t *testing.T) {
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
	_, err := client.Identifications.Get(context.TODO(), "id")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIdentificationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Identifications.Update(
		context.TODO(),
		"id",
		moderntreasury.IdentificationUpdateParams{
			ExpirationDate: moderntreasury.F(time.Now()),
			IDNumber:       moderntreasury.F("id_number"),
			IDType:         moderntreasury.F(moderntreasury.IdentificationUpdateParamsIDTypeArCuil),
			IssuingCountry: moderntreasury.F("issuing_country"),
			IssuingRegion:  moderntreasury.F("issuing_region"),
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
