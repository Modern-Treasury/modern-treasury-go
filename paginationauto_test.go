// File generated from our OpenAPI spec by Stainless.

package moderntreasury_test

import (
	"context"
	"testing"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestAutoPagination(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	iter := client.ExternalAccounts.ListAutoPaging(context.TODO(), moderntreasury.ExternalAccountListParams{})
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		externalAccount := iter.Current()
		t.Logf("%+v\n", externalAccount)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
