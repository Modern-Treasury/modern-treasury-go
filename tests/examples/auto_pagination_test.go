package examples

import (
	"context"
	"fmt"
	"testing"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
)

func TestAutoPagination(t *testing.T) {
	client := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	iter := client.ExternalAccounts.ListAutoPager(context.TODO(), &requests.ExternalAccountListParams{})
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		externalAccount := iter.Current()
		fmt.Printf("%+v\n", externalAccount)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
