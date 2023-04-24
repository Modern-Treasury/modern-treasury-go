package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/core"
	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
)

func TestVirtualAccountNewWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.VirtualAccounts.New(context.TODO(), &requests.VirtualAccountNewParams{Name: moderntreasury.F("string"), Description: moderntreasury.F("string"), CounterpartyID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), InternalAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), AccountDetails: moderntreasury.F([]requests.AccountDetail{{ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(requests.AccountDetailAccountNumberTypeClabe), AccountNumberSafe: moderntreasury.F("string")}, {ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(requests.AccountDetailAccountNumberTypeClabe), AccountNumberSafe: moderntreasury.F("string")}, {ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), AccountNumber: moderntreasury.F("string"), AccountNumberType: moderntreasury.F(requests.AccountDetailAccountNumberTypeClabe), AccountNumberSafe: moderntreasury.F("string")}}), RoutingDetails: moderntreasury.F([]requests.RoutingDetail{{ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(requests.RoutingDetailRoutingNumberTypeAba), PaymentType: moderntreasury.F(requests.RoutingDetailPaymentTypeACH), BankName: moderntreasury.F("string"), BankAddress: moderntreasury.F(requests.RoutingDetailBankAddress{ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")})}, {ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(requests.RoutingDetailRoutingNumberTypeAba), PaymentType: moderntreasury.F(requests.RoutingDetailPaymentTypeACH), BankName: moderntreasury.F("string"), BankAddress: moderntreasury.F(requests.RoutingDetailBankAddress{ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")})}, {ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), DiscardedAt: moderntreasury.F(time.Now()), RoutingNumber: moderntreasury.F("string"), RoutingNumberType: moderntreasury.F(requests.RoutingDetailRoutingNumberTypeAba), PaymentType: moderntreasury.F(requests.RoutingDetailPaymentTypeACH), BankName: moderntreasury.F("string"), BankAddress: moderntreasury.F(requests.RoutingDetailBankAddress{ID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Object: moderntreasury.F("string"), LiveMode: moderntreasury.F(true), CreatedAt: moderntreasury.F(time.Now()), UpdatedAt: moderntreasury.F(time.Now()), Line1: moderntreasury.F("string"), Line2: moderntreasury.F("string"), Locality: moderntreasury.F("string"), Region: moderntreasury.F("string"), PostalCode: moderntreasury.F("string"), Country: moderntreasury.F("string")})}}), DebitLedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), CreditLedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Metadata: moderntreasury.F(map[string]string{"foo": "string"})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVirtualAccountGet(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	err := c.VirtualAccounts.Get(
		context.TODO(),
		"string",
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVirtualAccountUpdateWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.VirtualAccounts.Update(
		context.TODO(),
		"string",
		&requests.VirtualAccountUpdateParams{Name: moderntreasury.F("string"), CounterpartyID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Metadata: moderntreasury.F(map[string]string{"foo": "string"})},
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVirtualAccountListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.VirtualAccounts.List(context.TODO(), &requests.VirtualAccountListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0)), InternalAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), CounterpartyID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"), Metadata: moderntreasury.F(map[string]string{"foo": "string"})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVirtualAccountDelete(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.VirtualAccounts.Delete(
		context.TODO(),
		"string",
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
