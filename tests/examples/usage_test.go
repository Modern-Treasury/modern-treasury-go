package examples

import (
	"context"
	"fmt"
	"testing"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
)

func TestUsage(t *testing.T) {
	client := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	externalAccount, err := client.ExternalAccounts.New(context.TODO(), &requests.ExternalAccountNewParams{
		CounterpartyID: moderntreasury.F("9eba513a-53fd-4d6d-ad52-ccce122ab92a"),
		Name:           moderntreasury.F("my bank"),
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", externalAccount)
}
