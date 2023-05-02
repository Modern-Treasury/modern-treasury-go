package main

import (
	"encoding/json"
	"fmt"
	mt "modern_treasury"
	"modern_treasury/account_details"
	"modern_treasury/expected_payments"
	"modern_treasury/internal_accounts"
	"os"
)

type CustomAccountDetailsResponse struct {
	account_details.AccountDetail
	CustomField bool `json:"custom_field"`
}

func toJSON(v interface{}) string {
	bytes, _ := json.MarshalIndent(v, "", "  ")
	return string(bytes)
}

// In this example, we will create, update, and delete an expected
// payment using the first internal account from the list API
func main() {
	client := mt.NewModernTreasuryWithParams(mt.ModernTreasuryParams{
		APIKey:         os.Getenv("MODERN_TREASURY_API_KEY"),
		OrganizationId: os.Getenv("MODERN_TREASURY_ORGANIZATION_ID"),
	})

	// Read

	internalAccounts, err := client.InternalAccounts.List(&internal_accounts.InternalAccountsListQuery{
		PerPage: mt.P(1),
	})

	if err != nil {
		panic(fmt.Errorf("Failed to list internal accounts: %w", err))
	} else if len(internalAccounts) == 0 {
		panic(fmt.Errorf("Internal account listing returned an empty array: %s", toJSON(internalAccounts)))
	}

	fmt.Printf("First internal account from listing: %v\n", toJSON(internalAccounts[0]))

	// Create

	createdExpectedPayment, err := client.ExpectedPayments.Create(expected_payments.ExpectedPaymentCreateRequest{
		Type:              mt.P(expected_payments.ExpectedPaymentTypeACH),
		AmountLowerBound:  mt.P(100),
		AmountUpperBound:  mt.P(100),
		Direction:         mt.P(expected_payments.ExpectedPaymentCreateRequestDirectionCredit),
		InternalAccountId: internalAccounts[0].Id,
	})

	if err != nil {
		panic(fmt.Errorf("Failed to create expected payment: %w", err))
	}

	fmt.Printf("Create expected payment result: %s\n", toJSON(createdExpectedPayment))

	// Update

	updatedExpectedPayment, err := client.ExpectedPayments.Update(*createdExpectedPayment.Id, expected_payments.ExpectedPaymentUpdateRequest{
		Description: mt.P("New description"),
	})

	if err != nil {
		panic(fmt.Errorf("Failed to update expected payment: %w", err))
	}

	fmt.Printf("Update expected payment result: %s\n", toJSON(updatedExpectedPayment))

	// Delete

	deletedPayment, err := client.ExpectedPayments.Delete(*updatedExpectedPayment.Id)

	if err != nil {
		panic(fmt.Errorf("Failed to delete expected payment: %w", err))
	}

	fmt.Printf("Delete expected payment result: %s\n", toJSON(deletedPayment))
}

type B struct{ A }
type A struct{}
