// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

// LedgerAccountSettlementAccountEntryService contains methods and other services
// that help with interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLedgerAccountSettlementAccountEntryService] method instead.
type LedgerAccountSettlementAccountEntryService struct {
	Options []option.RequestOption
}

// NewLedgerAccountSettlementAccountEntryService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewLedgerAccountSettlementAccountEntryService(opts ...option.RequestOption) (r *LedgerAccountSettlementAccountEntryService) {
	r = &LedgerAccountSettlementAccountEntryService{}
	r.Options = opts
	return
}

// Add ledger entries to a draft ledger account settlement.
func (r *LedgerAccountSettlementAccountEntryService) Update(ctx context.Context, id string, body LedgerAccountSettlementAccountEntryUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/ledger_account_settlements/%s/ledger_entries", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Remove ledger entries from a draft ledger account settlement.
func (r *LedgerAccountSettlementAccountEntryService) Delete(ctx context.Context, id string, body LedgerAccountSettlementAccountEntryDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/ledger_account_settlements/%s/ledger_entries", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type LedgerAccountSettlementAccountEntryUpdateParams struct {
	// The ids of the ledger entries that are to be added or removed from the ledger
	// account settlement.
	LedgerEntryIDs param.Field[[]string] `json:"ledger_entry_ids,required" format:"uuid"`
}

func (r LedgerAccountSettlementAccountEntryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerAccountSettlementAccountEntryDeleteParams struct {
	// The ids of the ledger entries that are to be added or removed from the ledger
	// account settlement.
	LedgerEntryIDs param.Field[[]interface{}] `json:"ledger_entry_ids,required"`
}

func (r LedgerAccountSettlementAccountEntryDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
