// File generated from our OpenAPI spec by Stainless.

package moderntreasury_test

import (
	"context"
	"errors"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestVirtualAccountNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.VirtualAccounts.New(context.TODO(), moderntreasury.VirtualAccountNewParams{
		InternalAccountID:     moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Name:                  moderntreasury.F("string"),
		AccountDetails:        moderntreasury.F([]moderntreasury.AccountDetailParam{{ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(moderntreasury.AccountDetailAccountNumberTypeClabe), AccountNumberSafe: moderntreasury.F("string")}, {ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(moderntreasury.AccountDetailAccountNumberTypeClabe), AccountNumberSafe: moderntreasury.F("string")}, {ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(moderntreasury.AccountDetailAccountNumberTypeClabe), AccountNumberSafe: moderntreasury.F("string")}}),
		CounterpartyID:        moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CreditLedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		DebitLedgerAccountID:  moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Description:           moderntreasury.F("string"),
		Metadata:              moderntreasury.F(map[string]string{"foo": "string"}),
		RoutingDetails:        moderntreasury.F([]moderntreasury.RoutingDetailParam{{ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(moderntreasury.RoutingDetailRoutingNumberTypeAba), PaymentType: moderntreasury.F(moderntreasury.RoutingDetailPaymentTypeACH), BankName: moderntreasury.F("string"), BankAddress: moderntreasury.F(moderntreasury.RoutingDetailBankAddressParam{ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")})}, {ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(moderntreasury.RoutingDetailRoutingNumberTypeAba), PaymentType: moderntreasury.F(moderntreasury.RoutingDetailPaymentTypeACH), BankName: moderntreasury.F("string"), BankAddress: moderntreasury.F(moderntreasury.RoutingDetailBankAddressParam{ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")})}, {ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(moderntreasury.RoutingDetailRoutingNumberTypeAba), PaymentType: moderntreasury.F(moderntreasury.RoutingDetailPaymentTypeACH), BankName: moderntreasury.F("string"), BankAddress: moderntreasury.F(moderntreasury.RoutingDetailBankAddressParam{ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")})}}),
		IdempotencyKey:        moderntreasury.F("string"),
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
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	err := c.VirtualAccounts.Get(
		context.TODO(),
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

func TestVirtualAccountUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.VirtualAccounts.Update(
		context.TODO(),
		"string",
		moderntreasury.VirtualAccountUpdateParams{
			CounterpartyID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Metadata:       moderntreasury.F(map[string]string{"foo": "string"}),
			Name:           moderntreasury.F("string"),
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
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.VirtualAccounts.List(context.TODO(), moderntreasury.VirtualAccountListParams{
		AfterCursor:       moderntreasury.F("string"),
		CounterpartyID:    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		InternalAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Metadata:          moderntreasury.F(map[string]string{"foo": "string"}),
		PerPage:           moderntreasury.F(int64(0)),
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
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.VirtualAccounts.Delete(
		context.TODO(),
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
