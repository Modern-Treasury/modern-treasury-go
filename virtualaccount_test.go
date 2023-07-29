// File generated from our OpenAPI spec by Stainless.

package moderntreasury_test

import (
	"context"
	"errors"
	"testing"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestVirtualAccountNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.VirtualAccounts.New(context.TODO(), moderntreasury.VirtualAccountNewParams{
		InternalAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Name:              moderntreasury.F("string"),
		AccountDetails: moderntreasury.F([]moderntreasury.VirtualAccountNewParamsAccountDetail{{
			AccountNumber:     moderntreasury.F("string"),
			AccountNumberType: moderntreasury.F(moderntreasury.VirtualAccountNewParamsAccountDetailsAccountNumberTypeClabe),
		}, {
			AccountNumber:     moderntreasury.F("string"),
			AccountNumberType: moderntreasury.F(moderntreasury.VirtualAccountNewParamsAccountDetailsAccountNumberTypeClabe),
		}, {
			AccountNumber:     moderntreasury.F("string"),
			AccountNumberType: moderntreasury.F(moderntreasury.VirtualAccountNewParamsAccountDetailsAccountNumberTypeClabe),
		}}),
		CounterpartyID:        moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CreditLedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		DebitLedgerAccountID:  moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Description:           moderntreasury.F("string"),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		RoutingDetails: moderntreasury.F([]moderntreasury.VirtualAccountNewParamsRoutingDetail{{
			RoutingNumber:     moderntreasury.F("string"),
			RoutingNumberType: moderntreasury.F(moderntreasury.VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeAba),
			PaymentType:       moderntreasury.F(moderntreasury.VirtualAccountNewParamsRoutingDetailsPaymentTypeACH),
		}, {
			RoutingNumber:     moderntreasury.F("string"),
			RoutingNumberType: moderntreasury.F(moderntreasury.VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeAba),
			PaymentType:       moderntreasury.F(moderntreasury.VirtualAccountNewParamsRoutingDetailsPaymentTypeACH),
		}, {
			RoutingNumber:     moderntreasury.F("string"),
			RoutingNumberType: moderntreasury.F(moderntreasury.VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeAba),
			PaymentType:       moderntreasury.F(moderntreasury.VirtualAccountNewParamsRoutingDetailsPaymentTypeACH),
		}}),
		IdempotencyKey: moderntreasury.F("string"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVirtualAccountGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.VirtualAccounts.Get(context.TODO(), "string")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVirtualAccountUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.VirtualAccounts.Update(
		context.TODO(),
		"string",
		moderntreasury.VirtualAccountUpdateParams{
			CounterpartyID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Metadata: moderntreasury.F(map[string]string{
				"foo": "string",
			}),
			Name: moderntreasury.F("string"),
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

func TestVirtualAccountListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.VirtualAccounts.List(context.TODO(), moderntreasury.VirtualAccountListParams{
		AfterCursor:       moderntreasury.F("string"),
		CounterpartyID:    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		InternalAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		PerPage: moderntreasury.F(int64(0)),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVirtualAccountDelete(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.VirtualAccounts.Delete(context.TODO(), "string")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
