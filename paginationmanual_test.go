// File generated from our OpenAPI spec by Stainless.

package moderntreasury_test

import (
	"context"
	"fmt"
	"testing"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestManualPagination(t *testing.T) {
	client := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	page, err := client.ExternalAccounts.List(context.TODO(), moderntreasury.ExternalAccountListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, externalAccount := range page.Items {
		fmt.Printf("%+v\n", externalAccount)
	}
	// Prism mock isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, externalAccount := range page.Items {
			fmt.Printf("%+v\n", externalAccount)
		}
	}
}
